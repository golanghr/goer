package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type RegWithMockStore struct {
	Registration
	storeMethod func()
}

func (r RegWithMockStore) Store() {
	r.storeMethod()
}

func NewDummyRegWithMockStore() RegWithMockStore {
	return RegWithMockStore{
		Registration{"", "", "", time.Now()},
		func() {}}
}

func TestFormDisplay(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	reg := NewDummyRegWithMockStore()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusOK))
	}

	if !strings.Contains(resp.Body.String(), "<form") {
		t.Fatal("Registraton dose not caontain form element!")
	}
}

func TestFormProcessBadRequest(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	reg := NewDummyRegWithMockStore()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusBadRequest))
	}

	// test bad mail format
	resp = httptest.NewRecorder()

	req, err = http.NewRequest("POST", "/", strings.NewReader("name=Test&surname=Test&email=Test"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	reg = NewDummyRegWithMockStore()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusBadRequest))
	}
}

func TestFormProcessOkRequest(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/", strings.NewReader("name=Test&surname=Test&email=test@test.com"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	reg := NewDummyRegWithMockStore()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusOK))
	}
}

func TestStoreCall(t *testing.T) {
	storeCalled := false
	defer func() {
		if storeCalled == false {
			t.Fatal(fmt.Sprintf("Store not called!"))
		}
	}()

	dummyReg := NewDummyRegWithMockStore()
	dummyReg.storeMethod = func() {
		storeCalled = true
	}

	reg := dummyReg

	req, err := http.NewRequest("POST", "/", strings.NewReader("name=Test&surname=Test&email=test@test.com"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp := httptest.NewRecorder()
	reg.ServeHTTP(resp, req)
}
