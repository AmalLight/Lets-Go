package main

import ( "fmt" ; "strings" ; "bufio" ; "os" )

type Animal interface {

    Eat   ()
    Move  ()
    Speak ()
}

var M map [ string ] Animal = map [ string ] Animal {}

type Cow   struct { info infoAnimal }
type Bird  struct { info infoAnimal }
type Snake struct { info infoAnimal }

type infoAnimal struct {

    food         string
    locomotion   string
    spoken_sound string
}

// func (A Animal) NewAnimal (name string) *Animal {}
// invalid receiver type Animal (Animal is an interface type)

func NewAnimal ( name string , either string ) {

    switch either {
    
      case "cow"   :

          M [ name ] = Cow   { infoAnimal { "grass" , "walk"    , "moo" } }
      
      case "bird"  :
      
          M [ name ] = Bird  { infoAnimal { "worms" , "fly"     , "peep" } }
      
      case "snake" :
      
          M [ name ] = Snake { infoAnimal { "mice"  , "slither" , "hsss" } }
    }
    
    if either == "cow"   ||
       either == "bird"  ||
       either == "snake"  { fmt.Println ( "Created it!" ) }
}

func Query ( name string , info string ) {

    var A Animal = nil
    A = M [ name ]

    if A != nil {

        switch info {
    
            case "eat" :
                A.Eat ()
        
            case "move" :
                A.Move ()
        
            case "speak" :
                A.Speak ()
        }
    }
}

func (C   Cow) Eat   () { fmt.Println ( C.info.food         ) }
func (B  Bird) Eat   () { fmt.Println ( B.info.food         ) }
func (S Snake) Eat   () { fmt.Println ( S.info.food         ) }

func (C   Cow) Move  () { fmt.Println ( C.info.locomotion   ) }
func (B  Bird) Move  () { fmt.Println ( B.info.locomotion   ) }
func (S Snake) Move  () { fmt.Println ( S.info.locomotion   ) }

func (C   Cow) Speak () { fmt.Println ( C.info.spoken_sound ) }
func (B  Bird) Speak () { fmt.Println ( B.info.spoken_sound ) }
func (S Snake) Speak () { fmt.Println ( S.info.spoken_sound ) }

func main () {

   scanner := bufio.NewScanner ( os.Stdin )

   for {
       userInput := ""

       fmt.Print ( "\n" , ">" , " " )
       
       scanner.Scan ()
       userInput = scanner.Text ()
       
       indexSpace := strings.Index ( userInput , " " )
       
       if indexSpace < 0 { continue }
       
       command := userInput [ : indexSpace ]
       
       other := userInput [ indexSpace + 1 : ]
       indexSpace = strings.Index ( other , " " )

       if indexSpace < 0 { continue }
       
       name        := other [ : indexSpace       ]
       either_info := other [   indexSpace + 1 : ]
       
       if command == "newanimal" {

           NewAnimal ( name , either_info ) 

       } else if command == "query" {
       
           Query ( name , either_info )
       }
   }
}
