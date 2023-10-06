package repository

import "github.com/notEpsilon/shorty/pkg/types"

type ShortyRepository interface {
	SaveUrl(url *types.ShortUrl) error
	FindUrlBySlug(slug string) (*types.ShortUrl, error)
	FindUrlByLink(url string) (*types.ShortUrl, error)
}

// ----------------------------- In Memory Repository -----------------------------

type InMemoryRepository struct {
	Urls map[string]*types.ShortUrl
}

func (r *InMemoryRepository) SaveUrl(url *types.ShortUrl) error {
	r.Urls[url.Slug] = url
	return nil
}

func (r *InMemoryRepository) FindUrlBySlug(slug string) (*types.ShortUrl, error) {
	return r.Urls[slug], nil
}

func (r *InMemoryRepository) FindUrlByLink(url string) (*types.ShortUrl, error) {
	for k, v := range r.Urls {
		if v.RedirectTo == url {
			return r.Urls[k], nil
		}
	}
	return nil, nil
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		Urls: make(map[string]*types.ShortUrl),
	}
}

var _ ShortyRepository = (*InMemoryRepository)(nil)

// ----------------------------- In Memory Repository -----------------------------
