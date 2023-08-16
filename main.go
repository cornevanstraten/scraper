package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main(){
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request){
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36");
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response){
		fmt.Println("Response Code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error){
		fmt.Println("error", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement){
		div := h.DOM
		quoteText := div.Find(".text").Text()
		author := div.Find(".author").Text()

		fmt.Println(quoteText, "by", author)
	})

	c.Visit("http://quotes.toscrape.com")
}