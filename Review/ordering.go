package main

import (
	"fmt"
	"sort"
)

func main(){
	var numberElements, element int
  for numberElements<4 {
	  fmt.Printf("Insert the number of elements you want to add (minimum 4): ")

	  fmt.Scanf("%d", &numberElements)
	  if numberElements < 4 {
		  fmt.Println("You should insert a minimum of 4 int elements!")
	  }
}
	elementSlice := make([]int, 0)
	fmt.Println("Insert a integer and press enter: ")

	for i := 0; i < numberElements; i++ {
		fmt.Scanf("%d", &element)
		elementSlice = append(elementSlice, element)
	}

	fmt.Println("Ordering elements: ", elementSlice)
  
	var result []int
	ch := make(chan []int)
	step := numberElements / 4

	for i := 0; i < numberElements; i += step {
		go sortSlice(elementSlice[i:i+step], ch)
	}
  
	for i := 0; i < numberElements/step; i++ {
		partialResult := <-ch
		result = merge(result, partialResult)
	}

	fmt.Printf("Final sorted output %v\n", result)
}

func sortSlice(slice []int, ch chan []int) {
	fmt.Printf("Sort the input:%v \n", slice)
	sort.Ints(slice)
	fmt.Printf("Sorted output: %v \n", slice)
	ch <- slice
}

func merge(left, right []int) []int {

	length, i, j := len(left)+len(right), 0, 0
	slice := make([]int, length, length)

	for k := 0; k < length; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
