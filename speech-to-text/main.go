package main

import (
	"flag"
	"fmt"

	"github.com/taylorchu/lana"
)

var (
	lang = flag.String("lang", "en", "language")
)

func main() {
	flag.Parse()

	r, err := lana.Listen(*lang)
	if err != nil || len(r.Alternative) == 0 {
		return
	}
	fmt.Println(r.Alternative[0].Transcript)
}
