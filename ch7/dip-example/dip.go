package main

import (
    "fmt"
    "net/http"
)

type DataStore interface {
    UserNameForId(userId string) (string, bool)
}

type Logger interface {
    Log(message string)
}

type Logic interface {
    SayHello(userId string) (string, error)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
    lg(message)
}

func main() {
    l := LoggerAdapter(LogOutput)
    ds := NewSimpleDataStore()
    logic := NewSimpleLogic(l, ds)
    c := NewController(l, logic)
    http.HandleFunc("/hello", c.SayHello)
    http.ListenAndServe(":8080", nil)
}

func LogOutput(message string) {
    fmt.Println(message)
}