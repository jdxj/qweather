package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ModeDebug = "debug"
)

func NewClient(mode string) *Client {
	c := &Client{
		httpClient: &http.Client{},
		mode:       mode,
	}
	return c
}

type Client struct {
	httpClient *http.Client

	mode string
}

func (c *Client) do(query string) (*GeneralResponse, error) {
	req := newRequest(query)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.mode == ModeDebug {
		fmt.Printf("response body: %s\n", data)
	}

	gr := new(GeneralResponse)
	if err := json.Unmarshal(data, gr); err != nil {
		return nil, err
	}

	if gr.Code != "200" {
		return nil, fmt.Errorf("query: %s, code: %s, body: %s",
			query, gr.Code, data)
	}
	return gr, nil
}

func newRequest(url string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return req
}
