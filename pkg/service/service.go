package service

import (
	"net/http"

	"github.com/notEpsilon/shorty/pkg/repository"
	"github.com/notEpsilon/shorty/pkg/types"
	"github.com/notEpsilon/shorty/pkg/util"
)

type ShortyService interface {
	Shorten(url string) (*types.ShortUrl, error)
	Redirect(slug string, w http.ResponseWriter, r *http.Request) error
}

// ----------------------------- Service Impl -----------------------------

type ServiceImpl struct {
	repository.ShortyRepository
}

func (s *ServiceImpl) Shorten(url string) (*types.ShortUrl, error) {
	if val, _ := s.ShortyRepository.FindUrlByLink(url); val != nil {
		return val, nil
	}

	resUrl := &types.ShortUrl{Slug: util.GetRandomSlug(), RedirectTo: url}
	if err := s.ShortyRepository.SaveUrl(resUrl); err != nil {
		return nil, err
	}

	return resUrl, nil
}

func (s *ServiceImpl) Redirect(slug string, w http.ResponseWriter, r *http.Request) error {
	redirectTo, err := s.ShortyRepository.FindUrlBySlug(slug)
	if err != nil {
		return err
	}
	http.Redirect(w, r, redirectTo.RedirectTo, http.StatusSeeOther)
	return nil
}

func NewServiceImpl(repo repository.ShortyRepository) *ServiceImpl {
	return &ServiceImpl{
		ShortyRepository: repo,
	}
}

var _ ShortyService = (*ServiceImpl)(nil)

// ----------------------------- Service Impl -----------------------------
