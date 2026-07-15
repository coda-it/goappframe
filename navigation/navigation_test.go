package navigation

import (
	"encoding/json"
	"testing"
)

func TestNavigationSerializesIsAdmin(t *testing.T) {
	serialized, err := json.Marshal(Navigation{
		ID:      "admin-panel",
		Label:   "Admin panel",
		Href:    "/admin",
		IsAdmin: true,
	})
	if err != nil {
		t.Fatalf("marshaling navigation failed: %v", err)
	}

	var deserialized map[string]interface{}
	if err := json.Unmarshal(serialized, &deserialized); err != nil {
		t.Fatalf("unmarshaling navigation failed: %v", err)
	}

	isAdmin, ok := deserialized["isAdmin"]
	if !ok {
		t.Fatal("expected isAdmin key in serialized navigation")
	}
	if isAdmin != true {
		t.Fatalf("expected isAdmin to be true, got %v", isAdmin)
	}
}
