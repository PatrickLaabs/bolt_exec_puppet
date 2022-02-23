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
		PuppetTest bool   `arg:"--test"`
		Noop       bool   `arg:"--noop"`
		Tags       string `arg:"--tags"`
		SkipTags   string `arg:"--skip_tags"`
	}

	var args struct {
		Agent *AgentCmd `arg:"subcommand:agent"`
	}

	if len(os.Args) < 3 {
		fmt.Println("General usage:\n" +
			"bolt_exec agent [--test] [--noop] [--tags TAGS] [--skip_tags SKIP_TAGS]\n\n" +
			"Some examples:\n" +
			"  ./bolt_exec_puppet agent --test\n" +
			"  ./bolt_exec_puppet agent --test --noop\n\n" +
			"  ./bolt_exec_puppet agent --test --noop --tags=<module>\n" +
			"  ./bolt_exec_puppet agent --test --noop --tags <module>\n\n" +
			"  ./bolt_exec_puppet agent --test --noop --skip_tags=<module>\n" +
			"  ./bolt_exec_puppet agent --test --noop --skip_tags <module>\n\n" +
			"A combination of both --tags and --skip_tags is also possible:\n" +
			"  ./bolt_exec_puppet agent --test --noop --tags=<module> --skip_tags=<module>\n\n" +
			"Every possibility can be run without --noop")
		os.Exit(1)
	}

	arg.MustParse(&args)

	var argu []string
	var pa string
	var pt string
	var putest string

	//flag.Parse()

	// === Switch on args / flags that are called on runtime ===
	switch {
	case args.Agent != nil:
		pa = "agent"

		if args.Agent.PuppetTest == true {
			putest = "--test"
		}

		argu = []string{pa, putest}

		if args.Agent.Noop == true {
			pt = "--noop"
			argu = []string{pa, putest, pt}
		}

		//if args.Agent.Tags == true {
		//	ptags = "--tags"
		//	argu = []string{pa, putest, pt, ptags, args.Agent.TagsAdd}
		//}

		if args.Agent.SkipTags != "" && args.Agent.Noop == true {
			data := args.Agent.SkipTags
			response := fmt.Sprintf("--skip_tags=%s", data)
			argu = []string{pa, putest, pt, response}
		} else if args.Agent.SkipTags != "" && args.Agent.Noop != true {
			data := args.Agent.SkipTags
			response := fmt.Sprintf("--skip_tags=%s", data)
			argu = []string{pa, putest, response}
		}

		if args.Agent.Tags != "" && args.Agent.Noop == true {
			data := args.Agent.Tags
			response := fmt.Sprintf("--tags=%s", data)
			argu = []string{pa, putest, pt, response}
		} else if args.Agent.Tags != "" && args.Agent.Noop != true {
			data := args.Agent.Tags
			response := fmt.Sprintf("--tags=%s", data)
			argu = []string{pa, putest, response}
		}

		if args.Agent.SkipTags != "" && args.Agent.Tags != "" && args.Agent.Noop == true {
			dataTags := args.Agent.Tags
			resTags := fmt.Sprintf("--tags=%s", dataTags)
			dataSkip := args.Agent.SkipTags
			resSkip := fmt.Sprintf("--skip_tags=%s", dataSkip)
			argu = []string{pa, putest, pt, resTags, resSkip}
		} else if args.Agent.SkipTags != "" && args.Agent.Tags != "" && args.Agent.Noop != true {
			dataTags := args.Agent.Tags
			resTags := fmt.Sprintf("--tags=%s", dataTags)
			dataSkip := args.Agent.SkipTags
			resSkip := fmt.Sprintf("--skip_tags=%s", dataSkip)
			argu = []string{pa, putest, resTags, resSkip}
		}

	}

	// == Command Executing ==
	p := "/usr/local/bin/puppet"
	// p := "/opt/puppetlabs/puppet/bin/puppet"
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
