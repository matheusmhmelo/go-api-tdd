package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubCredential func(w http.ResponseWriter, r *http.Request)

func (c stubCredential) Create(w http.ResponseWriter, r *http.Request) {
	c(w, r)
}

func TestNewRouter(t *testing.T) {
	want := http.StatusOK

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/credentials", nil)
	if err != nil {
		t.Fatal("error to create request POST - /credentials")
	}

	cred := stubCredential(func(_ http.ResponseWriter, _ *http.Request) {})
	router := NewRouter(cred)

	router.ServeHTTP(respRec, req)
	got := respRec.Result()

	if got.StatusCode != want {
		t.Errorf("router.ServeHTTP got unexpected status code: %v", got.StatusCode)
	}
}

func TestNewRouter_Error(t *testing.T) {
	t.Run("Method not allowed", func(t *testing.T) {
		want := http.StatusMethodNotAllowed

		respRec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/credentials", nil)
		if err != nil {
			t.Fatal("error to create request GET - /credentials")
		}

		cred := stubCredential(func(_ http.ResponseWriter, _ *http.Request) {})
		router := NewRouter(cred)

		router.ServeHTTP(respRec, req)
		got := respRec.Result()

		if got.StatusCode != want {
			t.Errorf("router.ServeHTTP got unexpected status code: %v", got.StatusCode)
		}
	})
	t.Run("Error on Handler", func(t *testing.T) {
		want := http.StatusBadRequest

		respRec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/credentials", nil)
		if err != nil {
			t.Fatal("error to create request POST - /credentials")
		}

		cred := stubCredential(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		})
		router := NewRouter(cred)

		router.ServeHTTP(respRec, req)
		got := respRec.Result()

		if got.StatusCode != want {
			t.Errorf("router.ServeHTTP got unexpected status code: %v", got.StatusCode)
		}
	})
}
