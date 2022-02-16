package main

import (
	// "bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	//var goFlagString string
	//flag.StringVar(&goFlagString, "go-version", "version", "using go version command")
	//var gitFlagString string
	//flag.StringVar(&gitFlagString, "git-version", "version", "using git version command")

	//var tagsFlagString string
	//// tagsFlagString := ""
	//flag.StringVar(&tagsFlagString, "tags", "", "tags help menu")
	//fmt.Printf("stringVar content %T\n", tagsFlagString)
	//var a []string
	//if string(tagsFlagString) == " " {
	//	fmt.Printf("stringVar content %v\n", tagsFlagString)
	//	a = append(a, "version")
	//	// a = append(a, tagsFlagString)
	//} else {
	//	fmt.Println("help")
	//	a = append(a, "help")
	//}

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("==> Choose your option, please: \n")
	// scanner.Scan()
	// var n string
	// switch string(scanner.Text()) {
	// case "go":
	// 	n = goFlagString
	// 	break
	// case "git":
	// 	n = gitFlagString
	// 	break
	// }

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goName := goCmd.String("version", "", "version")
	gitCmd := flag.NewFlagSet("git", flag.ExitOnError)
	gitName := gitCmd.String("version", "", "version")

	if len(os.Args) < 2 {
		fmt.Println("expected 'go' or 'git' subcommands")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "go":
		goCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'go'")
		fmt.Println("goName:", *goName)
		n = goName
	case "git":
		gitCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'git'")
		fmt.Println("gitname:", *gitName)
		n = gitName
	}
	fmt.Println("git and/or go name printing: ", n)
	// Command Execute

	// cmd := exec.Command("/usr/local/bin/puppet", "agent", "--test", "--noop")
	// cmd := exec.Command("bash", "-c", "go version")
	// cmd := exec.Command(string(scanner.Text()), string(n))
	// var a []string
	// a = append(a, "version")
	// a[0] = "go"
	// a[0] = "version"
	// fmt.Println("array print:", a[0], a[1])
	// fmt.Println(a)
	cmd := exec.Command("go", "help")

	// Maybe use a combinedOutput
	// Attaching to Stdout and Stderr
	cmdOutput := &bytes.Buffer{}
	cmdStderr := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdStderr

	// Printing func printCommand
	printCommand(cmd)

	var waitStatus syscall.WaitStatus
	// Starting the command saved inside cmd
	if err := cmd.Run(); err != nil {
		printError(err)
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println("Stderr ==> ", cmd.Stderr)
			fmt.Println("Stdout ==> ", cmd.Stdout)
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
			exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Println("Stderr ==> ", cmd.Stderr)
		fmt.Println("Stdout ==> ", cmd.Stdout)
		printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}

func printCommand(cmd *exec.Cmd) {
	// Printing executed command. Just for knowing that the run has started.
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

//func cmdArgsHandle() {
//
//}
