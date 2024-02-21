package main

import "fmt"

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    fmt.Println(makePerson("Val", "Valson", 33))
    fmt.Println(makePersonPointer("Point", "Pointerson", 33))
}

func makePerson(firstName string, lastName string, age int) Person {
    return Person{firstName, lastName, age}
}

func makePersonPointer(firstName string, lastName string, age int) *Person {
    person := Person{firstName, lastName, age}
    return &person
}
