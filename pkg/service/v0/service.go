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
func NewService(opts ...Option) Service {
	options := newOptions(opts...)

	m := chi.NewMux()
	m.Use(options.Middleware...)

	svc := MarkdownEditor{
		config: options.Config,
		mux:    m,
	}

	return svc
}

// MarkdownEditor defines implements the business logic for Service.
type MarkdownEditor struct {
	config *config.Config
	mux    *chi.Mux
}

// ServeHTTP implements the Service interface.
func (g MarkdownEditor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.mux.ServeHTTP(w, r)
}
