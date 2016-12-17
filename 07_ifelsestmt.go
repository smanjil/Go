
package main

import "fmt"

func main(){
    condition := 5 > 2

    if condition{
        fmt.Println("5 is indeed greater than 2.")
    }

    // && AND
    if 6 > 3 && 5 > 8{
        fmt.Println("Worked!")
    }else{
        fmt.Println("Didn't work!")
    }

    // || OR
    if 6 > 3 || 5 > 8{
        fmt.Println("Worked!")
    }else{
        fmt.Println("Not worked!")
    }
}
