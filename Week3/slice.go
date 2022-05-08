package main

import ( "fmt" ; "strconv" )

func main () {

    var size , i int = 3 , 1

    sli := make ( []int , 0 , size )

    for true {

        var userInput_str string = ""

        fmt.Printf ( "\nPlease Enter one integer: " ) // FOR debug
        
        fmt.Scan ( &userInput_str )
        
        // fmt.Printf ( "String value inserted before is: %s \n" , userInput_str ) // FOR debug
        
        var inputUser_int int = 0
        var err error = nil
        
        if userInput_str == "X" { 

            // fmt.Printf ( "\nX Pressed\n" ) // FOR debug

            break ;
        }
        
        inputUser_int , err = strconv.Atoi ( userInput_str )

        if ( err != nil ) { 
    
            fmt.Printf ( "\nError happened with Atoi \n" )

        } else {
    
            // fmt.Printf ( "Integer value converted is: %d \n" , inputUser_int ) // FOR debug
        
            sli = append ( sli , inputUser_int )
        
            // fmt.Printf ( "Slice in loop %d state is: %d \n" , i , sli  ) // FOR debug
        
            length := len ( sli )
        
            for i2 := 0 ; i2 < length ; i2 ++ {

                var copy_sli_i2 int = sli [ i2 ]

                var minus_index int = i2
                var minus int = sli [ i2 ]
            
                for i3 := i2+1 ; i3 < length ; i3 ++ {

                    if sli [ i3 ] < minus { minus_index = i3 ; minus = sli [ i3 ] }
                }
            
                if minus_index != i2 { 
            
                    sli [ i2 ] = minus
                    sli [ minus_index ] = copy_sli_i2
                }
            }
        }

        fmt.Printf ( "\nSorted Slice in loop %d state is: %d \n" , i , sli ) // FOR debug
        i = i+1
    }
    
    // fmt.Printf ( "Final Slice state is: %d \n" , sli ) // FOR debug
}

// https://pkg.go.dev/fmt
// The program takes strings from user input. For each input it checks if the input is equal to 'X' then break ; otherwise, if possible it converts the string to an integer, then adds this integer into a list. after that the code will sort the global list. All work will be done for each input.
