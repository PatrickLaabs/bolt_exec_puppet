package main

import (
	"bytes"
	"fmt"
	"github.com/alexflint/go-arg"
	"os"
	"os/exec"
	"runtime"
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
		fmt.Println("Version: v1.0.0\n\n" +
			"Options:\n" +
			"bolt_exec_puppet --help\n\n" +
			"General usage:\n" +
			"bolt_exec_puppet agent [--test] [--noop] [--tags TAGS] [--skip_tags SKIP_TAGS]\n\n" +
			"Exit Codes:\n" +
			"  Puppet Exit Codes 0, 2 are handled as Exit Code 0\n" +
			"  Puppet Exit Codes 1, 4, 6 are handled as Exit Code 1\n\n" +
			"Some examples:\n" +
			"  bolt_exec_puppet agent --test\n" +
			"  bolt_exec_puppet agent --test --noop\n\n" +
			"  bolt_exec_puppet agent --test --noop --tags=<module>\n" +
			"  bolt_exec_puppet agent --test --noop --tags <module>\n\n" +
			"  bolt_exec_puppet agent --test --noop --skip_tags=<module>\n" +
			"  bolt_exec_puppet agent --test --noop --skip_tags <module>\n\n" +
			"  bolt_exec_puppet agent --test --noop --tags=<module> --skip_tags=<module>")
		os.Exit(1)
	}

	arg.MustParse(&args)

	var argu []string
	var pa string
	var pt string
	var putest string

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
	// pw := "C:\\Program Files\\Puppet Labs\\Puppet\\bin\\puppet"
	pw := "puppet"
	cmd := exec.Command(p, argu...)
	if runtime.GOOS == "windows" {
		fmt.Println("Running on Windows:")
		cmd = exec.Command(pw, argu...)
	}

	// Streaming Stderr and Stdout into a single Buffer
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b

	var waitStatus syscall.WaitStatus
	// Starting the command saved inside cmd
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println(b.String())
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		fmt.Println(b.String())
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitHandle([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}

func exitHandle(outs []byte) {
	// Handling the returned exit codes from exec.Command inside a switch statement
	// and returns the desired exit code via os.Exit with defer.
	code := 0
	defer func() {
		os.Exit(code)
	}()
	switch n := string(outs); n {
	case "1":
		code = 1
	case "2":
		code = 0
	case "4":
		code = 1
	case "6":
		code = 1
	default:
		code = 0
	}
}
