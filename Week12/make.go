
  buffer := [ variable not constant ] int // not possible
  buffer := make ( [] int , var..not constant ) // possible
  buffer := make ( [] int ) // not possible --> [] int is slice
    
  buffer := make ( int , constant ) // not possible
  buffer := make ( int ) // not possible
    
  buffer := make ( chan int , constant ) // possible --> limited loop read/write to constant
  buffer := make ( chan int ) // possible --> infinite loop read/write
  
  // chan int +- == [] chan
