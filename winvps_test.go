package winvps

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// setup a test http server
func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client, err := NewClient("secret", BaseURL(server.URL))
	if err != nil {
		server.Close()
		t.Fatalf("Failed to create a new client: %v", err)
	}
	return mux, server, client
}

// close the httptest server
func teardown(s *httptest.Server) {
	s.Close()
}

// returns body from http.Request as string
func getBody(t *testing.T, r *http.Request) string {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	if err != nil {
		t.Fatalf("Failed to Read Body: %v", err)
	}
	return buffer.String()
}

// write fixture to specified io.Writer
func writeFixture(t *testing.T, w io.Writer, fixturePath string) {
	f, err := os.Open("testdata/" + fixturePath)
	if err != nil {
		t.Fatalf("failed to open fixture file: %v", err)
	}

	if _, err = io.Copy(w, f); err != nil {
		t.Fatalf("failed to copy fixture to writer: %v", err)
	}
}

func TestRequestOptions(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "limit=1&page=2", r.URL.RawQuery)
		writeFixture(t, w, "machines.json")
	})

	_, _, err := client.GetMachines(&RequestOptions{Limit: 1, Page: 2})
	if err != nil {
		t.Fatal(err)
	}
}
func TestRequestHeaders(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "secret", r.Header.Get("API-KEY"))
		require.Equal(t, "go-winvps", r.Header.Get("User-Agent"))
		require.Equal(t, "application/json", r.Header.Get("Accept"))
		writeFixture(t, w, "machines.json")
	})

	_, _, err := client.GetMachines()
	if err != nil {
		t.Fatal(err)
	}
}
