package main

import "fmt"

import "strings"

func main () {

    var userInput string = ""

    fmt.Printf ( "\nPlease Enter a string: " )
    fmt.Scan ( &userInput )
    
    fmt.Printf ( "String that it was took is: " + userInput + "\n" )

    var Lower_userInput string = strings.ToLower ( userInput )
    var   len_userInput    int =       len ( Lower_userInput )
    
    if len_userInput >= 3 && strings.HasPrefix ( Lower_userInput , "i" ) && 
                             strings.Contains  ( Lower_userInput , "a" ) {

        if "n" == Lower_userInput [ ( len_userInput - 1 ) : ] {

            fmt.Printf ( "Found!\n" ) ; return
        }
    }

    fmt.Printf ( "Not Found!\n" )
}

// user have to enter a string , and the code will show Found! or Not Found!

// If you really want to include the spaces ,
// you may consider using fmt.Scanf() with format %q a double-quoted string safely escaped with Go syntax.
// https://stackoverflow.com/questions/34647039/how-to-use-fmt-scanln-read-from-a-string-separated-by-spaces
