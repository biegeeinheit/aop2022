package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	maxCalories := 0
	secondMaxCalories := 0
	thirdMaxCalories := 0
	calorieCounter := 0

	for fileScanner.Scan() {

		nextLine := fileScanner.Text()
		if nextLine == "" { //possible problem at EOF
			if calorieCounter <= thirdMaxCalories {
				calorieCounter = 0
				continue
			}
			if calorieCounter > thirdMaxCalories && calorieCounter <= secondMaxCalories {
				thirdMaxCalories = calorieCounter
			} else {
				if calorieCounter > secondMaxCalories && calorieCounter <= maxCalories {
					thirdMaxCalories = secondMaxCalories
					secondMaxCalories = calorieCounter
				} else {
					if calorieCounter > maxCalories {
						thirdMaxCalories = secondMaxCalories
						secondMaxCalories = maxCalories
						maxCalories = calorieCounter
					}
				}

			}

			calorieCounter = 0
		} else {

			if number, err := strconv.ParseUint(nextLine, 10, 32); err == nil {
				finalIntNum := int(number) //Convert uint64 To int
				calorieCounter += finalIntNum
			}
		}
		//fmt.Println(fileScanner.Text())
	}

	readFile.Close()
	fmt.Println(maxCalories + secondMaxCalories + thirdMaxCalories)
}
