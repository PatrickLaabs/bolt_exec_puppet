package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("puppet", "agent", "--test", "--noop")
	//cmd := exec.Command("/usr/local/bin/puppet", "agent", "--test", "--noop")
	// cmd := exec.Command("/bin/sh", "-c", "ls | grep fdfd")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// if err := cmd.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	Bytes, _ := io.ReadAll(stdout)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(Bytes))
}

// func exportPath() {
// 	// Code for setting up path for puppet binary
// }
