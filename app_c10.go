package main

import (
	"example/hello/chapter_10/defering"
	"example/hello/chapter_10/math"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(math.Aggregate(2, 3, 4, math.Add))
	// prints 9
	fmt.Println(math.Aggregate(2, 3, 4, math.Mul))
	// prints 24
	squareFunc := math.SelfMath(math.Multiply)
	doubleFunc := math.SelfMath(math.Add)

	fmt.Println(squareFunc(5))
	// prints 25

	fmt.Println(doubleFunc(5))
	// prints 10
	users := map[string]defering.User{
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
