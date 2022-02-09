package main

import (
	"fmt"

	"github.com/bitfield/script"
)

func main() {
	// puppet cmd: /usr/local/bin/puppet agent --test --noop)
	fmt.Println(">> running puppet")
	p := script.Exec("bash -c 'puppet agent --test --noop'")
	fmt.Println("Exit Status:", p.ExitStatus())
	if err := p.Error(); err != nil {
		p.SetError(nil)
		out, _ := p.Stdout()
		fmt.Println("if err output:", out)
	} else {
		out, _ := p.Stdout()
		fmt.Printf("else err output: %v", out)
	}
}
