package winvps

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMachines(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "machines.json")
	})

	want := []*Machine{{Name: "VPS0123", Status: "Running"}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetMachines()
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}

func TestGetMachinesFull(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/full", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "machinesfull.json")
	})

	want := []*MachineFull{{
		Machine: &Machine{Name: "VPS0123", Status: "Running"},
		IPs:     []*IP{{Version: 4, Address: "127.0.0.1"}},
		Config:  &Limits{Bandwidth: 10, CpuCores: 1, CpuPercent: 100, DiskSize: 30, RamMin: 1024, RamMax: 1024},
		OS:      &OS{TemplateID: "1", BrandID: 1, UpdateStatus: &UpdateStatus{HResult: 1, RebootRequired: true, ResultCode: 1, UpdateTime: "2020-10-20 01:02:03"}},
	}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetMachinesFull()
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}

func TestGetMachinesRunning(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/running", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "machines.json")
	})

	want := []*Machine{{Name: "VPS0123", Status: "Running"}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetMachinesRunning()
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}
func TestGetMachinesStopped(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/stopped", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "machines.json")
	})

	want := []*Machine{{Name: "VPS0123", Status: "Running"}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetMachinesStopped()
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}
func TestGetMachine(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "machinefull.json")
	})

	want := &MachineFull{
		Machine: &Machine{Name: "VPS0123", Status: "Running"},
		IPs:     []*IP{{Version: 4, Address: "127.0.0.1"}},
		Config:  &Limits{Bandwidth: 10, CpuCores: 1, CpuPercent: 100, DiskSize: 30, RamMin: 1024, RamMax: 1024},
		OS:      &OS{TemplateID: "1", BrandID: 1, UpdateStatus: &UpdateStatus{HResult: 1, RebootRequired: true, ResultCode: 1, UpdateTime: "2020-10-20 01:02:03"}},
	}
	got, err := client.GetMachine("VPS0123")
	require.NoError(t, err)
	require.Equal(t, want, got)

	got, err = client.GetMachine("VPS01")
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetMachineJobs(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123/jobs", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "jobs.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetMachineJobs("VPS0123")
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)

	got, _, err = client.GetMachineJobs("VPS01")
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetMachineUsers(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123/users", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "users.json")
	})

	want := []*User{{Username: "admin", Role: "admin", Password: "secret"}}
	wantPage := &Pagination{Total: 1, Limit: 50, Page: 1, Pages: 1}
	got, gotPage, err := client.GetMachineUsers("VPS0123")
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)

	got, _, err = client.GetMachineUsers("VPS01")
	require.Error(t, err)
	require.Nil(t, got)
}

func TestChangeMachinePassword(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123/change_password", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		writeFixture(t, w, "change_password.json")
	})

	want := true
	got, err := client.ChangeMachinePassword("VPS0123", "secret")
	require.NoError(t, err)
	require.Equal(t, want, got)

	got, err = client.ChangeMachinePassword("VPS01", "secret")
	require.Error(t, err)
	require.False(t, got)
}

func TestCreateMachine(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, `{"product_id":1,"template_id":1,"location_id":1,"ui_language":"en-US"}`, getBody(t, r))
		writeFixture(t, w, "machinecreate.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Initialize", Status: "Inprogress", StartTime: "2020-10-20 01:02:03"}}
	wantName := "VPS0123"
	opts := &CreateMachineOptions{LocationID: 1, ProductID: 1, TemplateID: 1, UiLanguage: "en-US"}
	gotName, got, err := client.CreateMachine(opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantName, gotName)

	gotName, got, err = client.CreateMachine(nil)
	require.Error(t, err)
	require.Nil(t, got)
	require.Zero(t, gotName)
}
func TestReinstallMachine(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, `{"template_id":1}`, getBody(t, r))
		writeFixture(t, w, "jobspost.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}}
	opts := &ReinstallMachineOptions{TemplateID: 1}
	got, err := client.ReinstallMachine("VPS0123", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)

	got, err = client.ReinstallMachine("VPS01", opts)
	require.Error(t, err)
	require.Nil(t, got)
}

func TestSendMachineCommand(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123/start", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		writeFixture(t, w, "jobspost.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}}
	got, err := client.SendMachineCommand("VPS0123", "start")
	require.NoError(t, err)
	require.Equal(t, want, got)

	got, err = client.SendMachineCommand("VPS0123", "test")
	require.Error(t, err)
	require.Nil(t, got)
}

func TestAddMachineIP(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123/add_ip", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		writeFixture(t, w, "machineaddip.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}}
	wantAddr := "127.0.0.1"
	gotAddr, got, err := client.AddMachineIP("VPS0123")
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantAddr, gotAddr)

	gotAddr, got, err = client.AddMachineIP("VPS01")
	require.Error(t, err)
	require.Nil(t, got)
	require.Zero(t, gotAddr)
}
func TestUpdateMachine(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		require.Equal(t, `{"password":"secret"}`, getBody(t, r))
		writeFixture(t, w, "jobspost.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}}
	opts := &UpdateMachineOptions{Password: "secret"}
	got, err := client.UpdateMachine("VPS0123", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)

	got, err = client.UpdateMachine("VPS01", nil)
	require.Error(t, err)
	require.Nil(t, got)
}

func TestDeleteMachine(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"machines/VPS0123", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		writeFixture(t, w, "jobspost.json")
	})

	want := []*Job{{ID: 1, ParentID: 1, MachineID: 123, Type: "Change", Status: "Complete", StartTime: "2020-10-20 01:02:03"}}
	got, err := client.DeleteMachine("VPS0123")
	require.NoError(t, err)
	require.Equal(t, want, got)

	got, err = client.DeleteMachine("VPS01")
	require.Error(t, err)
	require.Nil(t, got)
}
