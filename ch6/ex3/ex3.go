package main

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    collection := make([]Person, 10_000_000)
    for i, _ := range collection {
        collection[i] = Person{"Bob", "Bobson", 30}
    }
}
