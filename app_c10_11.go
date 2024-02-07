package main

import (
	"example/hello/chapter_10/anony"
	"example/hello/chapter_10/closures"
	"example/hello/chapter_10/math"
	"fmt"
	"strings"
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

	harryPotterAggregator := closures.Concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive

	nums := []int{1, 2, 3, 4, 5}

	// Here we define an anonymous function that doubles an int
	// and pass it to DoMath
	allNumsDoubled := anony.DoMath(func(x int) int {
		return x + x
	}, nums)

	fmt.Println(allNumsDoubled)
	// prints:
	// [2 4 6 8 10]

	x := 42
	// "x" is the name of a location in memory. That location is storing the integer value of 42

	println(x)
	var p *int

	println(p)
	myString := "hello"
	myStringPtr := &myString

	println(myStringPtr)

	fmt.Println(*myStringPtr) // read myString through the pointer
	*myStringPtr = "world"    // set myString through the pointer

	fmt.Println(*myStringPtr)
	var x1 int = 50
	var y *int = &x1
	*y = 100
	println(*y)
	println(x1)

	messages := []string{
		"well shoot, this is awful",
		"",
		"dang robots",
		"dang them to heck",
		"",
	}

	messages2 := []string{
		"well shoot",
		"",
		"Allan is going straight to heck",
		"dang... that's a tough break",
		"",
	}

	testProf(messages)
	testProf(messages2)
}

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}

// don't touch below this line

func testProf(messages []string) {
	for _, message := range messages {
		if message == "" {
			removeProfanity(nil)
			fmt.Println("nil message detected")
		} else {
			removeProfanity(&message)
			fmt.Println(message)
		}
	}
}
