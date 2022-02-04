package main

import (
	"fmt"
	"io"
	"os/exec"
)

// func puppet() {
// 	puppetCmd := exec.Command("puppet", "apply", "--noop", "--test", "--debug", "manifest/manifest.pp")

// 	puppetIn, _ := puppetCmd.StdinPipe()
// 	puppetOut, _ := puppetCmd.StdoutPipe()
// 	puppetCmd.Start()
// 	puppetIn.Write([]byte())
// }

func main() {

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
}
