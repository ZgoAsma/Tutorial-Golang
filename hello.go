package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Prints 0 through 9

	for i := 0; ; i++ {
		// do something forever
		fmt.Println("In infinite loop,", i)
		//break
		if i == 10 {
			break
		}
	}

	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to ", plantHeight, "inches")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	// 1
	// 3
	// 5
	// 7
	// 9

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	// 0
	// 1
	// 2
	// 3
	// 4

}
