package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"log"
	"os/exec"
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

	p := "puppet"
	cmd := exec.Command(p, argu...)

	fmt.Println(cmd)
	cmdOut, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmdOut))
}
