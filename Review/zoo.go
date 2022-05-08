package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	m := make(map[string]Animal)
	m["cow"] = Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	m["bird"] = Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	m["snake"] = Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	for {
		animalName, animalAction := "", ""
		fmt.Println("Plase enter an animal name followed by an action")
		fmt.Scanf("%s %s", &animalName, &animalAction)

		switch animalAction {
		case "eat":
			m[animalName].Eat()
		case "move":
			m[animalName].Move()
		case "speak":
			m[animalName].Speak()
		default:
			fmt.Println("Invalid action name")
		}
	}

}
