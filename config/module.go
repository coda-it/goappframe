package config

// ProxyRoute - a single upstream endpoint exposed through a module's reverse
// proxy; Method and Path describe the proxied request (Path may contain
// {param} segments), Protected requires a logged-in session and InjectAuth
// lets the application replace the Authorization header with its own token
type ProxyRoute struct {
	Method     string `json:"method"`
	Path       string `json:"path"`
	Protected  bool   `json:"protected"`
	InjectAuth bool   `json:"injectAuth"`
}

// Module - module config
type Module struct {
	ID         string            `json:"id"`
	Properties map[string]string `json:"properties"`
	Proxy      []ProxyRoute      `json:"proxy,omitempty"`
}
