package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var intSlice = make([]int, 0, 100)
    for i := 0; i < len(intSlice); i++ {
        intSlice = append(intSlice, rand.Intn(100))
    }
    fmt.Println(intSlice)
    for _, v := range intSlice {
        switch {
        case v%2 == 0 && v%3 == 0:
            fmt.Println("Six!")
        case v%2 == 0:
            fmt.Println("Two!")
        case v%3 == 0:
            fmt.Println("Three!")
        default:
            fmt.Println("Never mind!")
        }
    }
}
