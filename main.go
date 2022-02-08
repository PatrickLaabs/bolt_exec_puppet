package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	//cmd := exec.Command("puppet", "agent", "--test", "--noop")
	//cmd := exec.Command("/usr/local/bin/puppet", "agent", "--test", "--noop")

	cmd := exec.Command("/bin/sh", "-c", "ls | grep sdsd")
	fmt.Println(">> executing binary...")

	// stdout, err := cmd.StdoutPipe()
	Stdin, err := cmd.StdinPipe()
	if err != nil {
		// fmt.Println("> stdoutpipe err")
		fmt.Println("> stdinpipe err")
		log.Fatal(err)
	}

	Stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("> stdoutpipe err")
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("> start err")
		log.Fatal(err)
	}

	if err := Stdin.Close(); err != nil {
		fmt.Println("> stdin.Close err")
		log.Fatal(err)
	}

	// if err := cmd.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	Bytes, _ := io.ReadAll(Stdout)

	if err := cmd.Wait(); err != nil {
		fmt.Println("> wait err")
		log.Fatal(err)
	}

	fmt.Println(string(Bytes))
	fmt.Println(">> executing binary succeeded")

	// exitCodes()
}

// func exportPath() {
// 	// Code for setting up path for puppet binary
// }

// func exitCodes() {
// 	// Need a way to get hold onto exit codes returned to stdin

// 	// CASE / Switch func?

// 	fmt.Println("foo:", os.Getenv("PWD"))
// 	fmt.Println("foo:", os.Getenv("$?"))
// }
