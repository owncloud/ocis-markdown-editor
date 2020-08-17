package http

import (
	"github.com/owncloud/ocis-markdown-editor/pkg/assets"
	gohttp "net/http"
	"github.com/go-chi/chi"
	"github.com/owncloud/ocis-markdown-editor/pkg/version"
	"github.com/owncloud/ocis-pkg/v2/service/http"
	"path"
	"strings"
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

	mux := chi.NewMux()
	mux.Use(Static(
			options.Config.HTTP.Root,
			assets.New(
				assets.Logger(options.Logger),
				assets.Config(options.Config),
			),
		))
	mux.Get("/", gohttp.HandlerFunc(func(writer gohttp.ResponseWriter, request *gohttp.Request) {
		writer.Write([]byte("noop"))
	}))

	service.Handle("/", mux)

	if err := service.Init(); err != nil {
		panic(err)
	}
	return service
}

// Static is a middleware that serves static assets.
func Static(root string, fs gohttp.FileSystem) func(gohttp.Handler) gohttp.Handler {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}

	static := gohttp.StripPrefix(
		root,
		gohttp.FileServer(
			fs,
		),
	)

	return func(next gohttp.Handler) gohttp.Handler {
		return gohttp.HandlerFunc(func(w gohttp.ResponseWriter, r *gohttp.Request) {
			if strings.HasPrefix(r.URL.Path, path.Join(root, "api")) {
				next.ServeHTTP(w, r)
			} else {
				if strings.HasSuffix(r.URL.Path, "/") {
					gohttp.NotFound(w, r)
				} else {
					static.ServeHTTP(w, r)
				}
			}
		})
	}
}