package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var myFlags arrayFlags

func main() {

	// puppet - agent, --test, --noop, --tags, --<module-name>

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goBuild := goCmd.String("build", "build", "build")
	goMain := goCmd.String("main", "", "main")
	goOpt := goCmd.String("o", "-o", "output value")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	puppetCmd := flag.NewFlagSet("agent", flag.ExitOnError)
	puppetNoop := puppetCmd.Bool("noop", true, "--noop")
	puppetOp := puppetCmd.Bool("op", false, "--op")
	puppetTags := puppetCmd.Bool("tags", true, "--tags")
	flag.Var(&myFlags, "tags", "--tags")
	fmt.Println("stored values:", myFlags)
	//puppetSkipTags := puppetCmd.String("skip_tags", "--skip_tags", "--skip_tags")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	var args []string
	var pnn string
	var ptags string
	switch os.Args[1] {

	case "go":
		goCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'go'")
		fmt.Println("> build:", *goBuild)
		fmt.Println("> main.go:", *goMain)
		fmt.Println("> goOpt:", *goOpt)
		args = []string{*goBuild, *goOpt, *goMain}
	case "agent":
		// gorun agent --test --noop --tags --skip_tags
		pa := "agent"
		pt := "--test"
		fmt.Printf("tags value %s\n", myFlags)
		if *puppetNoop == true {
			pnn = "--noop"
			fmt.Println("running in NO-operational mode")
		} else if *puppetOp == true {
			pnn = "--no-noop"
			fmt.Println("running in operational mode")
		}

		if *puppetTags == true {
			ptags = "--tags"
		}

		args = []string{pa, pt, pnn, *puppetTags}
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	p := "puppet"
	cmd := exec.Command(p, args...)

	fmt.Println(cmd)
	cmdOut, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmdOut))
}
