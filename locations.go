package winvps

import "net/http"

// Represents a location info
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Returns all available locations. Info from Pagination can be used to get locations using RequestOptions
// default Limit 50
func (c *Client) GetLocations(opts ...*RequestOptions) ([]*Location, *Pagination, error) {
	u := "locations"

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*Location
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}
