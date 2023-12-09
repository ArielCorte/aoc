package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(f *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	inputFile, err := os.Open("./inputs/1.txt")
	check(err)
	//fmt.Print(string(input))

	lines, err := ReadLines(inputFile)
	check(err)

	for _, line := range lines {
	}
}
