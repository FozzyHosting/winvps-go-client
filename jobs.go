package winvps

import (
	"fmt"
	"net/http"
)

// Represents a winvps job
type Job struct {
	ID        int    `json:"id"`
	ParentID  int    `json:"parent_id"`
	MachineID int    `json:"machine_id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	StartTime string `json:"start_time"`
}

// Returns all planned and completed jobs. Info from Pagination can be used to get jobs using RequestOptions
// default Limit 50
func (c *Client) GetJobs(opts ...*RequestOptions) ([]*Job, *Pagination, error) {
	u := "jobs"

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

// Returns all planned jobs. Info from Pagination can be used to get jobs using RequestOptions
// default Limit 50
func (c *Client) GetPendingJobs(opts ...*RequestOptions) ([]*Job, *Pagination, error) {
	u := "jobs/pending"

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

// Returns a single job info
func (c *Client) GetJob(id int) (*Job, error) {
	u := fmt.Sprintf("jobs/%d", id)

	req, err := c.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, err
	}

	result := new(Job)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Cancel specified job
func (c *Client) CancelJob(id int) error {
	u := fmt.Sprintf("jobs/%d", id)

	req, err := c.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return err
}
