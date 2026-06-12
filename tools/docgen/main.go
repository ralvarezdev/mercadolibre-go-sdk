// Command docgen fetches MercadoLibre developer documentation pages and writes
// one Markdown file per page under docs/<section>/<slug>.md.
//
// The pages are server-rendered; docgen extracts the article container
// (div.mesh-row__content), converts it to Markdown, and prepends YAML
// front-matter plus a list of api.mercadolibre.com endpoints referenced on the
// page.
//
// Usage:
//
//	go run ./tools/docgen                 # generate every page in the sitemap
//	go run ./tools/docgen -only items     # only slugs containing "items"
//	go run ./tools/docgen -section 04-items
//	go run ./tools/docgen -list           # print the sitemap and exit
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	docsBase  = "https://developers.mercadolibre.com.ve/es_ar/"
	siteHost  = "https://developers.mercadolibre.com.ve"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 " +
		"(KHTML, like Gecko) Chrome/124.0 Safari/537.36"
	headerUserAgent      = "User-Agent"
	headerAcceptLanguage = "Accept-Language"
	acceptLanguageValue  = "es-AR,es;q=0.9,en;q=0.8"
)

func main() {
	only := flag.String("only", "", "only process slugs containing this substring")
	section := flag.String("section", "", "only process this section directory")
	outDir := flag.String("out", "docs", "output directory")
	list := flag.Bool("list", false, "print the sitemap and exit")
	flag.Parse()

	if *list {
		for _, p := range sitemap {
			fmt.Printf("%-22s %-50s %s\n", p.Section, p.Slug, p.Title)
		}
		fmt.Printf("\n%d pages\n", len(sitemap))
		return
	}

	client := &http.Client{Timeout: 30 * time.Second}
	var ok, fail int
	for _, p := range sitemap {
		if *only != "" && !strings.Contains(p.Slug, *only) {
			continue
		}
		if *section != "" && p.Section != *section {
			continue
		}
		if err := generate(client, p, *outDir); err != nil {
			fmt.Printf("FAIL  %s/%s: %v\n", p.Section, p.Slug, err)
			fail++
			continue
		}
		fmt.Printf("ok    %s/%s.md\n", p.Section, p.Slug)
		ok++
		time.Sleep(250 * time.Millisecond) // be polite
	}
	fmt.Printf("\nDone: %d ok, %d failed\n", ok, fail)
	if fail > 0 {
		os.Exit(1)
	}
}

func generate(client *http.Client, p Page, outDir string) error {
	url := docsBase + p.Slug
	root, err := fetch(client, url)
	if err != nil {
		return err
	}
	content := findContent(root)
	if content == nil {
		return fmt.Errorf("article container not found")
	}

	var body strings.Builder
	blockMD(content, &body)
	md := cleanup(body.String())

	endpoints := extractEndpoints(content)

	var out strings.Builder
	out.WriteString("---\n")
	out.WriteString("title: " + yamlString(p.Title) + "\n")
	out.WriteString("slug: " + p.Slug + "\n")
	out.WriteString("section: " + p.Section + "\n")
	out.WriteString("source: " + url + "\n")
	out.WriteString("captured: " + time.Now().Format("2006-01-02") + "\n")
	out.WriteString("---\n\n")
	out.WriteString("# " + p.Title + "\n\n")
	out.WriteString("> Source: <" + url + ">\n\n")
	if len(endpoints) > 0 {
		out.WriteString("## Endpoints referenced\n\n")
		for _, e := range endpoints {
			out.WriteString("- `" + e + "`\n")
		}
		out.WriteString("\n")
	}
	out.WriteString("## Content\n\n")
	out.WriteString(md)
	out.WriteString("\n")

	dir := filepath.Join(outDir, p.Section)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, p.Slug+".md"), []byte(out.String()), 0o644)
}

func fetch(client *http.Client, url string) (*html.Node, error) {
	var lastErr error
	for attempt := 0; attempt < 4; attempt++ {
		req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
		req.Header.Set(headerUserAgent, userAgent)
		req.Header.Set(headerAcceptLanguage, acceptLanguageValue)
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(time.Duration(attempt+1) * time.Second)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			lastErr = fmt.Errorf("HTTP %d", resp.StatusCode)
			time.Sleep(time.Duration(attempt+1) * time.Second)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			continue
		}
		return html.Parse(strings.NewReader(string(b)))
	}
	return nil, lastErr
}

// findContent returns the div whose class contains "mesh-row__content",
// the article body container on MercadoLibre doc pages.
func findContent(n *html.Node) *html.Node {
	var found *html.Node
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if found != nil {
			return
		}
		if n.Type == html.ElementNode && n.DataAtom == atom.Div {
			if strings.Contains(classOf(n), "mesh-row__content") {
				found = n
				return
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return found
}

// ---- HTML -> Markdown ----------------------------------------------------

var wsRE = regexp.MustCompile(`\s+`)

func collapseWS(s string) string { return wsRE.ReplaceAllString(s, " ") }

// inlineMD renders inline content (text + inline elements) to Markdown.
func inlineMD(n *html.Node) string {
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.TextNode:
			sb.WriteString(collapseWS(c.Data))
		case html.ElementNode:
			switch c.DataAtom {
			case atom.A:
				txt := strings.TrimSpace(inlineMD(c))
				href := resolveHref(attr(c, "href"))
				if txt == "" {
					txt = href
				}
				if href == "" {
					sb.WriteString(txt)
				} else {
					sb.WriteString("[" + txt + "](" + href + ")")
				}
			case atom.Strong, atom.B:
				sb.WriteString("**" + strings.TrimSpace(inlineMD(c)) + "**")
			case atom.Em, atom.I:
				sb.WriteString("*" + strings.TrimSpace(inlineMD(c)) + "*")
			case atom.Code:
				sb.WriteString("`" + strings.TrimSpace(textContent(c)) + "`")
			case atom.Br:
				sb.WriteString("\n")
			case atom.Img:
				sb.WriteString("![" + attr(c, "alt") + "](" + attr(c, "src") + ")")
			default:
				sb.WriteString(inlineMD(c))
			}
		}
	}
	return sb.String()
}

// blockMD renders block-level content of n into sb.
func blockMD(n *html.Node, sb *strings.Builder) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.TextNode:
			if t := strings.TrimSpace(c.Data); t != "" {
				sb.WriteString("\n" + collapseWS(c.Data) + "\n")
			}
		case html.ElementNode:
			switch c.DataAtom {
			case atom.H1, atom.H2, atom.H3, atom.H4, atom.H5, atom.H6:
				level := int(c.Data[1] - '0')
				txt := strings.TrimSpace(inlineMD(c))
				if txt != "" {
					sb.WriteString("\n" + strings.Repeat("#", level) + " " + txt + "\n")
				}
			case atom.P:
				if txt := strings.TrimSpace(inlineMD(c)); txt != "" {
					sb.WriteString("\n" + txt + "\n")
				}
			case atom.Ul:
				sb.WriteString("\n")
				writeList(c, sb, false, 0)
			case atom.Ol:
				sb.WriteString("\n")
				writeList(c, sb, true, 0)
			case atom.Pre:
				writeCode(c, sb)
			case atom.Table:
				writeTable(c, sb)
			case atom.Blockquote:
				var inner strings.Builder
				blockMD(c, &inner)
				sb.WriteString("\n")
				for _, line := range strings.Split(strings.TrimSpace(inner.String()), "\n") {
					sb.WriteString("> " + line + "\n")
				}
			case atom.Hr:
				sb.WriteString("\n---\n")
			default:
				// Block-ish containers: recurse. Inline elements at block
				// level get wrapped as a paragraph.
				if isInline(c.DataAtom) {
					if txt := strings.TrimSpace(inlineMD(c)); txt != "" {
						sb.WriteString("\n" + txt + "\n")
					}
				} else {
					blockMD(c, sb)
				}
			}
		}
	}
}

func isInline(a atom.Atom) bool {
	switch a {
	case atom.A, atom.Strong, atom.B, atom.Em, atom.I, atom.Code, atom.Span, atom.Br, atom.Img:
		return true
	}
	return false
}

func writeList(list *html.Node, sb *strings.Builder, ordered bool, depth int) {
	idx := 0
	for li := list.FirstChild; li != nil; li = li.NextSibling {
		if li.Type != html.ElementNode || li.DataAtom != atom.Li {
			continue
		}
		idx++
		marker := "- "
		if ordered {
			marker = fmt.Sprintf("%d. ", idx)
		}
		indent := strings.Repeat("  ", depth)

		// Inline content of the li, excluding nested lists.
		var inline strings.Builder
		var nested []*html.Node
		for c := li.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && (c.DataAtom == atom.Ul || c.DataAtom == atom.Ol) {
				nested = append(nested, c)
				continue
			}
			inline.WriteString(inlineMD2(c))
		}
		text := strings.TrimSpace(collapseWS(inline.String()))
		sb.WriteString(indent + marker + text + "\n")
		for _, sub := range nested {
			writeList(sub, sb, sub.DataAtom == atom.Ol, depth+1)
		}
	}
}

// inlineMD2 renders a single node (used inside list items).
func inlineMD2(n *html.Node) string {
	wrap := &html.Node{Type: html.ElementNode}
	clone := *n
	clone.Parent, clone.PrevSibling, clone.NextSibling = nil, nil, nil
	wrap.FirstChild = &clone
	return inlineMD(wrap)
}

func writeCode(pre *html.Node, sb *strings.Builder) {
	lang := ""
	code := pre
	if c := firstElement(pre, atom.Code); c != nil {
		code = c
	}
	for _, cls := range strings.Fields(classOf(code) + " " + classOf(pre)) {
		if strings.HasPrefix(cls, "language-") {
			lang = strings.TrimPrefix(cls, "language-")
			break
		}
	}
	text := textContent(code)
	text = strings.Trim(text, "\n")
	sb.WriteString("\n```" + lang + "\n" + text + "\n```\n")
}

func writeTable(table *html.Node, sb *strings.Builder) {
	// MercadoLibre "resource reference" tables embed expandable modal panels
	// (div.modalmask) with per-method request/response examples. Render those
	// as structured endpoint sections instead of a cramped Markdown table.
	if hasClassDescendant(table, "modalmask") {
		writeResourceTable(table, sb)
		return
	}

	var rows [][]string
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.DataAtom == atom.Tr {
				var cells []string
				for cell := c.FirstChild; cell != nil; cell = cell.NextSibling {
					if cell.Type == html.ElementNode && (cell.DataAtom == atom.Td || cell.DataAtom == atom.Th) {
						v := strings.TrimSpace(collapseWS(inlineMD(cell)))
						v = strings.ReplaceAll(v, "|", "\\|")
						cells = append(cells, v)
					}
				}
				if len(cells) > 0 {
					rows = append(rows, cells)
				}
			} else {
				walk(c)
			}
		}
	}
	walk(table)
	if len(rows) == 0 {
		return
	}
	cols := 0
	for _, r := range rows {
		if len(r) > cols {
			cols = len(r)
		}
	}
	pad := func(r []string) []string {
		for len(r) < cols {
			r = append(r, "")
		}
		return r
	}
	sb.WriteString("\n")
	sb.WriteString("| " + strings.Join(pad(rows[0]), " | ") + " |\n")
	sb.WriteString("|" + strings.Repeat(" --- |", cols) + "\n")
	for _, r := range rows[1:] {
		sb.WriteString("| " + strings.Join(pad(r), " | ") + " |\n")
	}
	sb.WriteString("\n")
}

// writeResourceTable renders a MercadoLibre resource-reference table. Each body
// row becomes a "### `path`" section with the description and, for every modal
// panel, a "#### METHOD" block containing the request/response code samples.
func writeResourceTable(table *html.Node, sb *strings.Builder) {
	tbody := firstElement(table, atom.Tbody)
	if tbody == nil {
		tbody = table
	}
	for tr := tbody.FirstChild; tr != nil; tr = tr.NextSibling {
		if tr.Type != html.ElementNode || tr.DataAtom != atom.Tr {
			continue
		}
		var cells []*html.Node
		for td := tr.FirstChild; td != nil; td = td.NextSibling {
			if td.Type == html.ElementNode && (td.DataAtom == atom.Td || td.DataAtom == atom.Th) {
				cells = append(cells, td)
			}
		}
		if len(cells) == 0 {
			continue
		}
		path := strings.TrimSpace(collapseWS(textContent(cells[0])))
		if path == "" {
			continue
		}
		sb.WriteString("\n### `" + path + "`\n")
		if len(cells) > 1 {
			if desc := strings.TrimSpace(collapseWS(inlineMD(cells[1]))); desc != "" {
				sb.WriteString("\n" + desc + "\n")
			}
		}
		if len(cells) > 2 {
			writeModalPanels(cells[2], sb)
		}
	}
}

// writeModalPanels emits one "#### METHOD" block per div.modalmask found in the
// examples cell. The method label comes from the anchor that links to the
// panel's id (e.g. <a href="#modal1">GET</a> -> div id="modal1").
func writeModalPanels(cell *html.Node, sb *strings.Builder) {
	// Map panel id -> method label from the anchors in this cell.
	method := map[string]string{}
	var collectAnchors func(*html.Node)
	collectAnchors = func(n *html.Node) {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			if href := attr(n, "href"); strings.HasPrefix(href, "#") {
				method[strings.TrimPrefix(href, "#")] = strings.TrimSpace(textContent(n))
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			collectAnchors(c)
		}
	}
	collectAnchors(cell)

	var panels []*html.Node
	var collectPanels func(*html.Node)
	collectPanels = func(n *html.Node) {
		if n.Type == html.ElementNode && strings.Contains(classOf(n), "modalmask") {
			panels = append(panels, n)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			collectPanels(c)
		}
	}
	collectPanels(cell)

	for _, panel := range panels {
		m := method[attr(panel, "id")]
		if m == "" {
			m = "REQUEST"
		}
		sb.WriteString("\n#### " + strings.ToUpper(m) + "\n")
		// Walk the panel: h3 -> bold label, pre -> fenced code block. Skip the
		// go-back / close UI anchors.
		var walk func(*html.Node)
		walk = func(n *html.Node) {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type != html.ElementNode {
					continue
				}
				switch c.DataAtom {
				case atom.H1, atom.H2, atom.H3, atom.H4, atom.H5, atom.H6:
					if t := strings.TrimSpace(inlineMD(c)); t != "" {
						sb.WriteString("\n**" + t + "**\n")
					}
				case atom.Pre:
					writeCodeFence(c, sb)
				case atom.A:
					// skip UI controls
				default:
					walk(c)
				}
			}
		}
		walk(panel)
	}
}

// writeCodeFence writes a <pre> as a fenced code block, guessing the language
// from the content (curl -> bash, JSON -> json).
func writeCodeFence(pre *html.Node, sb *strings.Builder) {
	code := pre
	if c := firstElement(pre, atom.Code); c != nil {
		code = c
	}
	text := strings.Trim(textContent(code), "\n")
	sb.WriteString("\n```" + codeLang(text) + "\n" + text + "\n```\n")
}

func codeLang(text string) string {
	t := strings.TrimSpace(text)
	switch {
	case strings.HasPrefix(t, "$"), strings.HasPrefix(t, "curl"), strings.Contains(t, "curl -X"):
		return "bash"
	case strings.HasPrefix(t, "{"), strings.HasPrefix(t, "["):
		return "json"
	default:
		return ""
	}
}

func hasClassDescendant(n *html.Node, cls string) bool {
	if n.Type == html.ElementNode && strings.Contains(classOf(n), cls) {
		return true
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if hasClassDescendant(c, cls) {
			return true
		}
	}
	return false
}

// ---- helpers -------------------------------------------------------------

func extractEndpoints(content *html.Node) []string {
	text := textContent(content)
	re := regexp.MustCompile(`https?://api\.mercadolibre\.com[^\s"'` + "`" + `<>)\]]*`)
	set := map[string]bool{}
	for _, m := range re.FindAllString(text, -1) {
		m = strings.TrimRight(m, ".,;")
		set[m] = true
	}
	out := make([]string, 0, len(set))
	for e := range set {
		out = append(out, e)
	}
	sort.Strings(out)
	if len(out) > 60 {
		out = out[:60]
	}
	return out
}

func classOf(n *html.Node) string { return attr(n, "class") }

func attr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func firstElement(n *html.Node, a atom.Atom) *html.Node {
	var found *html.Node
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if found != nil {
			return
		}
		if n.Type == html.ElementNode && n.DataAtom == a {
			found = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return found
}

func textContent(n *html.Node) string {
	var sb strings.Builder
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode {
			sb.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return sb.String()
}

func resolveHref(href string) string {
	href = strings.TrimSpace(href)
	switch {
	case href == "" || strings.HasPrefix(href, "#"):
		return ""
	case strings.HasPrefix(href, "http://"), strings.HasPrefix(href, "https://"):
		return href
	case strings.HasPrefix(href, "//"):
		return "https:" + href
	case strings.HasPrefix(href, "/"):
		return siteHost + href
	default:
		return docsBase + href
	}
}

func yamlString(s string) string {
	if strings.ContainsAny(s, ":#\"'") {
		return `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
	}
	return s
}

var blankLinesRE = regexp.MustCompile(`\n{3,}`)

func cleanup(s string) string {
	s = blankLinesRE.ReplaceAllString(s, "\n\n")
	return strings.TrimSpace(s)
}
