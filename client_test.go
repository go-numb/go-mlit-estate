package mlit_test

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-numb/go-mlit-estate"
	"github.com/go-numb/go-mlit-estate/areas"
	"github.com/go-numb/go-mlit-estate/prices"
)

func TestGetAreas(t *testing.T) {
	c := mlit.New(false)

	res, err := c.Areas.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}

func Test(t *testing.T) {
	state := "東京都"
	code := areas.ToAreaCode(state)
	assert.NotEqual(t, "", code, "error")
	fmt.Printf("%v	%+v\n", state, code)
	t.Log(code)
}

func TestGetPrices(t *testing.T) {
	c := mlit.New(false)

	res, err := c.Prices.Get(&prices.Request{
		Area: "13",
		City: "13101",

		From: time.Now().Add(-365 * 24 * time.Hour),
		To:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}

	for i := range res.Data {
		if res.Data[i].Remarks != "" {
			fmt.Printf("%+v\n", res.Data[i])
		}
	}
}

func TestGetAll(t *testing.T) {
	c := mlit.New(false)
	// f, err := os.OpenFile("./areas/cities.csv", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// w := csv.NewWriter(f)

	// count := 47
	// var cities []string
	// for i := 0; i < count; i++ {
	// 	res, err := c.Areas.Get(i + 1)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	for j := range res.Data {
	// 		cities = append(cities, res.Data[j].ID)
	// 		w.Write([]string{
	// 			res.Data[j].ID, res.Data[j].Name,
	// 		})
	// 	}
	// 	w.Flush()
	// 	time.Sleep(time.Second)
	// }
	// f.Close()

	cities := read()

	var results []prices.Result
	for i := range cities {
		resp, err := c.Prices.Get(&prices.Request{
			Area: cities[i][1],
			City: cities[i][0],

			From: time.Now().Add(-365 * 24 * time.Hour),
			To:   time.Now(),
		})
		if err != nil {
			t.Fatal(err)
		}

		// fmt.Printf("%+v\n", resp.Data)
		results = append(results, resp.Data...)
		time.Sleep(time.Second)
	}

	file, err := os.OpenFile("./results.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body, err := json.Marshal(results)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := file.Write(body); err != nil {

	}

	t.Log("success")

}

func read() [][]string {
	f, err := os.OpenFile("./areas/cities.csv", os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil
	}
	r := csv.NewReader(f)

	results := make([][]string, 0)
	for {
		row, err := r.Read()
		if err != nil {
			break
		}

		results = append(results, []string{
			row[0],
			fmt.Sprintf("%s", string([]rune(row[0])[:2])),
		})
	}

	return results
}
