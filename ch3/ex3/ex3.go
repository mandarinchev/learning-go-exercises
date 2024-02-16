package main

import "fmt"

type Employee struct {
    firstName string
    lastName  string
    id        int
}

func main() {
    employee1 := Employee{"Dick", "Tracy", 0}
    employee2 := Employee{
        firstName: "Joe",
        lastName:  "Jobber",
        id:        1,
    }
    var employee3 Employee
    employee3.firstName = "Mr."
    employee3.lastName = "Buster"
    employee3.id = 2
    fmt.Println(employee1, employee2, employee3)
}
