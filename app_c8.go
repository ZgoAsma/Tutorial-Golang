package main

import (
	"example/hello/chapter_8/arrays"
	make_slice "example/hello/chapter_8/make"
	"example/hello/chapter_8/slices"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func main() {
	var myInts [10]int
	primes := [6]int{2, 3, 5, 7, 11, 13}

	println(myInts[2])
	println(primes[1])

	//app
	arrays.TestSend("Bob", 0)
	arrays.TestSend("Alice", 1)
	arrays.TestSend("Mangalam", 2)
	arrays.TestSend("Ozgur", 3)

	mySlice := primes[1:4]
	// mySlice = {3, 5, 7}

	println(len(primes))
	println(len(mySlice))

	//app
	slices.Test("Ozgur", 3, planFree)
	slices.Test("Jeff", 3, planPro)
	slices.Test("Sally", 2, planPro)
	slices.Test("Sally", 3, "no plan")

	// func make([]T, len, cap) []T
	mySlice = make([]int, 5, 10)

	println(len(mySlice))
	println(cap(mySlice))
	// the capacity argument is usually omitted and defaults to the length
	mySlice = make([]int, 6)
	println(len(mySlice))
	println(cap(mySlice))

	mySlice1 := []string{"I", "love", "go"}
	fmt.Println(len(mySlice1)) // 3
	fmt.Println(cap(mySlice1))

	// make app
	make_slice.Test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	make_slice.Test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	make_slice.Test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
	final := make_slice.Concat("Hello ", "there ", "friend!")
	fmt.Println(final)
	// Output: Hello there friend!
	names := []string{"bob", "sue", "alice"}
	make_slice.PrintStrings(names...)

	make_slice.TestNum(1.0, 2.0, 3.0)
	make_slice.TestNum(1.0, 2.0, 3.0, 4.0, 5.0)
	make_slice.TestNum(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	make_slice.TestNum(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)

}
