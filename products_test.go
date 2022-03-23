package winvps

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetProducts(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"products", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "products.json")
	})

	want := []*Product{{ID: 1, Name: "test", Limits: &Limits{CpuPercent: 100, CpuCores: 1, RamMin: 1024, RamMax: 1024, DiskSize: 30, Bandwidth: 10, Traffic: 1}}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetProducts()
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}
