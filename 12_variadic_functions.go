
package main

import "fmt"

func main() {
    add(6, 84, 786, 12, 8)
}

func add(nums ...int) {
    var total int

    for _, n := range nums {
        total += n
    }

    fmt.Println(total)
}
