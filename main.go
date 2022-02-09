package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func fooOutput(outs []byte) {
	// if len(outs) > 0 {
	// 	fmt.Printf("==> fooOutput: %s\n", string(outs))
	// }
	switch len(outs) {
	case 1:
		fmt.Printf("==> switch case 1: %s\n", string(outs))
	case 2:
		fmt.Printf("==> switch case 2: %s\n", string(outs))
	}

}

func main() {
	cmd := exec.Command("go", "blurb")
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		printError(err)
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
			fooOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		fooOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}
