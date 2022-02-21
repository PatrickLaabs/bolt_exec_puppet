package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
)

// https://pkg.go.dev/github.com/jessevdk/go-flags#readme-example
// https://lightstep.com/blog/getting-real-with-command-line-arguments-and-goflags/
type Options struct {
	Name string `short:"n" description:"Your name, for a greeting" default:"Unknown"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello", opts.Name)
}
