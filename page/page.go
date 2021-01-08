package page

import "github.com/coda-it/goappframe/navigation"

// Page - entity representing page
type Page struct {
	Version    string
	Title      string
	IsLogged   bool
	IsRoot     bool
	Params     map[string]interface{}
	Name       string
	Navigation []navigation.Navigation
	JSConfig   string
}
