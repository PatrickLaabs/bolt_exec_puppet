package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// func printCommand(cmd *exec.Cmd) {
// 	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
// }

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

	// 	Exit Codes
	// 0 & 2 als 0 zurückgeben (success)
	// 1 & 4 & 6 als 1 zurückgeben (err)

	switch n := string(outs); n {
	case "1":
		fmt.Printf("==> switch case 1 - exit code 1: %s\n", string(outs))
	case "2":
		fmt.Printf("==> switch case 2 - exit code 2: %s\n", string(outs))
	case "4":
		fmt.Printf("==> switch case 4 - exit code 4: %s\n", string(outs))
	case "6":
		fmt.Printf("==> switch case 6 - exit code 6: %s\n", string(outs))
	default:
		fmt.Printf("==> switch case default - exit code 0: %s\n", string(outs))
	}

	// 	// switch len(outs) {
	// 	// case 1: // should be exit code 0
	// 	// 	fmt.Printf("==> switch case 1 - exit code: %s\n", string(outs))
	// 	// case 2: // should be exit code 1
	// 	// 	fmt.Printf("==> switch case 2 - exit code: %s\n", string(outs))
	// 	// }

	// 	// i := len(outs)
	// 	// switch i {
	// 	// case 0:
	// 	// 	fmt.Printf("==> switch case 0 - exit code / should be 0: %s\n", string(outs))
	// 	// case 1: // should be exit code 1
	// 	// 	fmt.Printf("==> switch case 1 - exit code / should be 1: %s\n", string(outs))
	// 	// case 2: // should be exit code 2
	// 	// 	if len(outs) == 2 {
	// 	// 		fmt.Printf("==> switch case 2 - exit code / should be 2: %s\n", string(outs))
	// 	// 	}
	// 	// default:
	// 	// 	if len(outs) == 0 {
	// 	// 		fmt.Printf("==> switch case default - exit code / should be 0: %s\n", string(outs))
	// 	// 	}
	// 	// }

}

func main() {
	cmd := exec.Command("go", "version")
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
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
