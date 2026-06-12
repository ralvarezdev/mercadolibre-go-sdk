package mercadolibre

import "context"

// PageIterator walks offset/limit pages of any list endpoint. Create one with
// NewPageIterator.
//
// Unlike ItemScanIterator (which uses ML's scan/scroll protocol), PageIterator
// uses ordinary offset/limit pagination and works with any endpoint that
// returns a Paging field.
type PageIterator[T any] struct {
	fetch func(ctx context.Context, opts ListOptions) ([]T, Paging, error)
	opts  ListOptions
	done  bool
}

// NewPageIterator creates an iterator whose pages are produced by fetch.
//
//	iter := mercadolibre.NewPageIterator(
//	    mercadolibre.ListOptions{Limit: 50},
//	    func(ctx context.Context, o mercadolibre.ListOptions) ([]mercadolibre.Order, mercadolibre.Paging, error) {
//	        resp, err := client.Orders.SearchBySeller(ctx, sellerID, o, nil)
//	        if err != nil {
//	            return nil, mercadolibre.Paging{}, err
//	        }
//	        return resp.Results, resp.Paging, nil
//	    },
//	)
//	for {
//	    orders, ok, err := iter.Next(ctx)
//	    if err != nil { log.Fatal(err) }
//	    if !ok { break }
//	    for _, o := range orders { fmt.Println(o.ID) }
//	}
func NewPageIterator[T any](
	opts ListOptions,
	fetch func(context.Context, ListOptions) ([]T, Paging, error),
) *PageIterator[T] {
	return &PageIterator[T]{fetch: fetch, opts: opts}
}

// Next returns the next page of items.
// ok is false when iteration is complete (empty page or all pages exhausted).
func (it *PageIterator[T]) Next(ctx context.Context) (items []T, ok bool, err error) {
	if it.done {
		return nil, false, nil
	}
	items, paging, err := it.fetch(ctx, it.opts)
	if err != nil {
		return nil, false, err
	}
	if len(items) == 0 {
		it.done = true
		return nil, false, nil
	}
	it.opts.Offset = paging.Offset + len(items)
	// mark done when we got a partial page or exhausted the total
	if it.opts.Limit > 0 && len(items) < it.opts.Limit {
		it.done = true
	}
	if paging.Total > 0 && it.opts.Offset >= paging.Total {
		it.done = true
	}
	return items, true, nil
}
