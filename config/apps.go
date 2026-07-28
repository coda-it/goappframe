package config

import "github.com/coda-it/goappframe/navigation"

// App - application config
type App struct {
	ID string `json:"id"`
	// Domain is the primary/canonical domain, used when building outbound URLs
	// (activation e-mails, redirects, etc.).
	Domain string `json:"domain"`
	// Domains lists additional domains that also route to this app. Requests
	// whose host matches Domain or any entry here resolve to this app.
	Domains    []string                `json:"domains"`
	Modules    []Module                `json:"modules"`
	Navigation []navigation.Navigation `json:"navigation"`
}

// MatchesDomain - reports whether the given host resolves to this app. A host
// matches when it equals the primary Domain or any of the additional Domains.
func (a App) MatchesDomain(host string) bool {
	if a.Domain == host {
		return true
	}
	for _, domain := range a.Domains {
		if domain == host {
			return true
		}
	}
	return false
}
