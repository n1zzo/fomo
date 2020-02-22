package backends

import (
    "fmt"
    "strings"

    . "github.com/n1zzo/fomo/data"

    "github.com/gocolly/colly/v2"
    "github.com/kr/pretty"
)

//month_month_re := regexp.MustCompile(`(\d+)\s?(\w+)\s?-\s?(\d+)\s?(\w+)`)
//date_month_re := regexp.MustCompile(`(\d+)\s?-\s?(\d+)\s?(\w+)`)
//date_date_re := regexp.MustCompile(`(\d+)\s?(\w+)\s?(\d+)\s?(\w+)`"")

func Elfo() (events []Event) {
    events = make([]Event, 0)

    c := colly.NewCollector(colly.AllowedDomains("elfo.org"))

    // On the homepage search for events divs
    c.OnHTML("div .imgcontainer", func(e *colly.HTMLElement) {
        event_url := e.ChildAttr("a", "href")
        pretty.Println("Found element: ", event_url)
        c.Visit(e.Request.AbsoluteURL(event_url))
	})

    // When visiting an event, scrape off information
    c.OnHTML("div .scheda", func(e *colly.HTMLElement) {
        // Extract performer info
        performer := e.ChildText("div .autore")
        if performer == "" {
            performer = strings.Join(e.ChildTexts("div .cast"), " ")
        }
        // Extract timepoints
        date_str := e.ChildText("div .data")
        //time_str := e.ChildText("div .orario")
        //switch {
        //case month_month_re.FindStringSubmatch(date_str):
        //    start_time := 
        //}
        event := Event{
            Name: e.ChildText("div .titolo"),
            Performer: performer,
            Text: e.ChildText("div #testo"),
            Url: e.Request.AbsoluteURL(e.Request.URL.RequestURI()),
            Time: date_str,
            ImageUrl: e.Request.AbsoluteURL(e.ChildAttr("img", "src")),
        }
        events = append(events, event)
    })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

    // Start scraping
	c.Visit("https://elfo.org/")

    return events
}
