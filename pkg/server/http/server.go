package http

import (
	"github.com/go-chi/chi"
	"github.com/owncloud/ocis-markdown-editor/pkg/assets"
	svc "github.com/owncloud/ocis-markdown-editor/pkg/service/v0"
	"github.com/owncloud/ocis-markdown-editor/pkg/version"
	"github.com/owncloud/ocis-pkg/v2/middleware"
	"github.com/owncloud/ocis-pkg/v2/service/http"
)

// Server initializes the http service and server.
func Server(opts ...Option) http.Service {
	options := newOptions(opts...)

	service := http.NewService(
		http.Logger(options.Logger),
		http.Name(options.Name),
		http.Namespace(options.Config.HTTP.Namespace),
		http.Version(version.String),
		http.Address(options.Config.HTTP.Addr),
		http.Context(options.Context),
		http.Flags(options.Flags...),
	)

	markdownEditor := svc.NewService()

	{
		markdownEditor = svc.NewInstrument(markdownEditor, options.Metrics)
		markdownEditor = svc.NewLogging(markdownEditor, options.Logger)
		markdownEditor = svc.NewTracing(markdownEditor)
	}

	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Cache)
	mux.Use(middleware.Cors)
	mux.Use(middleware.Secure)

	mux.Use(middleware.Version(
		options.Name,
		version.String,
	))

	mux.Use(middleware.Logger(
		options.Logger,
	))

	mux.Use(middleware.Static(
		options.Config.HTTP.Root,
		assets.New(
			assets.Logger(options.Logger),
			assets.Config(options.Config),
		),
	))

	service.Handle(
		"/",
		mux,
	)

	if err := service.Init(); err != nil {
		panic(err)
	}
	return service
}
