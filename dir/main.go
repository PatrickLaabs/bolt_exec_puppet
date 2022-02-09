package main

import "fmt"

func main() {
	e := "6"
	switch n := e; n {
	case "1":
		fmt.Printf("==> switch case 1 - exit code 1: %s\n", string(n))
	case "2":
		fmt.Printf("==> switch case 2 - exit code 2: %s\n", string(n))
	case "4":
		fmt.Printf("==> switch case 4 - exit code 4: %s\n", string(n))
	case "6":
		fmt.Printf("==> switch case 6 - exit code 6: %s\n", string(n))
	default:
		fmt.Printf("==> switch case default - exit code 0: %s\n", string(n))
	}
}

// ToDo:
// 	Return 0 and 2 as successful
//	Return 1, 4 & 6 as error
