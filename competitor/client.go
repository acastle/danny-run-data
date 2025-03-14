package competitor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	baseUrl    *url.URL
	privateKey string
}

func NewClient(baseUrl string, privateKey string) *Client {
	url, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}

	return &Client{
		baseUrl:    url,
		privateKey: privateKey,
	}
}

func (c *Client) GetRunData(id string, limit int, skip int) (*EventDataPayload, error) {
	query := c.baseUrl.Query()
	query.Add("$limit", strconv.Itoa(limit))
	query.Add("$skip", strconv.Itoa(skip))
	query.Add("$sort[FinishRankOverall]", "1")
	addr := c.baseUrl.JoinPath("public", "result", "subevent", id)
	addr.RawQuery = query.Encode()
	req, err := http.NewRequest(http.MethodGet, addr.String(), nil)
	req.Header.Add("Wtc_priv_key", c.privateKey)

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
	var result EventDataPayload
	err = decoder.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return &result, nil
}
