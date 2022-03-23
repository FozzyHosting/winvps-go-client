package winvps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	baseURL    = "https://winvps.fozzy.com"
	apiVerPath = "/api/v2/"
	userAgent  = "go-winvps"
)

// Represents api client
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	token      string
	UserAgent  string
}

// Represents api response
type Response struct {
	Data       json.RawMessage `json:"data"`
	Pagination Pagination      `json:"pagination"`
	Error      string          `json:"error,omitempty"`
}

// Represents pagination info
type Pagination struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

// Method returns next page if available
func (p *Pagination) NextPage() int {
	if p.Pages-p.Page > 0 {
		p.Page++
		return p.Page
	}
	return 0
}

// Method returns previous page if available
func (p *Pagination) PreviousPage() int {
	if p.Page > 1 {
		p.Page--
		return p.Page
	}
	return 0
}

// Represents pagination options
type RequestOptions struct {
	Limit int `url:"limit,omitempty"`
	Page  int `url:"page,omitempty"`
}

// Represents option func for customize api client
type Option func(*Client) error

// Set BaseURL for api client
func BaseURL(baseURL string) Option {
	return func(c *Client) error {
		return c.setBaseURL(baseURL)
	}
}

// Set BaseURL, validate it and add api path to it
func (c *Client) setBaseURL(urlStr string) error {
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	baseURL.Path = apiVerPath
	c.baseURL = baseURL
	return nil
}

// Exec all passed option func to the client
func (c *Client) parseOptions(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}

// Creates a new instance of api client
func NewClient(token string, opts ...Option) (*Client, error) {
	c := &Client{
		UserAgent:  userAgent,
		token:      token,
		httpClient: &http.Client{Timeout: time.Second * 30},
	}
	c.setBaseURL(baseURL)
	if err := c.parseOptions(opts...); err != nil {
		return nil, err
	}
	return c, nil
}

// Used for validate request options
type Validator interface {
	Validate() error
}

// helper func, used for validate if struct contains required fields passed in opts
func checkRequiredOpts(v interface{}, opts []string) error {
	r := reflect.ValueOf(v)
	if reflect.Indirect(r).Kind() != reflect.Struct {
		return fmt.Errorf("not a struct, can't continue")
	}
	for _, opt := range opts {
		f := reflect.Indirect(r).FieldByName(opt)
		if !f.IsValid() {
			return fmt.Errorf("struct %+v not contain required field %s", v, opt)
		}
		switch f.Interface().(type) {
		case int:
			t := f.Interface().(int)
			if t == 0 {
				return fmt.Errorf("missing required option %s", opt)
			}
		case string:
			t := f.Interface().(string)
			if t == "" {
				return fmt.Errorf("missing required option %s", opt)
			}
		default:
			return fmt.Errorf("unsupported option type: %v", f.Type())
		}

	}
	return nil
}

// Make an http request, check and parse response
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	// Parse data field from response
	if v != nil {
		result := &Response{}
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return nil, fmt.Errorf("status: %d, unable to decode response, unknown format: %v", resp.StatusCode, err)
		}
		// Decode the data field
		if result.Data == nil {
			return nil, fmt.Errorf("status: %d, missing data from response", resp.StatusCode)
		}
		if err := json.Unmarshal(result.Data, v); err != nil {
			return result, fmt.Errorf("status: %d, unable to parse response data: %s", resp.StatusCode, err)
		}
		return result, nil
	}

	return nil, err
}

// Creates and validates a new request
// sets required headers
func (c *Client) NewRequest(method, path string, opt interface{}, opts []*RequestOptions) (*http.Request, error) {
	u := *c.baseURL
	u.Path = c.baseURL.Path + path

	// Prepare headers
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("API-KEY", c.token)
	reqHeaders.Set("User-Agent", c.UserAgent)

	// Validate and marshall request body if any
	var body []byte
	if method == http.MethodPost || method == http.MethodPut {
		reqHeaders.Set("Content-Type", "application/json")
		if opt != nil {
			if validator, ok := opt.(Validator); ok {
				if err := validator.Validate(); err != nil {
					return nil, err
				}
			}
			var err error
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			}
		}
	}

	// Prepare query params if any
	var options *RequestOptions
	if len(opts) >= 1 {
		options = opts[0]
		q, err := query.Values(options)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	// Create a new request
	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Set the request specific headers
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

// Checks the API response for errors
func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if data != nil {
		result := &Response{}
		if err := json.Unmarshal(data, result); err != nil {
			return fmt.Errorf("status: %d, can't parse error, unknown format, raw data: %s", r.StatusCode, data)
		}
		return fmt.Errorf("status: %d, error: %s", r.StatusCode, result.Error)
	}
	return fmt.Errorf("status: %d, empty response", r.StatusCode)
}
