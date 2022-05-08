package main

import ( "fmt" )

func swap ( sli [] int , i1 int ) {

    tmp := sli [ i1 ]

    sli [ i1     ] = sli [ i1 + 1 ]
    sli [ i1 + 1 ] = tmp
}

func BubbleSort ( sli []int ) {

    n := len ( sli )
    
    // cost := ( n * ( n/2 - 1 ) + n/2 ) < n^2

    for i1 := 0 ; i1 <   n            ; i1 ++ {
    for i2 := 0 ; i2 < ( n - 1 - i1 ) ; i2 ++ {

           element_1 , element_2 := sli [ i2 ] , sli [ i2+1 ]
        
        if element_1 > element_2 { swap ( sli , i2 ) }
    }}
}

func main () {

    const size = 10
    var userIntegers [ size ] int // FOR more safety
    
    for i := 0 ; i < size ; i++ { // FOR more safety
    
        fmt.Printf ( "\n Please Enter an Integer for the index %d for a max of %d: " , i , size ) // FOR debug
        
        fmt.Scan ( &userIntegers [ i ] )
    }

    fmt.Printf ( "\n\n" )

    var intSlice [] int
    for i := 0 ; i < size ; i++ { intSlice = append ( intSlice , userIntegers [ i ] ) }

    for i := 0 ; i < size ; i++ {

        fmt.Printf ( " Origin value %d found in position %d \n" , intSlice [ i ] , i )
    }

    fmt.Printf ( "\n" )

    BubbleSort ( intSlice )

    for i := 0 ; i < size ; i++ {

        fmt.Printf ( " Sorted value %d found in position %d \n" , intSlice [ i ] , i )
    }
}

// bubble.txt for cost function info in my GitHub
// input suggested: 9 8 7 6 5 4 3 2 1 0
