package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	bad_words := []string{"fuck", "shit", "crap"}

	fmt.Println("Working")

	c := colly.NewCollector(
		colly.AllowedDomains("www.songlyrics.com", "songlyrics.com", "https://www.songlyrics.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})

	c.OnHTML("#songLyricsDiv-outer", func(h *colly.HTMLElement) {
		div := h.DOM
		lyrics := div.Find(".songLyricsV14").Text()

		i := 0
		for i < len(bad_words) {
			lyrics = strings.Replace(lyrics, bad_words[i], "[****]", -1)
			i += 1
		}

		fmt.Println(lyrics)
	})

	c.Visit("https://www.songlyrics.com/eminem/not-afraid-lyrics/")
}
