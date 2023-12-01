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

	scoreCounter := 0

	for fileScanner.Scan() {

		nextLine := fileScanner.Text()
		turn := strings.Split(nextLine, " ")

		switch turn[1] {
		case "Z":
			scoreCounter += 6 + determineMyResult(turn[0], 1)
		case "Y":
			scoreCounter += 3 + determineMyResult(turn[0], 2)
		case "X":
			scoreCounter += 0 + determineMyResult(turn[0], 3)
		default:
			fmt.Println("some prob with my turnout")
		}

		//fmt.Println(fileScanner.Text())
	}

	readFile.Close()
	fmt.Println(scoreCounter)
}

func determineMyResult(opp string, outcome int) int {
	oppRes := 0
	switch opp {
	case "A":
		oppRes = 1
	case "B":
		oppRes = 2
	case "C":
		oppRes = 3
	default:
		fmt.Println("some prob with my turnout")
	}

	if outcome == 2 {
		return oppRes
	}
	if oppRes == 3 && outcome == 1 || oppRes == 2 && outcome == 3 {
		return 1
	}
	if oppRes == 1 && outcome == 1 || oppRes == 3 && outcome == 3 {
		return 2
	}
	return 3

}
