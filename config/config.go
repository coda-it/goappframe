package config

// Config - static app config
type Config struct {
	Apps            []App  `json:"apps"`
	DefaultLanguage string `json:"defaultLanguage"`
	Cert            string `json:"certificate"`
	CertPrvKey      string `json:"certificatePrivateKey"`
}
