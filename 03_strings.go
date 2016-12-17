

package main

import "fmt"

func main(){
    phrase1 := "The quick brown fox"
    phrase2 := " jumps over the lazy dog."

    fmt.Println(phrase1 + phrase2)

    aSliceOfText := phrase1[1:3]
    fmt.Println(aSliceOfText)

    noun1 := phrase1[16:]
    fmt.Println(noun1)

    noun2 := phrase2[21:]
    fmt.Println(noun2)

    character := phrase1[0]
    fmt.Println(character)
    fmt.Println(string(character))
}
