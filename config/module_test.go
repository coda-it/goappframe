package config

import (
	"encoding/json"
	"testing"
)

func TestModuleUnmarshalProxy(t *testing.T) {
	moduleJSON := `{
		"id": "eshop",
		"properties": {"shopDomain": "https://shop.example.com"},
		"proxy": [
			{"method": "GET", "path": "/shops/{shopId}/products"},
			{"method": "POST", "path": "/shops/{shopId}/products", "protected": true, "injectAuth": true}
		]
	}`

	var moduleConfig Module
	if err := json.Unmarshal([]byte(moduleJSON), &moduleConfig); err != nil {
		t.Fatalf("unmarshalling module config failed: %v", err)
	}

	if len(moduleConfig.Proxy) != 2 {
		t.Fatalf("expected 2 proxy routes, got %d", len(moduleConfig.Proxy))
	}

	firstRoute := moduleConfig.Proxy[0]
	if firstRoute.Method != "GET" || firstRoute.Path != "/shops/{shopId}/products" {
		t.Errorf("unexpected first proxy route: %+v", firstRoute)
	}
	if firstRoute.Protected || firstRoute.InjectAuth {
		t.Errorf("proxy route flags should default to false: %+v", firstRoute)
	}

	secondRoute := moduleConfig.Proxy[1]
	if !secondRoute.Protected || !secondRoute.InjectAuth {
		t.Errorf("expected protected and injectAuth to be true: %+v", secondRoute)
	}
}

func TestModuleUnmarshalWithoutProxy(t *testing.T) {
	moduleJSON := `{"id": "post"}`

	var moduleConfig Module
	if err := json.Unmarshal([]byte(moduleJSON), &moduleConfig); err != nil {
		t.Fatalf("unmarshalling module config failed: %v", err)
	}

	if moduleConfig.Proxy != nil {
		t.Errorf("expected nil proxy routes when not configured, got %+v", moduleConfig.Proxy)
	}
}
