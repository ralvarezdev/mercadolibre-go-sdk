package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// TokenSource supplies a valid bearer access token, refreshing it as needed.
type TokenSource interface {
	// Token returns a currently-valid access token string.
	Token(ctx context.Context) (string, error)
}

// TokenSourceFunc adapts a function to a TokenSource.
type TokenSourceFunc func(ctx context.Context) (string, error)

func (f TokenSourceFunc) Token(ctx context.Context) (string, error) { return f(ctx) }

// StaticTokenSource returns a TokenSource that always returns the same access
// token and never refreshes. Useful for scripts and tests.
func StaticTokenSource(accessToken string) TokenSource {
	return TokenSourceFunc(func(context.Context) (string, error) { return accessToken, nil })
}

// Store persists a Token across restarts. Because MercadoLibre refresh tokens
// are single-use and rotate on every refresh, a Store is required for long-lived
// apps so the latest refresh token is never lost.
type Store interface {
	Load(ctx context.Context) (*Token, error)
	Save(ctx context.Context, t *Token) error
}

// MemoryStore is an in-process Store. Not durable across restarts.
type MemoryStore struct {
	tok *Token
	mu  sync.Mutex
}

// NewMemoryStore returns a MemoryStore optionally seeded with an initial token.
func NewMemoryStore(initial *Token) *MemoryStore { return &MemoryStore{tok: initial} }

func (m *MemoryStore) Load(context.Context) (*Token, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.tok == nil {
		return (*Token)(nil), nil
	}
	cp := *m.tok
	return &cp, nil
}

func (m *MemoryStore) Save(_ context.Context, t *Token) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	cp := *t
	m.tok = &cp
	return nil
}

// FileStore persists the token as JSON at Path, writing atomically.
type FileStore struct {
	Path string
	mu   sync.Mutex
}

// NewFileStore returns a FileStore backed by the given file path.
func NewFileStore(path string) *FileStore { return &FileStore{Path: path} }

func (f *FileStore) Load(context.Context) (*Token, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	b, err := os.ReadFile(f.Path)
	if err != nil {
		if os.IsNotExist(err) {
			var t *Token
			return t, nil
		}
		return nil, err
	}
	var t Token
	if err := json.Unmarshal(b, &t); err != nil {
		return nil, fmt.Errorf("mercadolibre/auth: decoding token file: %w", err)
	}
	return &t, nil
}

func (f *FileStore) Save(_ context.Context, t *Token) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	dir := filepath.Dir(f.Path)
	tmp, err := os.CreateTemp(dir, ".token-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if _, err := tmp.Write(b); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, f.Path)
}

// RefreshingTokenSource returns valid access tokens, auto-refreshing when the
// current token is near expiry and persisting the rotated token via Store before
// returning it. Refreshes are serialized (single-flight) so concurrent callers
// never burn a single-use refresh token twice.
type RefreshingTokenSource struct {
	store  Store
	tok    *Token
	cfg    Config
	leeway time.Duration
	mu     sync.Mutex
}

// NewRefreshingTokenSource builds a refreshing source. If initial is nil it is
// loaded lazily from store on first use. leeway is how early to refresh before
// expiry (use 0 for a sensible default of 5 minutes).
func NewRefreshingTokenSource(cfg Config, initial *Token, store Store, leeway time.Duration) *RefreshingTokenSource {
	if leeway <= 0 {
		leeway = 5 * time.Minute
	}
	return &RefreshingTokenSource{cfg: cfg, store: store, leeway: leeway, tok: initial}
}

func (s *RefreshingTokenSource) Token(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tok == nil && s.store != nil {
		t, err := s.store.Load(ctx)
		if err != nil {
			return "", err
		}
		s.tok = t
	}
	if s.tok == nil {
		return "", fmt.Errorf("mercadolibre/auth: no token available; complete the authorization flow first")
	}
	if s.tok.Valid(s.leeway) {
		return s.tok.AccessToken, nil
	}
	if s.tok.RefreshToken == "" {
		return "", fmt.Errorf("mercadolibre/auth: token expired and no refresh token; re-authorize")
	}

	nt, err := s.cfg.Refresh(ctx, s.tok.RefreshToken)
	if err != nil {
		return "", err
	}
	// Persist the rotated token BEFORE returning, so a crash never loses it.
	if s.store != nil {
		if err := s.store.Save(ctx, nt); err != nil {
			return "", fmt.Errorf("mercadolibre/auth: persisting refreshed token: %w", err)
		}
	}
	s.tok = nt
	return nt.AccessToken, nil
}
