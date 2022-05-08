package main

import "fmt"

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5*a)*(t*t) + v0*t + s0
	}
	return fn
}

func main() {

	var a float64
	fmt.Print("Please enter the acceleration: ")
	fmt.Scanf("%f", &a)
	var v0 float64
	fmt.Print("Please enter the initial velocity: ")
	fmt.Scanf("%f", &v0)
	var s0 float64
	fmt.Print("Please enter the initial displacement: ")
	fmt.Scanf("%f", &s0)

	fn := GenDisplaceFn(a, v0, s0)

	var t float64
	fmt.Print("Please enter a time when you would like the displacement calculated: ")
	fmt.Scanf("%f", &t)

	fmt.Println(fn(t))

	fmt.Print("Please enter another time: ")
	fmt.Scanf("%f", &t)

	fmt.Println(fn(t))

}
