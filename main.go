package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// This line is needed for testing purposes, to run puppet with noop on manifest.pp
	puppetCmd := exec.Command("puppet", "apply", "--noop", "--test", "--debug", "manifest/manifest.pp", "|", "grep", "exit")

	// puppetCmd := exec.Command("puppet", "apply", "--noop", "--test", "--debug")

	// Input & Output Pipes
	puppetIn, _ := puppetCmd.StdinPipe()
	puppetOut, _ := puppetCmd.StdoutPipe()

	// Start the process
	puppetCmd.Start()
	puppetIn.Close()

	// Read the resulting output
	puppetBytes, _ := io.ReadAll(puppetOut)

	// Wait for process to exit
	puppetCmd.Wait()

	// Printing the resulting output
	fmt.Println(string(puppetBytes))
}