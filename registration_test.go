package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type DummyStore struct {
	storeMethod func(reg Registration) (err error)
}

func (s DummyStore) Store(reg Registration) (err error) {
	return s.storeMethod(reg)
}

func NewDummyReg() *Registration {
	dummyStore := DummyStore{func(reg Registration) (err error) { return nil }}

	return &Registration{
		"",
		"",
		"",
		time.Now(),
		dummyStore,
		""}
}

func TestFormDisplay(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	reg := NewDummyReg()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Server replied with status %v expected %v", resp.Code, http.StatusOK)
	}

	if !strings.Contains(resp.Body.String(), "<form") {
		t.Fatal("Registration does not contain form element!")
	}
}

func TestFormProcessBadRequest(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	reg := NewDummyReg()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("Server replied with status %v expected %v", resp.Code, http.StatusBadRequest)
	}

	// test bad mail format
	resp = httptest.NewRecorder()

	req, err = http.NewRequest("POST", "/", strings.NewReader("name=Test&surname=Test&email=Test"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	reg = NewDummyReg()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("Server replied with status %v expected %v", resp.Code, http.StatusBadRequest)
	}
}

func TestFormProcessOkRequest(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/", strings.NewReader("name=Test&surname=Test&email=test@test.com"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	reg := NewDummyReg()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Server replied with status %v expected %v", resp.Code, http.StatusOK)
	}
}

func TestStoreCall(t *testing.T) {
	storeCalled := false
	defer func() {
		if storeCalled == false {
			t.Fatal("Store not called!")
		}
	}()

	storeChecker := func(reg Registration) (err error) {
		storeCalled = true
		return nil
	}

	dummyStore := DummyStore{storeChecker}

	reg := &Registration{Storer: dummyStore}

	req, err := http.NewRequest("POST", "/", strings.NewReader("name=Test&surname=Test&email=test@test.com"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp := httptest.NewRecorder()
	reg.ServeHTTP(resp, req)
}

func TestStoreError(t *testing.T) {
	storeFailer := func(reg Registration) (err error) {
		return errors.New("Store failed!")
	}

	dummyStore := DummyStore{storeFailer}

	reg := &Registration{Storer: dummyStore}

	req, err := http.NewRequest("POST", "/", strings.NewReader("name=Test&surname=Test&email=test@test.com"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp := httptest.NewRecorder()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusInternalServerError {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusInternalServerError))
	}
}

func TestSuccessMsgDisplay(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/", strings.NewReader("name=Test&surname=Test&email=test@test.com"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	reg := NewDummyReg()
	reg.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("Server repled with status %v expected %v", resp.Code, http.StatusOK))
	}

	if !strings.Contains(resp.Body.String(), registrationSuccessful) {
		t.Fatal(fmt.Sprintf("Registraton dose not contain success message, expected \"%s\"", registrationSuccessful))
	}
}
