package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-test/deep"
	"github.com/matheusmhmelo/go-api-tdd/internal/pkg/credential"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubService func(ctx context.Context, receiverID string, c credential.Credentials) error

func (s stubService) Save(ctx context.Context, receiverID string, c credential.Credentials) error {
	return s(ctx, receiverID, c)
}

func TestCredential_Create(t *testing.T) {
	want := http.StatusCreated

	creds := credential.Credentials{
		ClientID:            "1",
		OrganizationID:      "2",
		SoftwareStatementID: "3",
	}
	receiverID := "abc"

	s := stubService(func(_ context.Context, id string, c credential.Credentials) error {
		if diff := deep.Equal(c, creds); diff != nil {
			t.Error("unexpected credential received on service: ", diff)
		}

		if id != receiverID {
			t.Error("unexpected receiverID received on service: ", id)
		}
		return nil
	})
	c := NewCredential(&s)

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/credentials", getBody(creds))
	if err != nil {
		t.Fatal("error to create request POST - /credentials")
	}
	req.Header.Set(partnerHeader, receiverID)

	c.Create(respRec, req)

	got := respRec.Result()

	if got.StatusCode != want {
		t.Errorf("Create got unexpected status code: %v", got.StatusCode)
	}
}

func TestCredential_Create_Error(t *testing.T) {
	tests := []struct {
		name       string
		getReq     func(t *testing.T) *http.Request
		service    stubService
		statusCode int
	}{
		{
			name: "EmptyPartnerID",
			getReq: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/credentials", getBody(credential.Credentials{
					ClientID:            "1",
					OrganizationID:      "2",
					SoftwareStatementID: "3",
				}))
				if err != nil {
					t.Fatal("error to create request POST - /credentials")
				}
				return req
			},
			service: stubService(func(_ context.Context, id string, c credential.Credentials) error {
				return nil
			}),
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Wrong Body",
			getReq: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/credentials", bytes.NewBuffer([]byte("test")))
				if err != nil {
					t.Fatal("error to create request POST - /credentials")
				}
				req.Header.Set(partnerHeader, "abc")
				return req
			},
			service: stubService(func(_ context.Context, id string, c credential.Credentials) error {
				return nil
			}),
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Error On Service",
			getReq: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/credentials", getBody(credential.Credentials{
					ClientID:            "1",
					OrganizationID:      "2",
					SoftwareStatementID: "3",
				}))
				if err != nil {
					t.Fatal("error to create request POST - /credentials")
				}
				req.Header.Set(partnerHeader, "abc")
				return req
			},
			service: stubService(func(_ context.Context, id string, c credential.Credentials) error {
				return errors.New("error")
			}),
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid Credentials",
			getReq: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/credentials", getBody(credential.Credentials{
					ClientID:       "1",
					OrganizationID: "2",
				}))
				if err != nil {
					t.Fatal("error to create request POST - /credentials")
				}
				req.Header.Set(partnerHeader, "abc")
				return req
			},
			service: stubService(func(_ context.Context, id string, c credential.Credentials) error {
				return nil
			}),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respRec := httptest.NewRecorder()

			c := NewCredential(tt.service)
			c.Create(respRec, tt.getReq(t))

			got := respRec.Result()

			if got.StatusCode != tt.statusCode {
				t.Errorf("Create got unexpected status code: %v", got.StatusCode)
			}
		})
	}
}

func getBody(credentials credential.Credentials) io.Reader {
	j, _ := json.Marshal(credentials)
	return bytes.NewBuffer(j)
}
