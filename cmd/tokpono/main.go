package main

import (
	"errors"
	"flag"
	"log"
	"os"

	_ "github.com/ghigt/gocyk"
)

var (
	jufra = flag.String("jufra", "", "lo jufra ka tokpono")
)

// la i'enai cu srera jungau
func ihenai(koha error) {
	log.Println(koha)
	flag.Usage()
	os.Exit(2)
}

func main() {
	flag.Parse()

	if *jufra == "" {
		ihenai(errors.New("na jufra"))
	}
}
