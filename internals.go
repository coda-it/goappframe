package goappframe

import (
	"github.com/coda-it/goappframe/module"
	"github.com/coda-it/goutils/mailer"
	"github.com/coda-it/gowebserver/router"
)

// Internals - application internals
type Internals struct {
	Port        string
	Modules     []module.Module
	Persistence interface{}
	DataKey     string
	Mailer      mailer.IMailer
	NotFound    router.ControllerHandler
}
