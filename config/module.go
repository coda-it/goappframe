package config

// Module - module config
type Module struct {
	ID         string            `json:"id"`
	Properties map[string]string `json:"properties"`
}
