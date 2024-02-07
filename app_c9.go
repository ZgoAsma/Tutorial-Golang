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

}
