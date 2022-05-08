package main

import ( "fmt" ; "bufio" ; "os" ; "io/ioutil" )

type Person struct {

    fname string
    lname string
}

func main () {

    fmt.Printf ( "Please Enter Name Text File: " ) // FOR debug

    scanner := bufio.NewScanner ( os.Stdin )
    scanner.Scan ()

    file_name := scanner.Text ()

    dat , err := ioutil.ReadFile ( file_name )

    var Array_Person [] Person

    if err != nil {

        fmt.Printf ( "Error happened with ioutil.ReadFile \n" )

    } else {

        name_last , name_last_bool := "" , true
        name_line , last_line      := "" ,   ""

        for _ , line := range dat {

           string_line := string ( line )

           if ( ( string_line != " " ) && ( string_line != "\n" ) ) {

               name_last += string_line

           } else {

               if name_last_bool { name_line = name_last } else
                                 { last_line = name_last }

               if name_last_bool == false {

                   // fmt.Printf ( "Find name: %s and last_name: %s \n" , name_line , last_line ) // FOR debug

                   Person_line := Person { fname : name_line , lname : last_line }

                   Array_Person = append ( Array_Person , Person_line )
               }

               name_last , name_last_bool = "" , ( ! name_last_bool )
           }
        }
    }

    for _ , P := range Array_Person {

        fmt.Printf ( "Print from Array, Name: %s with Last name: %s \n" , P.fname , P.lname )
    }
}

// /home/kaumi/Git/Lets-Go/Course/Week4/test_read.go.txt
// All done as requested by teacher
