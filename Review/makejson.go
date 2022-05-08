package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	myName := getName(scanner)
	myAddress := getAddress(scanner)
	myMap := getMap(myName, myAddress)
	barr := getJson(myMap)
	fmt.Println(string(barr))
}

func getName(scanner *bufio.Scanner) string {
	fmt.Printf("Enter your name: ")
	scanner.Scan()
	inputName := scanner.Text()
	return inputName
}

func getAddress(scanner *bufio.Scanner) string {
	fmt.Printf("Enter your address: ")
	scanner.Scan()
	inputAddress := scanner.Text()
	return inputAddress
}

func getMap(name string, address string) map[string]string {
	myMap := make(map[string]string)
	myMap["name"] = name
	myMap["address"] = address
	return myMap
}

func getJson(myMap map[string]string) []byte {
	barr, _ := json.Marshal(myMap)
	return barr
}
