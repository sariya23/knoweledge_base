package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	ScanCredists int    `json:"scan_credits"`
	QueryCredits int    `json:"query_credits"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
	Telnet       bool   `json:"telnet"`
}

func (c *Client) APIInfo() (*APIInfo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, c.apiKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result APIInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
