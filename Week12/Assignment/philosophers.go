package main

import ( "fmt" ; "sync" ; "time" ; "strconv" )

const common_size  int = 5
const common_limit int = 3
const common_time  int = 1

// --------------------------------------------------------------------------------

// object with Lock and Unlock // auto this ==> like a class extension

type Status struct { sync.Mutex ; database [common_limit][common_size] int }
type ChopS  struct { sync.Mutex                                            }
type Philo  struct {

    wg *sync.WaitGroup

    leftCS  *ChopS
    rightCS *ChopS
    
    name  string
    numb  int
    limit int
    time  time.Duration
}

// --------------------------------------------------------------------------------

func (P Philo) eat ( c chan int ) {

    for i := 0 ; i < P.limit ; i++ {
    
        <- c // P numb wait for Manage

        for {
            
            c <- 1 // I want to eat: i
            r := <- c // response
            
            if r == 1 { fmt.Println ( "( " + strconv.Itoa ( P.numb ) + " )" , P.name , "current: yes! for" , i + 1 )
            } else    { fmt.Println ( "( " + strconv.Itoa ( P.numb ) + " )" , P.name , "current: no! for"  , i + 1 ) }
            
            if r == 1 { break
            } else    { time.Sleep  ( ( P.time * 2 ) * time.Second ) }
        }

        P.leftCS. Lock ()
        P.rightCS.Lock ()
    
        fmt.Println ( "\n( " + strconv.Itoa ( P.numb ) + " )" , P.name ,
        "is starting to eat for STEP" , i + 1 , "\n" )

        time.Sleep ( P.time * time.Second )
    
        P.leftCS. Unlock()
        P.rightCS.Unlock()
    
        fmt.Println ( "\n( " + strconv.Itoa ( P.numb ) + " )" , P.name  ,
        "is finishing eating for STEP" , i + 1 , "\n" )
        
        for {
            
            c <- 2 // I finished to eat: i
            r := <- c // response
            
            if r == 2 { fmt.Println ( "( " + strconv.Itoa ( P.numb ) + " )" , P.name , "next:" , i + 2 , "from" , i + 1 ) ; break
            } else    { time.Sleep  (                                       ( P.time * 2 ) * time.Second )                ; }
        }
    }

    P.wg.Done ()
}

func CountS ( index int , S *Status ) (int,int) {

    count :=  0
    who   := -1
    
    for j , el := range S.database [ index ] { if ( el == 1 ) { count += 1 ; who = j } }
    
    return who , count
}

func WriteS ( index int , S *Status , numb int ) {

    S.Lock ()
    S.database [index][numb] = 2
    S.Unlock ()
}

func Manage ( index int , S *Status , numb int , c chan int , wg *sync.WaitGroup ) {

    c <- numb // P numb can start to ask for i
    
    for {
    
        r := <- c // request

        if r == 1 {

            S.Lock ()

            who , count := CountS ( index , S )
            
            Can_Eat := false
            
            if ( count == 0 ) || ( ( who > -1 ) && ( count == 1 ) ) {
                
                if ( count == 0 ) || ( ( count == 1 ) && (
                    
                    // 0 = 2 3 , 1 = 3 4 , 2 = 4 0 , 3 = 0 1 , 4 = 1 2
                    
                   ( numb != ( ( who + common_size - 1 ) % common_size )) &&
                   ( numb != ( ( who +               1 ) % common_size )) )) {

                    S.database [index][numb] = 1 ; Can_Eat = true ; 
                }
            }
            
            S.Unlock ()
            
            if Can_Eat == false { c <- 0
            } else              { c <- 1 ; }


        } else if r == 2 { WriteS ( index , S , numb ) ; c <- 2 ; break ; }
    }

    wg.Done ()
}

func main () { // HOST

    var CSticks [] ChopS
    var PhiloS  [] Philo
    
    for i := 0 ; i < common_size ; i++ { CSticks = append ( CSticks , ChopS {} ) }
    for i := 0 ; i < common_size ; i++ { PhiloS  = append ( PhiloS  , Philo {} ) }
    
    buffer := make ( [] chan int , common_size ) // build a Matrix! [5][1] chan int
    for i := 0 ; i < common_size ; i++ { buffer [ i ] = make ( chan int ) }
    
    var wg sync.WaitGroup  ;
    wg.Add ( common_size ) ;
    
    names := [] string { "Pippo", "Pluto", "Paperino", "Topolino", "Pietro" }
    
    for i , P :=
    range PhiloS {
    
        P.name = names [ i ] ; P.numb = i+1 ; P.limit = common_limit ; P.time = time.Duration ( common_time )
        
        P.leftCS  = & CSticks [   i                     ] ;
        P.rightCS = & CSticks [ ( i + 1 ) % common_size ] ;
        
        P.wg = & wg ; go P.eat ( buffer [ i ] )
    }
    
    var S Status

    for i := 0 ; i < common_limit ; i++ {
    
        var wg_tmp sync.WaitGroup  ;
        wg_tmp.Add ( common_size ) ;

        for j , _ :=
        range PhiloS { go Manage ( i , & S , j , buffer [ j ] , & wg_tmp ) ; }
        
        wg_tmp.Wait ()
    }
    
    wg.Wait ()
}
