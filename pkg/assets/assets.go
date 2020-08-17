package assets

import (
	"net/http"
	"os"

	"github.com/owncloud/ocis-markdown-editor/pkg/config"
	"github.com/owncloud/ocis-pkg/v2/log"

	// Fake the import to make the dep tree happy.
	_ "golang.org/x/net/context"

	// Fake the import to make the dep tree happy.
	_ "golang.org/x/net/webdav"
)

//go:generate go run github.com/UnnoTed/fileb0x embed.yml

// assets gets initialized by New and provides the handler.
type assets struct {
	logger log.Logger
	config *config.Config
}

// Open just implements the HTTP filesystem interface.
func (a assets) Open(original string) (http.File, error) {
	return FS.OpenFile(
		CTX,
		original,
		os.O_RDONLY,
		0644,
	)
}

// New returns a new http filesystem to serve assets.
func New(opts ...Option) http.FileSystem {
	options := newOptions(opts...)

	return assets{
		config: options.Config,
	}
}
