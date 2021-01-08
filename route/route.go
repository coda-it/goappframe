package route

import (
	"github.com/coda-it/gowebserver/router"
)

// Route - module route
type Route struct {
	Path      string
	Method    string
	Handler   router.ControllerHandler
	Protected bool
}
