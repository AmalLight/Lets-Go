package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var number string
	chosenNumbers := make([]int, 3)
	done := false
	i := 0

	for !done {
		fmt.Println("Enter a number")
		fmt.Scan(&number)

		if number == "X" {
			done = true
			fmt.Println("Here's your result...")
			fmt.Println(chosenNumbers)
			break
		}

		result, err := strconv.Atoi(number)
		if err != nil {
			fmt.Printf("Input must be a number, %s \n", err.Error())
			break
		}

		switch {
		case i < 3:
			chosenNumbers[i] = result
			tmpSlice := sortOnly(chosenNumbers, i)
			chosenNumbers = tmpSlice
			fmt.Println(chosenNumbers)
		default:
			tmpSlice := insertAndSort(chosenNumbers, result)
			chosenNumbers = tmpSlice
			fmt.Println(chosenNumbers)
		}

		i++
	}
}

func insertAndSort(chosenNumbers []int, numberToInsert int) []int {
	tmpSlice := make([]int, len(chosenNumbers), len(chosenNumbers)+1)
	tmpSlice = append(chosenNumbers, numberToInsert)
	sort.Slice(tmpSlice, func(i, j int) bool {
		return tmpSlice[i] < tmpSlice[j]
	})

	return tmpSlice
}

func sortOnly(sliceToSort []int, position int) []int {

	if position == 1 {
		if sliceToSort[1] < sliceToSort[0] {
			tmp := sliceToSort[0]
			sliceToSort[0] = sliceToSort[1]
			sliceToSort[1] = tmp
			return sliceToSort
		}

		return sliceToSort
	}

	if position == 2 {
		if sliceToSort[2] < sliceToSort[1] {
			tmp := sliceToSort[1]
			sliceToSort[1] = sliceToSort[2]
			sliceToSort[2] = tmp
			return sliceToSort
		}

		return sliceToSort
	}

	return sliceToSort

}
