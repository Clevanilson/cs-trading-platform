package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpAdapterRoutes(t *testing.T) {
	adapter := NewHttpAdapter()

	adapter.POST("/signup", func(input *HandlerInput) (*Response, error) {
		return &Response{StatusCode: http.StatusCreated, Body: map[string]string{"status": "created"}}, nil
	})
	adapter.GET("/get_account/:id", func(input *HandlerInput) (*Response, error) {
		return &Response{StatusCode: http.StatusOK, Body: map[string]string{"id": input.Params["id"]}}, nil
	})

	ts := httptest.NewServer(adapter.Handler())
	defer ts.Close()

	t.Run("POST signup", func(t *testing.T) {
		res, err := http.Post(ts.URL+"/signup", "application/json", nil)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("expected 201, got %d", res.StatusCode)
		}
	})

	t.Run("GET account by id", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/get_account/abc-123")
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected 200, got %d", res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		var payload map[string]string
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatal(err)
		}
		if payload["id"] != "abc-123" {
			t.Fatalf("expected id abc-123, got %q", payload["id"])
		}
	})
}
