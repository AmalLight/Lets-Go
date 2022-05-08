package main

import "fmt"

import "strconv"

func main () {

    var userInput float64 = 0.0
    var userInput_str string = ""
    
    var err error = nil
    var userInput_int int = 0

    fmt.Printf ( "\nPlease Enter a floating point number: " )
    fmt.Scan ( &userInput )

    userInput_str = strconv.FormatFloat ( userInput , 'f' , 0 , 64 )

    userInput_int , err = strconv.Atoi ( userInput_str )
    
    if ( err != nil ) { 
    
        fmt.Printf ( "Error happened with Atoi \n" )
        
        fmt.Printf ( "Truncated value inserted before as a string : " + userInput_str + "\n" )

    } else {
    
        fmt.Printf ( "Truncated value inserted before as an integer : " + strconv.Itoa ( userInput_int ) + "\n" )
    }
}

// user have to enter a floating point number , and the code will show the truncated version of the floating point number
