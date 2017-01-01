
package main

import "fmt"

func main() {
    defer doSomething()
    doSomethingElse()
}

func doSomething() {
    fmt.Println("doSomething()")
}

func doSomethingElse() {
    fmt.Println("doSomethingElse()")
}
