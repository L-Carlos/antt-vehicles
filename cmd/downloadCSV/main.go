package main

import (
	"flag"
	"log"
	"strings"

	"github.com/L-Carlos/antt-vehicles/generatename"

	"github.com/gocolly/colly/v2"
)

func main() {
	expectedCSV := generatename.ExpectedCSVName()
	fileOutput := flag.String("output", expectedCSV+".csv", "output [file name]")
	flag.Parse()

	c := colly.NewCollector(
		colly.AllowedDomains("dados.antt.gov.br"),
	)

	c.OnHTML("a.heading", func(e *colly.HTMLElement) {
		if strings.TrimSpace(e.Text) == expectedCSV {
			log.Printf("Found file: %s", expectedCSV)
			c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
		}

	})

	c.OnHTML(".module-resource", func(e *colly.HTMLElement) {
		c.OnResponse(func(r *colly.Response) {
			r.Save(*fileOutput)
		})
		c.Visit(e.ChildAttr("a", "href"))

	})

	c.OnRequest(func(request *colly.Request) {
		log.Println("Visiting", request.URL.String())
	})

	c.Visit("https://dados.antt.gov.br/dataset/veiculos-habilitados")
}
