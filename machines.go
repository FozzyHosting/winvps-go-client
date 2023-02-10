package winvps

import (
	"fmt"
	"net/http"
	"net/url"
)

// List of available CreateMachine() options
type CreateMachineOptions struct {
	Description string `json:"description,omitempty"`
	Password    string `json:"password,omitempty"`
	ProductID   int    `json:"product_id,omitempty"`
	TemplateID  int    `json:"template_id,omitempty"`
	BrandID     int    `json:"brand_id,omitempty"`
	DiskType    string `json:"disk_type,omitempty"`
	LocationID  int    `json:"location_id,omitempty"`
	AddDisk     int    `json:"add_disk,omitempty"`
	AddRam      int    `json:"add_ram,omitempty"`
	AddCpu      int    `json:"add_cpu,omitempty"`
	AddBand     int    `json:"add_band,omitempty"`
	AutoStart   int    `json:"auto_start,omitempty"`
	AddIPv6     int    `json:"add_ipv6,omitempty"`
	UiLanguage  string `json:"ui_language,omitempty"`
}

// List of available UpdateMachine() options
type UpdateMachineOptions struct {
	Password  string `json:"password,omitempty"`
	ProductID int    `json:"product_id,omitempty"`
	AddDisk   int    `json:"add_disk,omitempty"`
	AddRam    int    `json:"add_ram,omitempty"`
	AddCpu    int    `json:"add_cpu,omitempty"`
	AddBand   int    `json:"add_band,omitempty"`
}

// List of available ReinstallMachine() options
type ReinstallMachineOptions struct {
	Password   string `json:"password,omitempty"`
	TemplateID int    `json:"template_id,omitempty"`
	BrandID    int    `json:"brand_id,omitempty"`
	AutoStart  int    `json:"auto_start,omitempty"`
}

// Represents a winvps machine info
type Machine struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Notes  string `json:"notes"`
}

// Represents a full winvps machine info
type MachineFull struct {
	*Machine
	IPs    []*IP   `json:"ips"`
	OS     *OS     `json:"os"`
	Config *Limits `json:"config"`
}

// Represents an IP info
type IP struct {
	Version int    `json:"version"`
	Address string `json:"address"`
}

// Represents a winvps OS info
type OS struct {
	TemplateID   string        `json:"template_id"`
	BrandID      int           `json:"brand_id"`
	UpdateStatus *UpdateStatus `json:"update_status"`
}
type UpdateStatus struct {
	HResult        int    `json:"h_result"`
	RebootRequired bool   `json:"reboot_required"`
	ResultCode     int    `json:"result_code"`
	UpdateTime     string `json:"update_time"`
}

// Represents a winvps user info
type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

// Represents a simple result response
type result struct {
	Result bool `json:"result"`
}

// Represents a change_password request
type password struct {
	Password string `json:"password"`
}

type AdditionalUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate CreateMachineOptions for required options
func (t *CreateMachineOptions) Validate() error {
	rOpts := []string{"ProductID", "TemplateID", "LocationID"}
	if err := checkRequiredOpts(t, rOpts); err != nil {
		return err
	}
	if len(t.DiskType) > 0 && (t.DiskType != "hdd" && t.DiskType != "ssd") {
		return fmt.Errorf("allowed disk type 'hdd' or 'ssd' but '%s' passed", t.DiskType)
	}
	return nil
}

// Create a new machine with specified CreateMachineOptions
// returns new machine name and Jobs list
func (c *Client) CreateMachine(opt *CreateMachineOptions) (string, []*Job, error) {
	u := "machines"
	req, err := c.NewRequest(http.MethodPost, u, opt, nil)
	if err != nil {
		return "", nil, err
	}
	result := new(struct {
		Name string `json:"name"`
		Jobs []*Job `json:"jobs"`
	})
	_, err = c.Do(req, result)
	if err != nil {
		return "", nil, err
	}

	return result.Name, result.Jobs, nil
}

// Update machine with specified UpdateMachineOptions
func (c *Client) UpdateMachine(name string, opt *UpdateMachineOptions) ([]*Job, error) {
	u := fmt.Sprintf("machines/%s", url.PathEscape(name))

	req, err := c.NewRequest(http.MethodPut, u, opt, nil)
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Jobs []*Job `json:"jobs"`
	})
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result.Jobs, nil
}

// Reinstall machine with specified ReinstallMachineOptions
func (c *Client) ReinstallMachine(name string, opt *ReinstallMachineOptions) ([]*Job, error) {
	u := fmt.Sprintf("machines/%s", url.PathEscape(name))

	req, err := c.NewRequest(http.MethodPost, u, opt, nil)
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Jobs []*Job `json:"jobs"`
	})
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result.Jobs, nil
}

// Returns all machines. Limit and Page can be set via RequestOptions
// default Limit 50
func (c *Client) GetMachines(opts ...*RequestOptions) ([]*Machine, *Pagination, error) {
	u := "machines"

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*Machine
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}

// Returns all machines with full info. Info from Pagination can be used to get machines using RequestOptions
// default Limit 50
func (c *Client) GetMachinesFull(opts ...*RequestOptions) ([]*MachineFull, *Pagination, error) {
	u := "machines/full"

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*MachineFull
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}

// Returns all running machines. Info from Pagination can be used to get machines using RequestOptions
// default Limit 50
func (c *Client) GetMachinesRunning(opts ...*RequestOptions) ([]*Machine, *Pagination, error) {
	u := "machines/running"

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*Machine
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}

// Returns all stopped machines. Info from Pagination can be used to get machines using RequestOptions
// default Limit 50
func (c *Client) GetMachinesStopped(opts ...*RequestOptions) ([]*Machine, *Pagination, error) {
	u := "machines/stopped"

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*Machine
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}

// Return specific machine full info
func (c *Client) GetMachine(name string) (*MachineFull, error) {
	u := fmt.Sprintf("machines/%s", url.PathEscape(name))

	req, err := c.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, err
	}

	result := new(MachineFull)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Returns all jobs assigned to machine. Info from Pagination can be used to get jobs using RequestOptions
// default Limit 50
func (c *Client) GetMachineJobs(name string, opts ...*RequestOptions) ([]*Job, *Pagination, error) {
	u := fmt.Sprintf("machines/%s/jobs", url.PathEscape(name))

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*Job
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}

// Returns list of additional system users. Info from Pagination can be used to get users using RequestOptions
// default Limit 50
func (c *Client) GetMachineUsers(name string, opts ...*RequestOptions) ([]*User, *Pagination, error) {
	u := fmt.Sprintf("machines/%s/users", url.PathEscape(name))

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*User
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}

// Change VPS machine password
func (c *Client) ChangeMachinePassword(name, pass string) (bool, error) {
	u := fmt.Sprintf("machines/%s/change_password", url.PathEscape(name))

	opt := &password{Password: pass}
	req, err := c.NewRequest(http.MethodPost, u, opt, nil)
	if err != nil {
		return false, err
	}

	result := new(result)
	_, err = c.Do(req, result)
	if err != nil {
		return false, err
	}

	return result.Result, nil
}

// helper func to validate command for SendMachineCommand()
func validateCommand(command string) error {
	availableCommands := []string{
		"start",
		"stop",
		"restart",
		"enable_rdp",
		"enable_network",
		"restart_mt",
		"run_updates_install"}

	for _, c := range availableCommands {
		if command == c {
			return nil
		}
	}
	return fmt.Errorf("wrong command passed '%s', available commands: %s", command, availableCommands)
}

// Send command to machine. Available commands is:
// start, stop, restart, enable_rdp, enable_network, restart_mt, run_updates_install
func (c *Client) SendMachineCommand(name, command string) ([]*Job, error) {
	if err := validateCommand(command); err != nil {
		return nil, err
	}
	u := fmt.Sprintf("machines/%s/%s", url.PathEscape(name), url.PathEscape(command))

	req, err := c.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Jobs []*Job `json:"jobs"`
	})
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result.Jobs, nil
}

// Add IP to specified machine
// returns new IP address and Jobs list
func (c *Client) AddMachineIP(name string) (string, []*Job, error) {
	u := fmt.Sprintf("machines/%s/add_ip", url.PathEscape(name))

	req, err := c.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return "", nil, err
	}
	result := new(struct {
		Address string `json:"address"`
		Jobs    []*Job `json:"jobs"`
	})
	_, err = c.Do(req, result)
	if err != nil {
		return "", nil, err
	}

	return result.Address, result.Jobs, nil
}

// Create machine deletion job
func (c *Client) DeleteMachine(name string) ([]*Job, error) {
	u := fmt.Sprintf("machines/%s", url.PathEscape(name))

	req, err := c.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, err
	}
	result := new(struct {
		Jobs []*Job `json:"jobs"`
	})
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result.Jobs, nil
}
