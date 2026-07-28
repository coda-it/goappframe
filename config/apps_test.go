package config

import (
	"encoding/json"
	"testing"
)

func TestAppMatchesDomain(t *testing.T) {
	app := App{
		Domain:  "app.example.com",
		Domains: []string{"www.example.com", "example.com"},
	}

	matching := []string{"app.example.com", "www.example.com", "example.com"}
	for _, host := range matching {
		if !app.MatchesDomain(host) {
			t.Errorf("expected host %q to match app domains", host)
		}
	}

	if app.MatchesDomain("other.example.com") {
		t.Errorf("did not expect host %q to match app domains", "other.example.com")
	}
}

func TestAppMatchesDomainWithoutAdditionalDomains(t *testing.T) {
	app := App{Domain: "app.example.com"}

	if !app.MatchesDomain("app.example.com") {
		t.Errorf("expected primary domain to match")
	}
	if app.MatchesDomain("www.example.com") {
		t.Errorf("did not expect unlisted host to match")
	}
}

func TestAppUnmarshalDomains(t *testing.T) {
	appJSON := `{"id": "primary", "domain": "app.example.com", "domains": ["www.example.com", "example.com"]}`

	var app App
	if err := json.Unmarshal([]byte(appJSON), &app); err != nil {
		t.Fatalf("unmarshalling app config failed: %v", err)
	}

	if len(app.Domains) != 2 {
		t.Fatalf("expected 2 additional domains, got %d", len(app.Domains))
	}
	if !app.MatchesDomain("www.example.com") {
		t.Errorf("expected unmarshalled domains to be matchable")
	}
}
