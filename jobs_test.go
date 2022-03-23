package winvps

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetJobs(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"jobs", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "jobs.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetJobs()
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}

func TestGetPendingJobs(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"jobs/pending", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "jobs.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetPendingJobs()
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}

func TestGetJob(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"jobs/1", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "job.json")
	})

	want := &Job{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}

	got, err := client.GetJob(1)
	require.NoError(t, err)
	require.Equal(t, want, got)

	got, err = client.GetJob(2)
	require.Error(t, err)
	require.Nil(t, got)
}

func TestCancelJob(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"jobs/1", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
	})

	err := client.CancelJob(1)
	require.NoError(t, err)
}
