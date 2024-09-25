package main

import (
	"fmt"
	"os"

	"practice/prac1"
	"practice/prac2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify which practice to run: 1, 2, or 3")
		return
	}

	switch os.Args[1] {
	case "1":
		prac1.Run()
	case "2":
		prac2.Run()
	default:
		fmt.Println("Invalid practice number. Please choose 1, 2, or 3")
	}
}
