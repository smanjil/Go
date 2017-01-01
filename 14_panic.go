
package main

import (
    "io/ioutil"
    "fmt"
)

func main() {
    // read the whole line at once
    b, err := ioutil.ReadFile("names.txt")
    if err != nil {
        panic(err)
    }

    for _, e := range b {
        fmt.Print(string(e))
    }
}
