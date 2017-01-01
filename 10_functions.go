
package main

import "fmt"

func main() {
    doSomething()
    add(2, 4)
}

func doSomething() {
    fmt.Println("Performing the doSomething() function...")
}

func add(a int, b int) {
    fmt.Println(a + b)
}
