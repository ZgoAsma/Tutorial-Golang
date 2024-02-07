package defering

import (
	"fmt"
)

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]MyUser, name string) (log string) {
	user, ok := users[name]
	defer delete(users, name)
	if !ok {
		return logNotFound
	}
	if user.Admin {
		return logAdmin
	}
	return logDeleted
}

// don't touch below this line

type MyUser struct {
	Name   string
	Number int
	Admin  bool
}

func Test(users map[string]MyUser, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}
