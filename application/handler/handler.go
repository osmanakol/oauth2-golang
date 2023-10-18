package handler

import (
	"github.com/osmanakol/oauth2-golang/pkg/google"
)

type Handler struct {
	GoogleRoute GoogleHandler
}

type HandlerConfigurations func(h *Handler) error

func NewHandler(configs ...HandlerConfigurations) (*Handler, error) {
	h := &Handler{}

	for _, cfg := range configs {
		err := cfg(h)

		if err != nil {
			return nil, err
		}
	}

	return h, nil
}

func WithGoogleRoute() HandlerConfigurations {
	return func(h *Handler) error {
		g := google.NewGoogleOauth2()

		h.GoogleRoute = NewGoogleHandler(&g)
		return nil
	}
}
