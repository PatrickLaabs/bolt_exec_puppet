package main

import (
	"flag"
	"fmt"
	"log"
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
	flag.Var(&myFlags, "list1", "Some description for this param.")
	flag.Parse()

	fmt.Println("stored values:", myFlags)
	fmt.Println("first arg:", myFlags[1])

	cmd := exec.Command("go")
	if len(myFlags) == 1 {
		cmd = exec.Command(myFlags[0])
		fmt.Println("len == 1", myFlags)
	} else if len(myFlags) == 2 {
		cmd = exec.Command(myFlags[0], myFlags[1])
		fmt.Println("len == 2", myFlags)
	} else if len(myFlags) == 3 {
		cmd = exec.Command(myFlags[0], myFlags[1], myFlags[2])
		fmt.Println("len == 3", myFlags)
	}

	cmdOut, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmdOut))
}
