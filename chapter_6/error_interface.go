package chapter_6

import "fmt"

type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintln("can not divide DIVIDEND by zero\n")
}

// don't edit below this line

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// We convert the `divideError` struct to an `error` type by returning it
		// as an error. As an error type, when it's printed its default value
		// will be the result of the Error() method
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func TestDiv(dividend, divisor float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}
