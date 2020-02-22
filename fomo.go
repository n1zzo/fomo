package fomo

import (
    "fmt"

    "github.com/n1zzo/fomo/backends"
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

    backends.Elfo()
}
