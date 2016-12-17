
package main

import "fmt"

func main(){
    // constants always need to be initialized
    const goldenRatio float64 = 1.618034
    fmt.Println(goldenRatio)

    // constant enumeration
    const(
        First = iota
        Second
        Third
    )

    fmt.Println(First, Second, Third)
}
