# go-mlit-estate
go-mlit-estate is estate price API wrapper.
[不動産取引価格情報取得API](https://www.land.mlit.go.jp/webland/api/)

# Usage 
``` go
package main

import (
    "github.com/go-numb/go-mlit-estate"
    "github.com/go-numb/go-mlit-estate/prices"
)

func main() {
    // 日本語 or English
    isEnglish := false
    client := mlit.New()

    res, err := client.Prices.Get(&prices.Request{
        Area: "04",
        City: "04101",

        From: time.Now().Add(-365 * 24 * time.Hour),
        To:   time.Now(),
    })
    if err != nil {
        return
    }

    for i := range res.Data {
        fmt.Printf("%+v\n", res.Data[i])
    }
}

```


## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-mlit-estate/blob/master/LICENSE)