package main

import (
	"bytes"
	"fmt"
	"github.com/alexflint/go-arg"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func main() {

	type AgentCmd struct {
		PuppetTest  bool   `arg:"--test"`
		Noop        bool   `arg:"--noop"`
		Op          bool   `arg:"--no-noop"`
		Tags        bool   `arg:"--tags"`
		TagsAdd     string `arg:"-a"`
		SkipTags    bool   `arg:"--skip_tags"`
		SkipTagsAdd string `arg:"-b"`
	}

	var args struct {
		Agent *AgentCmd `arg:"subcommand:agent"`
	}

	arg.MustParse(&args)

	var argu []string
	var pa string
	var pt string
	var ptags string
	var stags string
	var putest string

	//flag.Parse()

	// === Switch on args / flags that are called on runtime ===
	switch {
	case args.Agent != nil:
		pa = "agent"

		if args.Agent.PuppetTest == true {
			putest = "--test"
		}

		if args.Agent.Noop == true {
			pt = "--noop"
		} else if args.Agent.Op == true {
			pt = "--no-noop"
		}

		argu = []string{pa, putest, pt}

		if args.Agent.Tags == true {
			ptags = "--tags"
			argu = []string{pa, putest, pt, ptags, args.Agent.TagsAdd}
		}

		if args.Agent.SkipTags == true {
			stags = "--skip_tags"
			argu = []string{pa, putest, pt, stags, args.Agent.SkipTagsAdd}
		}
	}

	// == Command Executing ==
	// p := "/usr/local/bin/puppet"
	p := "/opt/puppetlabs/puppet/bin/puppet"
	// Windows Pathes needs to be in \
	pw := "C:/ProgramFiles/Puppet Labs/bin"
	cmd := exec.Command(p, argu...)
	if runtime.GOOS == "windows" {
		fmt.Println("Running on Windows:")
		cmd = exec.Command(pw, argu...)
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
