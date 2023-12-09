package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func isNumeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
	// calling regexp.MustCompile() function to create the regular expression.
	// calling MatchString() function that returns a bool that
	// indicates whether a pattern is matched by the string.
}

func hasLetterNumber(s string) (bool, string) {
	letterNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	lncLen := len(s)
	if lncLen >= 3 {
		for j := 0; j < 9; j++ {
			if strings.Contains(s, letterNumbers[j]) {
				//fmt.Print(j + 1)
				return true, strconv.Itoa(j + 1)
			}
		}
	}
	return false, ""
}

func main() {
	inputFile, err := os.Open("./inputs/1.txt")
	check(err)
	//fmt.Print(string(input))

	lines, err := ReadLines(inputFile)
	check(err)

	sum := 0

	for _, line := range lines {
		foundFirst := false
		foundLast := false
		firstDigit := ""
		lastDigit := ""
		letterNumberChain := ""
		for i := 0; i < len(line); i++ {
			if !foundFirst {
				if isNumeric(string(line[i])) {
					//fmt.Print(string(line[i]))
					firstDigit = string(line[i])
					foundFirst = true
					break
				} else {
					letterNumberChain += string(line[i])
					hasLN, ln := hasLetterNumber(letterNumberChain)
					if hasLN {
						//fmt.Print(ln)
						firstDigit = ln
						foundFirst = true
						break
					}
				}
			}
		}
		letterNumberChain = ""
		for i := len(line) - 1; i >= 0; i-- {
			if !foundLast {
				if isNumeric(string(line[i])) {
					//fmt.Print(string(line[i]))
					lastDigit = string(line[i])
					foundLast = true
					break
				} else {
					letterNumberChain = string(line[i]) + letterNumberChain
					//fmt.Print(letterNumberChain)
					hasLN, ln := hasLetterNumber(letterNumberChain)
					if hasLN {
						//fmt.Print(ln)
						lastDigit = ln
						foundLast = true
						break
					}
				}
			}

		}
		result := firstDigit + lastDigit
		intResult, err := strconv.Atoi(result)
		check(err)
		sum += intResult
		fmt.Print(line, " ", result)
		println()
	}

	fmt.Println(sum)
}
