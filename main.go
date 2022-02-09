package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	//cmd := exec.Command("/usr/local/bin/puppet", "agent", "--test", "--noop")
	cmd := exec.Command("/bin/sh", "-c", "ls | grep main")
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
		log.Fatalf("cmd.Start: %v", err)
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
		fmt.Println("NEw Line: ", err)
		log.Fatal("\nfatal wait err: ", err)
	} else {
		fmt.Println("printing wait err - else", err)
	}

	errOut := err
	fmt.Println("printing err out of 'scope'", errOut)

	switch errOut {
	case 1:
		fmt.Println("exitCode 1")
	default:
		fmt.Print("succeeded", err)
	}

	// if err != nil {
	// 	i :=
	// 	switch i {
	// 	case 0:
	// 		fmt.Println("zero")
	// 	case 1:
	// 		fmt.Println("one")
	// 	case 2:
	// 		fmt.Println("two")
	// 	default:
	// 		fmt.Println("Unknown Number")
	// 	}
	// }

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

// }
