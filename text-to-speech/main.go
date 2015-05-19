package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/taylorchu/lana"
)

var (
	lang = flag.String("lang", "en", "language")
)

func main() {
	flag.Parse()

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return
	}
	lana.Say(string(b), *lang)
}
