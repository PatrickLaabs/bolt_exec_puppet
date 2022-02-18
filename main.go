package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func main() {

	noopCmd := flag.NewFlagSet("noop", flag.ExitOnError)
	noopName := noopCmd.String("noop", "--noop", "puppet agent --noop")
	// noopEnable := noopCmd.Bool("enable", false, "enable")

	opCmd := flag.NewFlagSet("op", flag.ExitOnError)
	opName := opCmd.String("op", "--no-noop", "puppet agent --no-noop")
	// opEnable := opCmd.Bool("enable", false, "enable")

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	helpName := helpCmd.String("help", "", "-h")
	//
	// == ToDo:
	// == Value von tag Flag muss mit --tags=xxx befüllt werden. Also: tags-como=--tags=siguv_como
	// == Besser Lösung wird gesucht!
	//tagComoCmd := flag.NewFlagSet("tagcomo", flag.ExitOnError)
	//tagComoName := tagComoCmd.String("tags-como", "--tags=siguv_como", "puppet agent --tags=siguv_como")

	if len(os.Args) < 2 {
		fmt.Println(">> Usage:\n>> ./bolt_puppet_exec noop or op")
		os.Exit(1)
	}

	var n string
	flag.Parse()
	switch os.Args[1] {
	case "noop":
		noopCmd.Parse(os.Args[2:])
		//fmt.Println(" > enable noop:", *noopEnable)
		//if *noopEnable == false {
		//	fmt.Println("exiting noop case")
		//	os.Exit(1)
		//}
		n = *noopName
	case "op":
		opCmd.Parse(os.Args[2:])
		//fmt.Println(" > enable op:", *opEnable)
		//if *opEnable == false {
		//	fmt.Println("exiting op case")
		//	os.Exit(1)
		//}
		n = *opName
	case "help":
		helpCmd.Parse(os.Args[2:])
		fmt.Println(">> Usage:\n>> ./bolt_puppet_exec noop or op")
		n = *helpName
		os.Exit(1)
	}

	// == Command Executing ==
	// p := "/usr/local/bin/puppet"
	p := "/opt/puppetlabs/puppet/bin/puppet"
	pw := "puppet"
	pa := "agent"
	t := "--test"
	cmd := exec.Command(p, pa, t, n)
	if runtime.GOOS == "windows" {
		fmt.Println("Running on Windows:")
		cmd = exec.Command(pw, pa, t, n)
	}

	// Streaming Stderr and Stdout into a single Buffer
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b

	// Printing func printCommand
	printCommand(cmd)

	var waitStatus syscall.WaitStatus
	// Starting the command saved inside cmd
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println("CombinedOut:\n", string(b.Bytes()))
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
			exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		fmt.Println("CombinedOut:\n", string(b.Bytes()))
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
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
