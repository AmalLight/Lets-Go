package main

import ( "fmt" ; "encoding/json" )

type Person struct {

    name string
    address string
}

func main () {

    var name , address string = "" , ""

    fmt.Printf ( "Please Enter Name: " ) // FOR debug
    fmt.Scan ( &name )

    fmt.Printf ( "Please Enter Address: " ) // FOR debug
    fmt.Scan ( &address )
    
    var mapMarshal =
    map [ string ] string { "name" : name , "address" : address }
    
    /* // FOR debug

    for key , value := range mapMarshal {
    
        fmt.Printf ( "\n Map Found Key: %s with value: %s \n" , key , value )

    } */
    
    /* // FOR debug

    var pgMarshal =
    Person { name: mapMarshal [ "name" ], address: mapMarshal [ "address" ] }
    
    var pgUnMarshal Person
    
    fmt.Printf ( "\n Testing Person Struct: name = %s with value = %s \n" , pgMarshal.name , pgMarshal.address )

    json_byte , err := json.Marshal ( pgMarshal )

    */
    
    json_byte , err := json.Marshal ( mapMarshal )
    
    var mapUnMarshal map [ string ] string
    
    if err != nil {

        fmt.Printf ( "\n Error happened with json.Marshal \n" )

    } else {
    
        err := json.Unmarshal ( json_byte , &mapUnMarshal )
        
        if err != nil {
        
            fmt.Printf ( "\n Error happened with json.Unmarshal \n" )

        } else { 
        
            // fmt.Printf ( "\n json.Unmarshal done correctly \n" ) // FOR debug
            // fmt.Printf ( "\n %s \n " , mapUnMarshal ) // FOR debug
            
            fmt.Printf ( "%s \n " , string ( json_byte ) )
        }
    }
}

// It Takes name and address from user and converts they to one json byte array
