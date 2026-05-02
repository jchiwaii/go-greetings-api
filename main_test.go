package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	response := performRequest(http.MethodGet, "/", nil)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	body := response.Body.String()
	if !strings.Contains(body, "Go Greeting API") {
		t.Fatalf("expected welcome text, got %q", body)
	}
}

func TestGetGreetingWithName(t *testing.T) {
	response := performRequest(http.MethodGet, "/greet?name=Amina", nil)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var payload greetingResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if payload.Name != "Amina" {
		t.Fatalf("expected name Amina, got %q", payload.Name)
	}

	if payload.Message != "Hello, Amina! Welcome to Go." {
		t.Fatalf("unexpected message: %q", payload.Message)
	}
}

func TestGetGreetingDefaultsToWorld(t *testing.T) {
	response := performRequest(http.MethodGet, "/greet", nil)

	var payload greetingResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if payload.Name != "World" {
		t.Fatalf("expected default name World, got %q", payload.Name)
	}
}

func TestPostGreetingTrimsExtraSpaces(t *testing.T) {
	body := bytes.NewBufferString(`{"name":"  Grace   Hopper  "}`)
	response := performRequest(http.MethodPost, "/greet", body)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var payload greetingResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if payload.Name != "Grace Hopper" {
		t.Fatalf("expected normalized name, got %q", payload.Name)
	}
}

func TestGreetingRejectsInvalidJSON(t *testing.T) {
	response := performRequest(http.MethodPost, "/greet", strings.NewReader(`{bad json}`))

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}
}

func TestGreetingRejectsUnsupportedMethod(t *testing.T) {
	response := performRequest(http.MethodDelete, "/greet", nil)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", response.Code)
	}
}

func performRequest(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, body)
	response := httptest.NewRecorder()
	newRouter().ServeHTTP(response, request)
	return response
}
