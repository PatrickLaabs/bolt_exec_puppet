package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// This line is needed for testing purposes, to run puppet with noop on manifest.pp
	// puppetCmd := exec.Command("/bin/sh", "-c",  "puppet apply --noop --test --debug manifest/manifest.pp | grep -E 'exit'")

	// puppetCmd := exec.Command("/bin/sh", "-c", "sudo /usr/local/bin/puppet agent --test --noop")

	puppetCmd := exec.Command("sudo", "/usr/local/bin/puppet", "agent", "--test", "--noop")

	// puppetCmd := exec.Command("/bin/sh", "-c", "sudo /usr/local/bin/puppet agent --test --noop | grep -E 'exit'") <= detailed exit codes from puppet

	// puppetCmd := exec.Command("/bin/sh", "-c", "ls | grep -E Docker; echo 'Hello'")

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