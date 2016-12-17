
package main

import "fmt"

func main(){
    // var myAge int = 21
    // myVariable := 34

    myAge, myVariable := 21, 34

    fmt.Println(myAge)
    fmt.Println(myVariable)

    fmt.Println("--------MODIFIED---------")

    // myAge = 37
    // myVariable = 765

    myAge, myVariable = 37, 765

    fmt.Println(myAge)
    fmt.Println(myVariable)
}
