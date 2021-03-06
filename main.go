package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func readLines(fileToRead string) ([]string, error) {
	file, err := os.Open(fileToRead)
	if err != nil {
		fmt.Println("retard")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func symbolsArray(symbolsToUse string) []string {
	switch {
	case symbolsToUse == "minimal":
		return []string{"!", "@", "$", "%", "^", "&", "*", "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9"}

	case symbolsToUse == "all":
		return []string{"!", "@", "$", "%", "^", "&", "*", "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9", "`", "~", "#", "-", "=", "+", "[", "{", "]", "}"}
	}
	return nil
}

func main() {
	var password bytes.Buffer
	var symbols []string

	rand.Seed(time.Now().UTC().UnixNano())

	args := os.Args[1:]
	fmt.Println(args[0], args[1])

	words, err := readLines(args[0])
	if err != nil {
		fmt.Println("Something went wrong in words")
	}

	wordCount, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Something went wrong in wordCount")
	}

	if len(args) > 2 {
		symbols = symbolsArray(args[2])
	} else {
		symbols = symbolsArray("minimal")
	}

	for i := 0; i < wordCount; i++ {
		n := rand.Intn(len(words))
		n0 := rand.Intn(len(symbols))
		password.WriteString(words[n] + symbols[n0])
	}
	fmt.Println(password.String())
}
