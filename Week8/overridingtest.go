package main

import ( "fmt" )

// be carefull to use *_test.go, Doesn't work!!

type Test interface {

    Do ( x int )
}

type MyInt struct { i int }

func (M MyInt) made () { fmt.Println ( M.i ) } // Receiver requires with force a Struct?
                                               // No You can use also a simple type ( like done in Week7 )

func (M MyInt) Do ( x int ) { M.i = x ; M.made () } // it is required for polymorphism.

func main () {

    var M MyInt
    
    var T Test
    T = M
    
    T.Do ( 5 )
}
