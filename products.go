package winvps

import "net/http"

// Represents a winvps product info
type Product struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Limits *Limits `json:"limits"`
}

// Represents product limits
type Limits struct {
	CpuPercent int `json:"cpu_percent"`
	CpuCores   int `json:"cpu_cores"`
	RamMin     int `json:"ram_min"`
	RamMax     int `json:"ram_max"`
	DiskSize   int `json:"disk_size"`
	Bandwidth  int `json:"bandwidth"`
	Traffic    int `json:"traffic"`
}

// Returns all available products. Info from Pagination can be used to get products using RequestOptions
// default Limit 50
func (c *Client) GetProducts(opts ...*RequestOptions) ([]*Product, *Pagination, error) {
	u := "products"

	req, err := c.NewRequest(http.MethodGet, u, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result []*Product
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Pagination, nil
}
