package main

import "fmt"

var version = "dev"

func main() {
	fmt.Printf("Version: %s\n", version)
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello Golang"
}
