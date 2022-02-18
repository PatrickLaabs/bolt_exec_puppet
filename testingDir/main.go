package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	// == Command Executing ==
	// p := "/usr/local/bin/puppet"
	//p := "/opt/puppetlabs/puppet/bin/puppet"
	//pw := "puppet"
	//pa := "agent"
	//t := "--test"

	cmd := exec.Command("go", "blurb")
	//if runtime.GOOS == "windows" {
	//	fmt.Println("Running on Windows:")
	//	cmd = exec.Command(pw, pa, t, n)
	//}

	// Maybe use a combinedOutput
	// Attaching to Stdout and Stderr
	//cmdOutput := &bytes.Buffer{}
	//cmdStderr := &bytes.Buffer{}
	//cmd.Stdout = cmdOutput
	//cmd.Stderr = cmdStderr
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	// https://www.sobyte.net/post/2021-06/go-os-exec-short-tutorial/

	// Printing func printCommand
	printCommand(cmd)
	var waitStatus syscall.WaitStatus
	// Starting the command saved inside cmd
	if err := cmd.Run(); err != nil {
		// printError(err)
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println("CombinedOut:\n", string(b.Bytes()))
			//fmt.Println("Stderr ==> ", cmd.Stderr)
			//fmt.Println("Stdout ==> ", cmd.Stdout)
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
			exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Println("CombinedOut:\n", string(b.Bytes()))
		//fmt.Println("Stderr ==> ", cmd.Stderr)
		//fmt.Println("Stdout ==> ", cmd.Stdout)
		printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}

func printCommand(cmd *exec.Cmd) {
	// Printing executed command. Just for knowing that the run has started.
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

//func printError(err error) {
//	if err != nil {
//		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
//	}
//}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func exitHandle(outs []byte) {
	// Handling the returned exit codes from exec.Command inside a switch statement
	// and returns the desired exit code via os.Exit with defer.
	code := 0
	defer func() {
		os.Exit(code)
	}()

	// Printing depends on case.
	// Currently it prints the used case and the expected exit code along with it.
	// This is only for debuging and tests - can be removed afterwards.
	switch n := string(outs); n {
	case "1":
		fmt.Printf("==> switch case 1 - exit code 1: %s\n", string(outs))
		code = 1
	case "2":
		fmt.Printf("==> switch case 2 - exit code 2: %s\n", string(outs))
		code = 0
	case "4":
		fmt.Printf("==> switch case 4 - exit code 4: %s\n", string(outs))
		code = 1
	case "6":
		fmt.Printf("==> switch case 6 - exit code 6: %s\n", string(outs))
		code = 1
	default:
		fmt.Printf("==> switch case default - exit code 0: %s\n", string(outs))
		code = 0
	}
}
