package main

import (
	"fmt"
	"log"
	"os/exec"
	"io"
)

func main() {
	// cmd := exec.Command("/usr/local/bin/puppet", "agent", "--test", "--noop")
	cmd := exec.Command("/bin/sh", "-c", "ls | grep Docker")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Close(); err != nil {
		log.Fatal(err)
	}

	Bytes, _ := io.ReadAll(stdout)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf(Bytes)
}
