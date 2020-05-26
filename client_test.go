package mlit_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-numb/go-mlit-estate"
	"github.com/go-numb/go-mlit-estate/areas"
	"github.com/go-numb/go-mlit-estate/prices"
)

func TestGetAreas(t *testing.T) {
	c := mlit.New(false)

	res, err := c.Areas.Get("宮城")
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
	c := mlit.New(true)

	res, err := c.Prices.Get(&prices.Request{
		Area: "04",
		City: "04101",

		From: time.Now().Add(-365 * 24 * time.Hour),
		To:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}

	for i := range res.Data {
		fmt.Printf("%+v\n", res.Data[i])
	}
}
