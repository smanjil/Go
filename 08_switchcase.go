
package main

import "fmt"

func main(){
    command := "close"

    switch command{
        case "create":
            fmt.Println("Creating....")
        case "open":
            fmt.Println("Opening....")
        case "close":
            fmt.Println("Closing....")
        default:
            fmt.Println("Unrecognized command!")
    }
}
