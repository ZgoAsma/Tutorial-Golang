package mutations

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]User, name string) (deleted bool, err error) {
	// get user
	u, ok := users[name]
	if !ok {
		return false, errors.New("User not found!")
	}
	if u.ScheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

// don't touch below this line

type User struct {
	Name                 string
	Number               int
	ScheduledForDeletion bool
}

func Test(users map[string]User, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted:", name)
		return
	}
	fmt.Println("Did not delete:", name)
}
