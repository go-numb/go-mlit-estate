package areas

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	Endpoint = "/CitySearch"
)

type Client struct {
	Endpoint string
}

type Response struct {
	Status string `json:"status"`
	Data   []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

func (p *Client) Get(area interface{}) (*Response, error) {
	u, err := url.ParseRequestURI(p.Endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("area", toCode(area))
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	s := new(Response)
	if err := json.NewDecoder(res.Body).Decode(s); err != nil {
		return nil, err
	}

	return s, nil
}

func toCode(v interface{}) string {
	switch area := v.(type) {
	case int:
		return fmt.Sprintf("%d", area)
	case string:
		return ToAreaCode(area)
	}
	return ""
}

func ToAreaCode(state string) string {
	abs, err := os.Getwd()
	if err != nil {
		return ""
	}
	f, err := os.Open(filepath.Join(abs, "areas", "list.csv"))
	if err != nil {
		return ""
	}
	defer f.Close()

	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err != nil {
			break
		}

		if !strings.Contains(row[1], state) {
			continue
		}
		return row[0]
	}
	return ""
}
