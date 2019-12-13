package main

import (
	"github.com/vanhtuan0409/tget"
)

func main() {
	r := tget.NewRequest("https://docs.google.com/presentation/u/0/d/1gbFnt7hRff0tZK2B6WzTs4f-WbJgIdjHtaB7L9zQMC4/export/pptx")
	if err := r.Download(".", 5); err != nil {
		panic(err)
	}
}
