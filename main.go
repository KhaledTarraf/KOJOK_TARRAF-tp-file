package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	data, err := os.Open("Hangman.txt") // Open the file

	if err != nil { // Check for errors
		fmt.Println(err) // If there is an error, print it
	}
	scanner := bufio.NewScanner(data) // Create a new scanner

	scanner.Split(bufio.ScanWords) // Split the file by words

	tab := make([]string, 0) // Create a slice of strings

	for scanner.Scan() { // Scan the file
		tab = append(tab, scanner.Text()) // Append the words to the slice
	}

	fmt.Println(tab[0])          // Print the first word
	fmt.Println(tab[len(tab)-1]) // Print the last word

	for index, a := range tab {
		if a == "before" { // If the word is "before"
			result, _ := strconv.Atoi(tab[index+1]) // Convert the next word to an integer
			fmt.Println(tab[result-1])              // Print the word at the index of the integer
		}
		if a == "now" { // If the word is "now"
			are := tab[index-1]               // Get the word before "now"
			result1 := int(are[1]) / len(tab) // Divide the second letter of the word by the length of the slice
			fmt.Println(tab[result1-1])       // Print the word at the index of the result
		}
	}

	min := 0                                                    // Set the minimum value for the random number
	max := 1000                                                 // Set the maximum value for the random number
	rand.Seed(time.Now().Unix())                                // Seed the random number generator
	fmt.Println("The random number is", rand.Intn(max-min)+min) // Print the random number

	data.Close() // Close the file
}
