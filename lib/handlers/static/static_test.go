package static

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileServe(t *testing.T) {
	resp := httptest.NewRecorder()

	uri := "/static_test.go"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	staticHandler := Handler("./")

	staticHandler.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusOK))
	}
}

func TestFileServe404(t *testing.T) {
	resp := httptest.NewRecorder()

	uri := "/not_existing"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	staticHandler := Handler("./")

	staticHandler.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusNotFound))
	}
}
