package main

/*
* A race condition occurs when two or more concurrent goroutines try to access and modify the same data.
* For example, if one goroutine tries to read a variable, meanwhile other goroutines are trying to update
* the value of the same variable.
 */

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func universe() {
	defer wg.Done()
	msg = "Hello Universe"
}

func cosmos() {
	defer wg.Done()
	msg = "Hello Cosmos"
}

/*
* In this code example, we are deploying two goroutines named universe and cosmos.
* Both goroutines are updating a variable named message to "Hello Universe" and "Hello Cosmos" accordingly.
* When we run this code, It will deploy the two goroutines.
* These goroutines eventually update the value of msg and join back to the main goroutine.
* Since two goroutines are updating the same variable in the above code, we cannot determine the output of the fmt.Println(msg), i.e,
* It will either print "Hello Universe" or "Hello Cosmos".
* Since we cannot determine the value stored in the msg variable, this will be a serious bug in our code that can break our software.
 */

func main() {

	msg = "Hello world"

	wg.Add(2)
	go universe()
	go cosmos()

	wg.Wait()
	fmt.Println(msg)

}

