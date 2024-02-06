package main

import (
	"example/hello/chapter_8/arrays"
	make_slice "example/hello/chapter_8/make"
	"example/hello/chapter_8/matrixes"
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

	fmt.Println("********************Slices App********************")
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

	fmt.Println("********************Make , variadic********************")
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

	anotherSlice := []string{"but", "I", "prefer", "Python", "!"}
	oneThing := "!"
	//anti-pattern, this way we can ran to some bugs, we should reassign the result of the append to mySlice1
	mySlice2 := append(mySlice1, oneThing)
	fmt.Println(mySlice2)
	firstThing := "and"
	secondThing := "Python"
	mySlice2 = append(mySlice1, firstThing, secondThing)
	fmt.Println(mySlice2)
	mySlice1 = append(mySlice1, anotherSlice...)
	fmt.Println(mySlice1)

	make_slice.TestSliceAppend([]make_slice.Cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	make_slice.TestSliceAppend([]make_slice.Cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})

	matrixes.Test(3, 3)
	matrixes.Test(5, 5)
	matrixes.Test(10, 10)
	matrixes.Test(15, 15)

	// bugs from assigning to a new slice using append :

	fmt.Println("********************Example 1********************")
	a := make([]int, 3)
	fmt.Println("len of a:", len(a))
	// len of a: 3
	fmt.Println("cap of a:", cap(a))
	// cap of a: 3
	fmt.Println("appending 4 to b from a")
	// appending 4 to b from a
	b := append(a, 4)
	fmt.Println("b:", b)
	// b: [0 0 0 4]
	fmt.Println("addr of b:", &b[0])
	// addr of b: 0x44a0c0
	fmt.Println("appending 5 to c from a")
	// appending 5 to c from a
	c := append(a, 5)
	fmt.Println("addr of c:", &c[0])
	// addr of c: 0x44a180
	fmt.Println("a:", a)
	// a: [0 0 0]
	fmt.Println("b:", b)
	// b: [0 0 0 4]
	fmt.Println("c:", c)
	// c: [0 0 0 5]

	//example 2
	fmt.Println("********************Example 2********************")
	i := make([]int, 3, 8)
	fmt.Println("len of i:", len(i))
	// len of i: 3
	fmt.Println("cap of i:", cap(i))
	// cap of i: 8
	fmt.Println("appending 4 to j from i")
	// appending 4 to j from i
	j := append(i, 4)
	fmt.Println("j:", j)
	// j: [0 0 0 4]
	fmt.Println("addr of j:", &j[0])
	// addr of j: 0x454000
	fmt.Println("appending 5 to g from i")
	// appending 5 to g from i
	g := append(i, 5)
	fmt.Println("addr of g:", &g[0])
	// addr of g: 0x454000
	fmt.Println("i:", i)
	// i: [0 0 0]
	fmt.Println("j:", j)
	// j: [0 0 0 5]
	fmt.Println("g:", g)
	// g: [0 0 0 5]

	//ranges
	fmt.Println("********************Ranges********************")
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	fmt.Println("create conflict!")
	fmt.Println("Chapter 8 is over!")

}
