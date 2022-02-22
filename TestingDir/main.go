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

	// puppet agent --test --noop --tags --skip_tags

	// Build the Flags like puppet would handle them originally.
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goName := goCmd.String("version", "version", "go version")

	buildCmd := flag.NewFlagSet("build", flag.ContinueOnError)
	buildName := buildCmd.String("build", "build", "go build")
	buildAdd := buildCmd.String("add", "main.go", "main.go")
	buildBool := buildCmd.Bool("noop", true, "bool for noop")

	// === Setting up flags ===
	// ./bolt_exec noop
	noopCmd := flag.NewFlagSet("noop", flag.ExitOnError)
	noopName := noopCmd.String("noop", "--noop", "puppet agent --noop")

	// ./bolt_exec op
	opCmd := flag.NewFlagSet("op", flag.ExitOnError)
	opName := opCmd.String("op", "--no-noop", "puppet agent --no-noop")

	// ./bolt_exec help
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	helpName := helpCmd.String("help", "", "-h")

	// ./bolt_exec tags -add=<module> -start=--noop
	tagsCmd := flag.NewFlagSet("tags", flag.ExitOnError)
	tagsName := tagsCmd.String("tags", "--tags", "puppet agent --tags=")
	tagsStart := tagsCmd.String("start", "--noop", "choose between op or noop")
	tagsAdd := tagsCmd.String("add", "", "additional args like your module name")

	// ./bolt_exec skip -add=<module> -start=--noop
	skipTagsCmd := flag.NewFlagSet("skip_tags", flag.ExitOnError)
	skipTagsName := skipTagsCmd.String("skip", "--skip_tags", "skipping tags")
	skipTagsStart := skipTagsCmd.String("start", "--noop", "choose between op or noop")
	skipTagsAdd := skipTagsCmd.String("add", "", "module name to skip")

	// === Err checking, since we need at least 2 args ===
	// may be deleted if not needed
	if len(os.Args) < 2 {
		fmt.Println(">> Usage:\n./puppet_bolt_exec noop \n" +
			"./puppet_bolt_exec op\n" +
			"./puppet_bolt_exec help\n" +
			"./puppet_bolt_exec tags -add=<module> -start=--noop\n" +
			"./puppet_bolt_exec skip -add=<module> -start=--noop")
		os.Exit(1)
	}

	// === Init var's for usage out of scope ===
	var n string
	var nn string
	var nm string
	var nb bool
	var nbn string
	var args []string
	// === Parsing the flags ===
	flag.Parse()

	// === Switch on args / flags that are called on runtime ===
	switch os.Args[1] {
	case "noop":
		pa := "agent"
		t := "--test"
		noopCmd.Parse(os.Args[2:])
		n = *noopName
		args = []string{pa, t, n}
		// if tags (n) == empty remove n from args []string
	case "op":
		pa := "agent"
		t := "--test"
		opCmd.Parse(os.Args[2:])
		n = *opName
		args = []string{pa, t, n}
	case "go":
		goCmd.Parse(os.Args[2:])
		n = *goName
		args = []string{n}
	case "build":
		buildCmd.Parse(os.Args[2:])
		n = *buildName
		nm = *buildAdd
		nb = *buildBool
		if nb == true {
			fmt.Println("bool is false")
			nbn = "--noop"
		} else if nb != true {
			fmt.Println("bool is true")
			nbn = "--no-noop"
		}
		if nm == "" {
			fmt.Println("no additional args set")
			args = []string{n, nbn}
			fmt.Printf("args passed: %s\n", args)
		} else {
			fmt.Printf("Args set to: %s\n", nm)
			args = []string{n, nm, nbn}
			fmt.Printf("args passed: %s\n", args)
		}
	case "tags":
		pa := "agent"
		t := "--test"
		tagsCmd.Parse(os.Args[2:])
		n = *tagsName
		nn = *tagsAdd
		nm = *tagsStart
		args = []string{pa, t, nm, n, nn}
	case "skip":
		pa := "agent"
		t := "--test"
		skipTagsCmd.Parse(os.Args[2:])
		n = *skipTagsName
		nn = *skipTagsAdd
		nm = *skipTagsStart
		args = []string{pa, t, nm, n, nn}
	case "help":
		helpCmd.Parse(os.Args[2:])
		fmt.Println(">> Usage:\n./puppet_bolt_exec noop \n" +
			"./puppet_bolt_exec op\n" +
			"./puppet_bolt_exec help\n" +
			"./puppet_bolt_exec tags -add=<module> -start=--noop\n" +
			"./puppet_bolt_exec skip -add=<module> -start=--noop")
		n = *helpName
		os.Exit(1)
	}

	// == Command Executing ==
	// p := "/usr/local/bin/puppet"
	// p := "/opt/puppetlabs/puppet/bin/puppet"
	// Windows Pathes needs to be in \
	//pw := "C:/ProgramFiles/Puppet Labs/bin"
	//cmd := exec.Command(p, args...)
	//if runtime.GOOS == "windows" {
	//	fmt.Println("Running on Windows:")
	//	cmd = exec.Command(pw, args...)
	//}

	g := "go"
	cmd := exec.Command(g, args...)

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
			// fmt.Println("CombinedOut:\n", string(b.Bytes()))
			fmt.Println("CombinedOut:\n", b.String())
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
			exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		// fmt.Println("CombinedOut:\n", string(b.Bytes()))
		fmt.Println("CombinedOut:\n", b.String())
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}

func printCommand(cmd *exec.Cmd) {
	// Printing executed command. Just for knowing that the run has started.
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
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