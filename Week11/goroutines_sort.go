package main

import ( "fmt" ; "bufio" ; "os" ; "unicode" ; "strings" ; "strconv" ; "sync" )

func swap ( sli [] int , i1 int ) {

    tmp := sli [ i1 ]

    sli [ i1     ] = sli [ i1 + 1 ]
    sli [ i1 + 1 ] = tmp
}

func BubbleSort ( sli []int , wg *sync.WaitGroup ) {

    n := len ( sli )
    
    // cost := ( n * ( n/2 - 1 ) + n/2 ) < n^2

    for i1 := 0 ; i1 <   n            ; i1 ++ {
    for i2 := 0 ; i2 < ( n - 1 - i1 ) ; i2 ++ {

           element_1 , element_2 :=       sli [ i2 ] , sli [ i2+1 ]        
        if element_1 > element_2 { swap ( sli , i2                ) }
    }}

    fmt.Println ( "BubbleSort sort: " , sli )
    
    wg.Done ()
}

func sort ( sli1 [] int , sli2 [] int ) [] int {

    var sli_copy [] int
    
    i ,       len1   , j ,       len2 :=
    0 , len ( sli1 ) , 0 , len ( sli2 )
    
    for ; ( i < len1 ) && ( j < len2 ) ; {
    
        if sli1 [ i ] <= sli2 [ j ] {
        
                 sli_copy = append ( sli_copy , sli1 [ i ] ) ; i++
        } else { sli_copy = append ( sli_copy , sli2 [ j ] ) ; j++ }
    }

    for ; ( i < len1 ) ; i++ { sli_copy = append ( sli_copy , sli1 [ i ] ) }
    for ; ( j < len2 ) ; j++ { sli_copy = append ( sli_copy , sli2 [ j ] ) }
    
    return sli_copy
}

func merge ( sli1 [] int , sli2 [] int , sli3 [] int , sli4 [] int ) [] int {

    sli1_copy := sort ( sli1      , sli2      )
    sli2_copy := sort ( sli3      , sli4      )
    sli3_copy := sort ( sli1_copy , sli2_copy )
    
    fmt.Println ( ""                          )
    fmt.Println ( "Merge1 sort: " , sli1_copy )
    fmt.Println ( "Merge2 sort: " , sli2_copy )
    fmt.Println ( "Merge3 sort: " , sli3_copy )

    return sli3_copy
}

func main () {

    scanner := bufio.NewScanner ( os.Stdin )

    fmt.Print ( "\n" , "enter a unique big series: " )
    
    scanner.Scan ()
    text :=  scanner.Text ()
    text  = string ( text )

    var number_list [] int
    len_list := 0
    
    f := func ( c rune ) bool { return ( ( !unicode.IsNumber ( c ) ) && 
                                         ( !unicode.IsPunct  ( c ) ) )} // https://pkg.go.dev/strings#FieldsFunc
    
    for _ , number := range strings.FieldsFunc ( text , f ) {

       char_int , err := strconv.Atoi ( number )
       
       if err != nil { continue }
    
       number_list = append ( number_list , char_int )
                
       len_list = len ( number_list )
    }
    
    division_four , i , rest := 0 , 1 , 0
    
    if      len_list >= 4 {
       if ( len_list  % 4 ) > 0 { rest = len_list % 4 }
       
       division_four = len_list / 4
    }

    fmt.Println ( ""                                                   )
    fmt.Println ( "Working on length" , len_list                       )
    fmt.Println ( "It was divided in" , division_four , "part by four" )
    fmt.Println ( "With a rest of"    , rest                           )
    
    var wg sync.WaitGroup
    
    fmt.Println ( "" ) ; wg.Add ( 4 )

    for i1 , i2 := 0 , 0 ; ( ( len_list > 0 ) && ( i2 <= len_list ) ) ; i , i1 = i+1 , i2 {
        
        if    ( ( i2 + division_four ) <= len_list ) { i2 += division_four
        } else                                       { i2  = len_list     }

        fmt.Println ( "original array of group" , i , "is:" , number_list [ i1 : i2 ] )

        if    i == 4 { go BubbleSort ( number_list [ i1 :    ] , &wg ) ; i2 = len_list
        } else       { go BubbleSort ( number_list [ i1 : i2 ] , &wg ) }

        if i2 >= len_list { break }
    }
    
    wg.Wait ()

    if len_list >= 4 {

        division_fn :=
        func ( multp int ) int { return ( division_four * multp ) }

        sorted_number_list :=
        merge (             /* less cost =  2*(2*n/4) + (2*n/2) = 2n */

            number_list [                   : division_fn ( 1 ) ] ,
            number_list [ division_fn ( 1 ) : division_fn ( 2 ) ] ,
            number_list [ division_fn ( 2 ) : division_fn ( 3 ) ] ,
            number_list [ division_fn ( 3 ) :                   ] )

        fmt.Println ( "" )
        fmt.Println ( "Final sorted array" , sorted_number_list )
    }
}

// tested on : 14 13 12 11 10 9 8 7 6 5 4 3 2 1 0
// tested on : 6 4 8 89 46 78 0 35 7
// tested on : 6 4 8 9 4 3 8 90 0 5 2 2 2 8 -99 -433
