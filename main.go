package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// run task list
	// 執行瀏覽器任務
	var myMapPtr map[string]string

	err := chromedp.Run(ctx,
		chromedp.Navigate(encodeURIComponent()),
		chromedp.WaitVisible("iframe[src]", chromedp.ByQuery),
		chromedp.Attributes("iframe[src]", &myMapPtr),
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	if myMapPtr["src"] != "" {
		fmt.Println("src:", myMapPtr["src"])
		getData(myMapPtr["src"])
	} else {
		fmt.Println("Attributes:", myMapPtr)
	}

}

func getData(path string) {

	req, err := http.Get(path)

	if err != nil {
		return
	}

	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)

	if err != nil {
		return
	}
	fmt.Println(string(data))
}

func encodeURIComponent() string {
	udd := "https://codesandbox.io/p/devbox/github/laof/ssssandbox/tree/main/?file=/index.js:1,10"
	// req, err := http.Get("https://codesandbox.io/p/devbox/github/laof/ssssandbox/tree/main/?file=%2Findex.js%3A1%2C10")
	u, _ := url.Parse(udd)

	q := u.Query()
	u.RawQuery = q.Encode()
	return (u.String())
}
