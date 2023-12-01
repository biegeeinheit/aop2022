package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	priorityCounter := 0

	for fileScanner.Scan() {

		nextLine := fileScanner.Text()

		h1 := nextLine[0 : len(nextLine)/2]
		h2 := nextLine[len(nextLine)/2:]

		res := SimpleGeneric[string](strings.Split(h1, ""), strings.Split(h2, ""))
		//fmt.Printf(res[0])
		b := mapCharacters(res[0])
		priorityCounter += b
		fmt.Printf(res[0])
		fmt.Println("aaaaaaaaaa")
		fmt.Println(h1)
		fmt.Println(h2)
	}

	readFile.Close()
	fmt.Println(priorityCounter)
}

func mapCharacters(char string) int {
	// Check if the string has at least one character
	if len(char) == 0 {
		return 0 // or handle the empty string case as needed
	}

	charRune := []rune(char)[0] // Convert the string to a rune

	mapping := map[rune]int{
		'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35,
		'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44,
		'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52,
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9,
		'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18,
		's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
	}

	return mapping[charRune]
}

func SimpleGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}
