package winvps

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLocations(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"locations", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "locations.json")
	})

	want := []*Location{{ID: 1, Name: "test"}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetLocations()
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}
