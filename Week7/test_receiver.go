package main

import ( "fmt" )

type MyInt int // cannot define new methods on non-local type int

func (x MyInt) overwrite (y int) {

    z := MyInt ( y )

    if y > 0 { x = z }

    fmt.Println ( x )
}

func (x * MyInt) overwrite2 (y int) {

    z := MyInt ( y )

    if y > 0 { * x = z }

    fmt.Println ( * x )
}

/* DOESN'T EXIST

func (x & MyInt) overwrite3 (y int) {

    // nothing...
}

*/


func main () {

   var x MyInt = 5
   
   x.overwrite ( 9 ) // 9
   x.overwrite ( 0 ) // 5

   var x2 * MyInt
   five := MyInt ( 5 )
   x2 = & five
   
   x2.overwrite2 ( 9 ) // 9
   x2.overwrite2 ( 0 ) // 5 if { x = & z } else ( * x = z ) 9
   
   // 5 maybe because after the call of function the 9 index is dead
}
