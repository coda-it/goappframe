package config

import (
	"github.com/coda-it/goappframe/navigation"
)

// Config - static app config
type Config struct {
	Navigation []navigation.Navigation `json:"navigation"`
}
