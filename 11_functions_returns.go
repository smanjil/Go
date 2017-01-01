
package main

import "fmt"

func main() {
    thestring, thenumber := superComplexFunction()
    fmt.Println(thestring)
    fmt.Println(thenumber)
}

func superComplexFunction() (string, int) {
    e := "some string"
    num := 5

    return e, num
}
