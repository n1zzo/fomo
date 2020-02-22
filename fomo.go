package main

import (
    "fmt"

    "github.com/n1zzo/fomo/backends"
    "github.com/n1zzo/fomo/frontends"

    "github.com/kr/pretty"
)

type Event struct {
    name string
    performer string
    text string
    url string
    time string
    image_url string
}

func main() {
    fmt.Println("Hello FOMO!")

    events := backends.Elfo()
    for _, e := range events {
        pretty.Printf("\n")
        pretty.Println(e)
        frontends.Instagram(e)
        return
    }

}
