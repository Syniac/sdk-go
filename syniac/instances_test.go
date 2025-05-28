package syniac

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func TestListInstances(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/instances" {
			t.Fatalf("expected /v1/instances, got %s", r.URL.Path)
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode([]Instance{{ID: "123", Name: "Test", Status: "running"}})
	}))
	defer server.Close()

	client := NewClient("fake-token")
	client.BaseURL = server.URL

	instances, err := client.ListInstances()
	if err != nil {
		t.Fatal(err)
	}

	if len(instances) != 1 || instances[0].ID != "123" {
		t.Fatalf("unexpected response: %+v", instances)
	}
}
