package main

import (
	"fmt"
	"os"

	"practice/prac1"
	"practice/prac2"
	"practice/prac3"
	"practice/prac4"
	"practice/prac5"
	"practice/prac6"
	"practice/prac7"
	"practice/practice_quiz"
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
	case "3":
		prac3.Run()
	case "4":
		prac4.Run()
	case "5":
		prac5.Run()
	case "6":
		prac6.Run()
	case "7":
		prac7.Run()
	case "practice":
		practice_quiz.Run()
	default:
		fmt.Println("Invalid practice number. Please choose 1, 2, or 3")
	}
}
