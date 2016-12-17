package main

import "fmt"

func main(){
    for i := 0; i <= 10; i++{
        fmt.Println(i)
    }

    myName := "Manjil"
    for _, char := range myName{
        fmt.Print(string(char) + " ")
    }
    fmt.Println()

    j := 0

    for j < 5{
        fmt.Println("golang!")
        j++
        if j == 3{
            break
        }
    }
}
