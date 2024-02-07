package main

import (
	"example/hello/chapter_10/defering"
	"fmt"
	"sort"
)

func main() {
	var users = map[string]defering.MyUser{
		"john": {
			Name:   "john",
			Number: 18965554631,
			Admin:  true,
		},
		"elon": {
			Name:   "elon",
			Number: 19875556452,
			Admin:  true,
		},
		"breanna": {
			Name:   "breanna",
			Number: 98575554231,
			Admin:  false,
		},
		"kade": {
			Name:   "kade",
			Number: 10765557221,
			Admin:  false,
		},
	}
	fmt.Println("Initial users:")
	usersSorted := []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	defering.Test(users, "john")
	defering.Test(users, "santa")
	defering.Test(users, "kade")

	fmt.Println("Final users:")
	usersSorted = []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")
}
