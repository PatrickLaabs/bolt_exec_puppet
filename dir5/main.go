package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goName := goCmd.String("version", "version", "version")

	gitCmd := flag.NewFlagSet("git", flag.ExitOnError)
	gitName := gitCmd.String("version", "blurb", "version")

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	helpName := helpCmd.String("help", "", "-h")

	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	buildName := buildCmd.String("build", "build", "build")
	// buildN := buildCmd.String("main.go", "main.go", "main.go")

	// == SETTING ARGS ==
	// Args: --noop, --no-noop, --tags-como
	// --no-noop => https://puppet.com/docs/puppet/7/config_about_settings.html
	//noopCmd := flag.NewFlagSet("noop", flag.ExitOnError)
	//noopName := noopCmd.String("noop", "--noop", "puppet agent --noop")
	//
	//opCmd := flag.NewFlagSet("op", flag.ExitOnError)
	//opName := opCmd.String("op", "--no-noop", "puppet agent --no-noop")
	//
	// == ToDo:
	// == Value von tag Flag muss mit --tags=xxx befüllt werden. Also: tags-como=--tags=siguv_como
	// == Besser Lösung wird gesucht!
	//tagComoCmd := flag.NewFlagSet("tagcomo", flag.ExitOnError)
	//tagComoName := tagComoCmd.String("tags-como", "--tags=siguv_como", "puppet agent --tags=siguv_como")

	//if len(os.Args) < 2 {
	//	fmt.Println(">> Usage:\n>> ./bolt_puppet_exec noop, op or tags-como")
	//	os.Exit(1)
	//}
	if len(os.Args) < 2 {
		fmt.Println(">> Usage:\n>> ./main go or git")
		os.Exit(1)
	}

	var n string
	var t string
	// var t string
	flag.Parse()
	// args := flag.Args()
	switch os.Args[1] {
	case "go":
		goCmd.Parse(os.Args[2:])
		fmt.Println("  tail:", goCmd.Args())
		fmt.Println("  > tail:", flag.Args())
		//fmt.Println("subcommand 'go'")
		//tail := goCmd.Args()
		//tailConv := strings.Join(tail, " ")
		n = *goName
		//t = tailConv
	case "git":
		gitCmd.Parse(os.Args[2:])
		fmt.Println("  tail:", gitCmd.Args())
		fmt.Println("  > tail:", flag.Args())
		//fmt.Println("subcommand 'git'")
		// fmt.Println("gitname:", *gitName)
		//tail := gitCmd.Args()
		//tailConv := strings.Join(tail, " ")
		n = *gitName
		// t = tailConv
	case "build":
		buildCmd.Parse(os.Args[2:])
		fmt.Println("  tail:", buildCmd.Args())
		fmt.Println("  > tail:", flag.Args())
		// var tail []string = buildCmd.Args()
		//tail := buildCmd.Args()
		//tailConv := strings.Join(tail, " ")
		n = *buildName
		tail := flag.Args()
		tailConv := strings.Join(tail, " ")
		t = tailConv
		// t = tailConv
	case "help":
		helpCmd.Parse(os.Args[2:])
		fmt.Println(">> Usage:\n>> ./main go or git")
		//fmt.Println("helpName:", *helpName)
		n = *helpName
		os.Exit(1)
	}

	// Command Executing

	//p := "/usr/local/bin/puppet"
	//pa := "agent"
	//cmd := exec.Command(p, pa, n)
	// g := "go"
	cmd := exec.Command("go", n, t)
	//if runtime.GOOS == "windows" {
	//	cmd = exec.Command("tasklist")
	//}

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
