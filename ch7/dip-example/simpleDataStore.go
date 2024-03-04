package main

type SimpleDataStore struct {
    userData map[string]string
}

func NewSimpleDataStore() SimpleDataStore {
    return SimpleDataStore{
        userData: map[string]string{
            "1": "Fred",
            "2": "Mary",
            "3": "Pat",
        },
    }
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
    name, ok := sds.userData[userId]
    return name, ok
}
