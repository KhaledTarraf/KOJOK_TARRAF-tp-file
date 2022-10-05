package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("Hangman.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
