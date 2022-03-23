package winvps

import "net/http"

// Represents a winvps template
type Template struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Returns all available templates. Info from Pagination can be used to get templates using RequestOptions
// default Limit 50
func (c *Client) GetTemplates(opts ...*RequestOptions) ([]*Template, *Pagination, error) {
	u := "templates"

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*Template
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}
