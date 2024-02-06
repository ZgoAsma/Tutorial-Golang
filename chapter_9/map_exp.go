package chapter_9

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Number of names and phone numbers doesnt match!")
	}
	userMap := make(map[string]user)
	for i, name := range names {
		u := user{name: name, phoneNumber: phoneNumbers[i]}
		userMap[name] = u
	}
	return userMap, nil
}

// don't touch below this line

type user struct {
	name        string
	phoneNumber int
}

func Test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}
