package main

import ( "fmt" ; "strings" ; "bufio" ; "os" )

type Animal struct {

    food       string
    locomotion string
    noiseame   string
}

func NewAnimal (name string) *Animal {

    var tmpA * Animal
    tmpA = & Animal { "" , "" , "" }

    eat , move , speak := "" , "" , ""

    switch name {
   
      case "cow"   :
          eat = "grass"
          move = "walk"
          speak = "moo"
      
      case "bird"  :
          eat = "worms"
          move = "fly"
          speak = "peep"
      
      case "snake" :
          eat = "mice"
          move = "slither"
          speak = "hsss"
   }
   
   tmpA.food       = eat
   tmpA.locomotion = move
   tmpA.noiseame   = speak
   
   return tmpA
}

func (A Animal) Eat   () { fmt.Println ( A.food       ) }
func (A Animal) Move  () { fmt.Println ( A.locomotion ) }
func (A Animal) Speak () { fmt.Println ( A.noiseame   ) }

func main () {

   var A * Animal = nil

   scanner := bufio.NewScanner ( os.Stdin )

   for {
       userInput := ""

       fmt.Print ( "\n" , ">" , " " )
       
       scanner.Scan ()
       userInput = scanner.Text ()
       
       indexSpace := strings.Index ( userInput , " " )       
       
       name  := userInput [ : indexSpace       ]
       about := userInput [   indexSpace + 1 : ]
       
       A = NewAnimal ( name )
       
       switch about {
       
          case "eat"   : (*A).Eat   ()
          case "move"  : (*A).Move  ()
          case "speak" : (*A).Speak ()
       }
   }
}
