package shodan

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HostLocation struct {
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   string  `json:"postal_code"`
	CountryCode  string  `json:"country_code"`
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     int     `json:"area_code"`
	DMACode      int     `json:"dma_code"`
	Longitude    float32 `json:"longitude"`
	Latitude     float32 `json:"latitude"`
}
type Host struct {
	Domains   []string     `json:"domains"`
	Hostnames []string     `json:"hostnames"`
	OS        string       `json:"os"`
	Timestamp string       `json:"timestamp"`
	ISP       string       `json:"isp"`
	ASN       string       `json:"asn"`
	Org       string       `json:"org"`
	Data      string       `json:"data"`
	IPString  string       `json:"ip_str"`
	Location  HostLocation `json:"location"`
	IP        int64        `json:"ip"`
	Port      int          `json:"port"`
}
type HostSearch struct {
	Matches []Host `json:"matches"`
}

func (c *Client) HostSearch(q string) (*HostSearch, error) {
	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s?query=%s", BaseURL, c.apiKey, q))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result HostSearch
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(d))
	err = json.Unmarshal(d, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
