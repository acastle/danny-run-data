package xact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	baseUrl *url.URL
}

func NewClient(baseUrl string) *Client {
	url, _ := url.Parse("https://results.xacte.com")
	xclient := Client{
		baseUrl: url,
	}
	return &xclient
}

func (c *Client) GetRunData(id string, skip int) (*Result, error) {
	query := c.baseUrl.Query()
	query.Add("eventId", id)
	query.Add("offset", strconv.Itoa(skip))
	addr := c.baseUrl.JoinPath("json", "search")
	addr.RawQuery = query.Encode()
	req, err := http.NewRequest(http.MethodGet, addr.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get run data: %w", err)
	}

	println(addr.String())
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non success status code from api: %s, %d", addr.String(), resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	var result Result
	err = decoder.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return &result, nil
}
