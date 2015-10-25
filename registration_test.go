package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFormDisplay(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	reg := &registration{}
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusOK))
	}

	if !strings.Contains(resp.Body.String(), "<form") {
		t.Fatal("Registraton dose not caontain form element!")
	}
}
