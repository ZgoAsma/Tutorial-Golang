package main

import (
	"errors"
	"example/hello/chapter_6"
	"example/hello/chapter_6/error_package"
	"fmt"
	"strconv"
)

func main() {

	// Atoi converts a stringified number to an integer
	//i, err := strconv.Atoi("42b")
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("couldn't convert:", err)
		// because "42b" isn't a valid integer, we print:
		// couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
		// Note:
		// 'parsing "42b": invalid syntax' is returned by the .Error() method
		return
	}
	// if we get here, then
	// i was converted successfully
	fmt.Println(i)

	fmt.Println("-----------Errors Assignment----------------")

	chapter_6.Test(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	chapter_6.Test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	chapter_6.Test(
		"Thank you.",
		"Enjoy!",
	)
	chapter_6.Test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)

	fmt.Println("-----------Errors Formatting----------------")

	const name = "Kim"
	const age = 22
	s := fmt.Sprintf("%v is %v years old.", name, age)
	// s = "Kim is 22 years old."
	println(s)

	fmt.Printf("I am %f years old", 10.523)
	// I am 10.523000 years old

	// The ".2" rounds the number to 2 decimal places
	fmt.Printf("I am %.2f years old", 10.523)
	// I am 10.52 years old

	chapter_6.TestFormatting(1.4, "+1 (435) 555 0923")
	chapter_6.TestFormatting(2.1, "+2 (702) 555 3452")
	chapter_6.TestFormatting(32.1, "+1 (801) 555 7456")
	chapter_6.TestFormatting(14.4, "+1 (234) 555 6545")

	chapter_6.TestDiv(10, 0)
	chapter_6.TestDiv(10, 2)
	chapter_6.TestDiv(15, 30)
	chapter_6.TestDiv(6, 3)

	var err1 error = errors.New("something went wrong")
	println(err1.Error())

	error_package.Test(10, 0)
	error_package.Test(10, 2)
	error_package.Test(15, 30)
	error_package.Test(6, 3)

}
