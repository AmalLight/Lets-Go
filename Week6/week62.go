package main

/*

Let us assume the following formula for
displacement space_final as a function of:
  - time t
  - acceleration a
  - initial velocity v_0
  - and initial splace s_0.

space_final = 1/2 * a * t^2 + v_0 * t + s_0

Write a program which first prompts
to enter values for: ( acceleration, initial velocity, and initial splace ).

Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64 arguments ( acceleration a, initial velocity vo, and initial splace 0 ).

GenDisplaceFn() should return a function which computes new splace as a function of time, assuming the given values for { acceleration, initial velocity, and initial splace }.

The function returned by GenDisplaceFn() should take one ( float64 argument t ), representing time, and return one float64 argument which is the displacement travelled after time t.

*/

import ( "fmt" )

func GenDisplaceFn ( acceleration , initial_velocity , initial_splace float64 ) func ( float64 ) float64 {

    saved_acceleration , saved_initial_velocity , saved_initial_splace :=
          acceleration ,       initial_velocity ,       initial_splace

    // fmt.Println ( "\n\n Acceleration: " , saved_acceleration     )
    // fmt.Println ( " Initial Velocity: " , saved_initial_velocity )
    // fmt.Println ( " Initial Splace: "   , saved_initial_splace   )

    return func ( passed_time float64 ) float64 {

        to_return := ( ( saved_acceleration / 2 )  *  (  passed_time  * passed_time  ) )
        to_return += ( ( saved_initial_velocity * passed_time ) + saved_initial_splace )

        return to_return
    }
}

func main () {

    var GenDisplace_Vs [] float64
    
    var tmp float64 = 0.0
    
    var GenDisplace_Fn func ( float64 ) float64

    for i := 1 ; i < 4 ; i++ {

         tmp = 0.0

         switch i {

         case 1:
             fmt.Printf ( " Please Enter an Acceleration: " )
         case 2:
             fmt.Printf ( " Please Enter an Initial Velocity: " )
         case 3:
             fmt.Printf ( " Please Enter an Initial splace: " )
         }

         fmt.Scan ( &tmp )
         
         GenDisplace_Vs = append ( GenDisplace_Vs , tmp )
         
         if i == 3 {
         
             GenDisplace_Fn = GenDisplaceFn ( GenDisplace_Vs [ 0 ] ,
                                              GenDisplace_Vs [ 1 ] , GenDisplace_Vs [ 2 ] )
         }
    }

    tmp = 0.0
    fmt.Printf ( "\n Please Enter Your Passed Time: " )
    fmt.Scan   ( &tmp                                 )
    
    fmt.Println ( "\n Show Final Splace value: " , GenDisplace_Fn ( tmp ) )
    
    // fn := GenDisplaceFn ( 10 , 2 , 1 )
    // fmt.Println ( fn ( 3 ) )
    // fmt.Println ( fn ( 5 ) )
}

// 
