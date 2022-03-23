package winvps

import (
	"net/http"
)

// Represens a brand info
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Returns all available brands. Info from Pagination can be used to get brands using RequestOptions
// default Limit 50
func (c *Client) GetBrands(opts ...*RequestOptions) ([]*Brand, *Pagination, error) {
	u := "brands"

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*Brand
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}
