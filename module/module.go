package module

import (
	"github.com/coda-it/goappframe/route"
)

// Module - application module
type Module struct {
	Enabled bool
	Routes  []route.Route
}
