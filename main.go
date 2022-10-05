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

	data, err := os.Open("Hangman.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(data)

	scanner.Split(bufio.ScanWords)

	tab := make([]string, 0)

	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}

	fmt.Println(tab[0])
	fmt.Println(tab[len(tab)-1])

	for index, a := range tab {
		if a == "before" {
			result, _ := strconv.Atoi(tab[index+1])
			fmt.Println(tab[result-1])
		}
		if a == "now" {
			are := tab[index-1]
			result1 := int(are[1]) / len(tab)
			fmt.Println(tab[result1-1])
		}
	}

	min := 0
	max := 1000
	rand.Seed(time.Now().Unix())
	fmt.Println("The random number is", rand.Intn(max-min)+min)

	data.Close()
}
