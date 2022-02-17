package main

import (
	"fmt"
	"os"
)

// 	Return 0 and 2 as successful (0)
//	Return 1, 4 & 6 as error (1)

func main() {
	code := 0
	defer func() {
		os.Exit(code)
	}()

	e := "2"
	switch n := e; n {
	case "1":
		fmt.Printf("==> switch case 1 - exit code 1: %s\n", string(n))
		code = 1
	case "2":
		fmt.Printf("==> switch case 2 - exit code 2: %s\n", string(n))
		code = 0
	case "4":
		fmt.Printf("==> switch case 4 - exit code 4: %s\n", string(n))
		code = 1
	case "6":
		fmt.Printf("==> switch case 6 - exit code 6: %s\n", string(n))
		code = 1
	default:
		fmt.Printf("==> switch case default - exit code 0: %s\n", string(n))
		code = 0
	}
}
