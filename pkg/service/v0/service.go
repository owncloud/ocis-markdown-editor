package svc

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/owncloud/ocis-markdown-editor/pkg/config"
)

// Service defines the extension handlers.
type Service interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// NewService returns a service implementation for Service.
func NewService(cfg *config.Config) Service {
	m := chi.NewMux()

	svc := MarkdownEditor{
		config: cfg,
		mux:    m,
	}

	return svc
}

// MarkdownEditor defines implements the business logic for Service.
type MarkdownEditor struct {
	config *config.Config
	mux    *chi.Mux
}

// AddMiddleware adds a middleware to the service
func (g MarkdownEditor) AddMiddleware() {
}

// ServeHTTP implements the Service interface.
func (g MarkdownEditor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.mux.ServeHTTP(w, r)
}
