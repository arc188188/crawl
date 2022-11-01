package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Actor struct {
	CnName string
	JpName string
	Age    uint8
}

func Actress() {
	c := colly.NewCollector(colly.AllowedDomains("avday.tv", "www.avday.tv"))

	c.OnHTML("body > div.container.container-content.no-cate-container-content > div > div.row.row-actor > div > a > img", func(e *colly.HTMLElement) {
		fmt.Printf("%q %s\n", e.Attr("title"), e.Attr("src"))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit("https://avday.tv/actress")
	if err != nil {
		panic(err.Error())
	}
}
