package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func getFirstDigit(s string) (int, rune) {
	for i, char := range s {
		if isDigit(char) {
			return i, char
		}
	}
	return -1, 0
}

func getLastDigit(s string) (int, rune) {
	outputIndex := -1
	outputChar := rune(0)

	for i, char := range s {
		if isDigit(char) {
			outputChar = char
			outputIndex = i
		}
	}
	return outputIndex, outputChar
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		i, char := getFirstDigit(line)

		if i < 0 {
			log.Fatalf("Invalid input found: %s", line)
		}

		firstDigit := char
		secondDigit := firstDigit

		i, char = getLastDigit(line[i:])
		if i >= 0 {
			secondDigit = char
		}

		number := (int(firstDigit)-'0')*10 + (int(secondDigit) - '0')

		total += number

	}
	fmt.Printf("ANSWER: %d\n", total)
}
