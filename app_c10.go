package main

import (
	"example/hello/chapter_10/anony"
	"example/hello/chapter_10/closures"
	"example/hello/chapter_10/math"
	"fmt"
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
}
