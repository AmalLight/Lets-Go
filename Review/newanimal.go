package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalStruct struct {
	food, locomotion, noise string
}

type Cow AnimalStruct
type Bird AnimalStruct
type Snake AnimalStruct

func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
	fmt.Println(c.noise)
}
func (b Bird) Eat() {
	fmt.Println(b.food)
}
func (b Bird) Move() {
	fmt.Println(b.locomotion)
}
func (b Bird) Speak() {
	fmt.Println(b.noise)
}
func (s Snake) Eat() {
	fmt.Println(s.food)
}
func (s Snake) Move() {
	fmt.Println(s.locomotion)
}
func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		inputType, targetName, inputParam := GetInput(scanner)
		switch inputType {
		case "newanimal":
			CreateAnimal(targetName, inputParam, animals)
		case "query":
			QueryAnimal(targetName, inputParam, animals)
		default:
			fmt.Printf("Unknown inputType %s, must be one of (newanimal, query)\n", inputType)
			continue
		}
	}
}

func GetInput(scanner *bufio.Scanner) (string, string, string) {
	var inputType, targetName, inputParam string
	fmt.Printf("> ")
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	inputElems := strings.Split(input, " ")
	if len(inputElems) != 3 {
		fmt.Println("You must input exactly three strings, separated by a space")
		inputType = "unknown"
		targetName = "unknown"
		inputParam = "unknown"
	} else {
		inputType = inputElems[0]
		targetName = inputElems[1]
		inputParam = inputElems[2]

	}
	return inputType, targetName, inputParam
}

func CreateAnimal(targetName string, inputParam string, animals map[string]Animal) {
	_, present := animals[targetName]
	if present {
		fmt.Printf("Warning! Animal named %s already present, overwriting.\n", targetName)
	}
	switch inputParam {
	case "cow":
		animals[targetName] = Cow{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		animals[targetName] = Bird{food: "worms", locomotion: "fly", noise: "peep"}

	case "snake":
		animals[targetName] = Snake{food: "mice", locomotion: "slither", noise: "hiss"}
	default:
		fmt.Printf("Unknown animal type %s, skipping\n", inputParam)
		return
	}
	fmt.Println("Created it!")
}
func QueryAnimal(targetName string, inputParam string, animals map[string]Animal) {
	animalType, present := animals[targetName]
	if !present {
		fmt.Printf("Warning! Animal named %s not present, skpping.\n", targetName)
		return
	}
	switch inputParam {
	case "speak":
		animalType.Speak()
	case "eat":
		animalType.Eat()
	case "move":
		animalType.Move()
	default:
		fmt.Printf("Unknown action %s, skipping\n", inputParam)
	}

}
