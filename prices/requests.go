package prices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	Endpoint = "/TradeListSearch"
)

type Client struct {
	Endpoint string
}

type Response struct {
	Status string   `json:"status"`
	Data   []Result `json:"data"`
}

type Result struct {
	Type             string `json:"Type"`
	Region           string `json:"Region,omitempty"`
	MunicipalityCode string `json:"MunicipalityCode"`
	Prefecture       string `json:"Prefecture"`
	Municipality     string `json:"Municipality"`
	DistrictName     string `json:"DistrictName"`
	TradePrice       string `json:"TradePrice"`
	Area             string `json:"Area"`
	LandShape        string `json:"LandShape,omitempty"`
	Frontage         string `json:"Frontage,omitempty"`
	TotalFloorArea   string `json:"TotalFloorArea,omitempty"`
	BuildingYear     string `json:"BuildingYear,omitempty"`
	Structure        string `json:"Structure"`
	Use              string `json:"Use"`
	Direction        string `json:"Direction,omitempty"`
	Classification   string `json:"Classification,omitempty"`
	Breadth          string `json:"Breadth,omitempty"`
	CityPlanning     string `json:"CityPlanning"`
	CoverageRatio    string `json:"CoverageRatio"`
	FloorAreaRatio   string `json:"FloorAreaRatio"`
	Period           string `json:"Period"`
	FloorPlan        string `json:"FloorPlan,omitempty"`
	Renovation       string `json:"Renovation,omitempty"`
	Remarks          string `json:"Remarks,omitempty"`
}

func (p *Response) Num() int {
	return len(p.Data)
}

type Request struct {
	Area    string
	City    string
	Station string
	From    time.Time
	To      time.Time
}

type Query struct {
	Area    string `url:"area,omitempty"`
	City    string `url:"city,omitempty"`
	Station string `url:"station,omitempty"`
	From    string `url:"from"`
	To      string `url:"to"`
}

func (p *Client) Get(r *Request) (*Response, error) {
	u, err := url.ParseRequestURI(p.Endpoint)
	if err != nil {
		return nil, err
	}
	u.RawQuery = r.toQuery()
	fmt.Printf("%+v %v\n", u.String(), r.toQuery())
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: 10 * time.Second,
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

func toDate(t time.Time) string {
	year := t.Year()

	var num int
	if 10 < t.Month() {
		num = 4
	} else if 6 < t.Month() {
		num = 3
	} else if 3 < t.Month() {
		num = 2
	} else {
		num = 1
	}

	return fmt.Sprintf("%d%d", year, num)
}

func (p *Request) toQuery() string {
	q := new(Query)
	q.Area = p.Area
	q.City = p.City
	q.Station = p.Station
	q.From = toDate(p.From)
	q.To = toDate(p.To)
	v, err := query.Values(q)
	if err != nil {
		return ""
	}
	return v.Encode()
}
