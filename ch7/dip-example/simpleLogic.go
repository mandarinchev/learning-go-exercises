package main

import "errors"

type SimpleLogic struct {
    l Logger
    ds DataStore
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
    return SimpleLogic{
        l:  l,
        ds: ds,
    }
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
    sl.l.Log("in SayHello for " + userId)
    name, ok := sl.ds.UserNameForId(userId)
    if !ok {
        return "", errors.New("unknown user")
    }
    return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userId string) (string, error) {
    sl.l.Log("in SayGoodbye for " + userId)
    name, ok := sl.ds.UserNameForId(userId)
    if !ok {
        return "", errors.New("unknown user")
    }
    return "Goodbye, " + name, nil
}
