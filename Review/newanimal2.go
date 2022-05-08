package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, locomotion, noise string
}

func NewCow(name string) Animal {
	return &Cow{name: name, food: "grass", locomotion: "walk", noise: "moo"}
}

func (a Cow) Eat() {
	fmt.Printf("eats %s\n", a.food)
}

func (a Cow) Move() {
	fmt.Printf("%s \n", a.locomotion)
}
func (a Cow) Speak() {
	fmt.Printf("says %s \n", a.noise)
}

type Bird struct {
	name, food, locomotion, noise string
}

func NewBird(name string) Animal {
	return &Bird{name: name, food: "worms", locomotion: "fly", noise: "peep"}
}

func (a Bird) Eat() {
	fmt.Printf("eats %s\n", a.food)
}

func (a Bird) Move() {
	fmt.Printf("%s \n", a.locomotion)
}
func (a Bird) Speak() {
	fmt.Printf("says %s \n", a.noise)
}

type Snake struct {
	name, food, locomotion, noise string
}

func NewSnake(name string) Animal {
	return &Snake{name: name, food: "mice", locomotion: "slither", noise: "hsss"}
}

func (a Snake) Eat() {
	fmt.Printf("eats %s\n", a.food)
}

func (a Snake) Move() {
	fmt.Printf("%s \n", a.locomotion)
}
func (a Snake) Speak() {
	fmt.Printf("says %s \n", a.noise)
}

var methods map[string]func(a Animal) = map[string]func(a Animal){
	"eat":   func(a Animal) { a.Eat() },
	"speak": func(a Animal) { a.Speak() },
	"move":  func(a Animal) { a.Move() },
}

var Animals map[string]func(name string) Animal = map[string]func(name string) Animal{
	"cow":   NewCow,
	"bird":  NewBird,
	"snake": NewSnake,
}

const exitChar string = "x"

var uInputStr string

func main() {

	m := make(map[string]Animal)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")

		uInputStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		uInputStr = strings.ToLower(strings.TrimSpace(uInputStr))

		if (uInputStr == exitChar) {
			fmt.Println("Bye!")
			break
		}

		requestSlice := strings.Split(uInputStr, " ")
		if len(requestSlice) < 2 {
			fmt.Println("Wrong request, try again!")
			continue
		}

		if (uInputStr == exitChar) {
			fmt.Println("Bye!")

			break
		}

		if requestSlice[0] == "newanimal" {
			anim := Animals[requestSlice[2]](requestSlice[1])
			m[requestSlice[1]] = anim
			fmt.Println("Created it!")
			continue
		}

		if requestSlice[0] == "query" {
			//fmt.Println(m[requestSlice[1]])
			methods[requestSlice[2]](m[requestSlice[1]])
			//fmt.Println(m[requestSlice[1]])
		}

	}

}
