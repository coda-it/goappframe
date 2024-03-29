package config

// App - application config
type App struct {
	ID      string   `json:"id"`
	Modules []Module `json:"modules"`
	Domain  string   `json:"domain"`
}
