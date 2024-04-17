package config

import "github.com/coda-it/goappframe/navigation"

// App - application config
type App struct {
	ID         string                  `json:"id"`
	Modules    []Module                `json:"modules"`
	Domain     string                  `json:"domain"`
	Navigation []navigation.Navigation `json:"navigation"`
}
