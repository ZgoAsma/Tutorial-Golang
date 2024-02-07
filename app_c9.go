package main

import "fmt"

func main() {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	fmt.Println(timeZone)

	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	person := "Asma"
	if attended[person] { // will be false if person is not in the map
		fmt.Println(person, "was at the meeting")
	}

	fmt.Println(aggregate(2, 3, 4, add))
	// prints 9
	fmt.Println(aggregate(2, 3, 4, mul))
	// prints 24

}

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}
