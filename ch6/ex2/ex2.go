package main

import "fmt"

func main() {
    slc := make([]string, 10)
    fmt.Println(slc)
    UpdateSlice(slc, "Update")
    fmt.Println(slc)
    GrowSlice(slc, "Grow")
    fmt.Println(slc)
}

func UpdateSlice(slc []string, str string) {
    slc[cap(slc)-1] = str
    fmt.Println(slc)
}

func GrowSlice(slc []string, str string) {
    slc = append(slc, str)
    fmt.Println(slc)
}
