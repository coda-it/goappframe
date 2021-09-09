package config

import (
	"github.com/coda-it/goappframe/navigation"
)

// Config - static app config
type Config struct {
	Navigation      []navigation.Navigation `json:"navigation"`
	Apps            []App                   `json:"apps"`
	DefaultLanguage string                  `json:"defaultLanguage"`
	Cert            string                  `json:"certificate`
	CertPrvKey      string                  `json:"certificatePrivateKey`
}
