package main

import (
	"example/hello/chapter_5"
	"example/hello/chapter_5/type_switches"
	"fmt"
)

func main() {
	chapter_5.Test(chapter_5.FullTime{
		Name:   "Jack",
		Salary: 50000,
	})
	chapter_5.Test(chapter_5.Contractor{
		Name:         "Bob",
		HourlyPay:    100,
		HoursPerYear: 73,
	})
	chapter_5.Test(chapter_5.Contractor{
		Name:         "Jill",
		HourlyPay:    872,
		HoursPerYear: 982,
	})

	/*
		s := chapter_5.Circle{}
		c, ok := s.(chapter_5.Circle)
		if !ok {
			// log an error if s isn't a circle
			log.Fatal("s is not a circle")
		}

		radius := c.Radius
		println(radius)

	*/

	// multiple interfaces
	chapter_5.TestPhones(chapter_5.EmailAlt{
		IsSubscribed: true,
		Body:         "hello there",
		ToAddress:    "john@does.com",
	})
	chapter_5.TestPhones(chapter_5.EmailAlt{
		IsSubscribed: false,
		Body:         "This meeting could have been an email",
		ToAddress:    "jane@doe.com",
	})
	chapter_5.TestPhones(chapter_5.EmailAlt{
		IsSubscribed: false,
		Body:         "This meeting could have been an email",
		ToAddress:    "elon@doe.com",
	})
	chapter_5.TestPhones(chapter_5.SMS{
		IsSubscribed:  false,
		Body:          "This meeting could have been an email",
		ToPhoneNumber: "+155555509832",
	})
	chapter_5.TestPhones(chapter_5.SMS{
		IsSubscribed:  false,
		Body:          "This meeting could have been an email",
		ToPhoneNumber: "+155555504444",
	})
	chapter_5.TestPhones(chapter_5.Invalid{})

	// type switches
	type_switches.TestSwitch(type_switches.Email{
		IsSubscribed: true,
		Body:         "hello there",
		ToAddress:    "john@does.com",
	})
	type_switches.TestSwitch(type_switches.Email{
		IsSubscribed: false,
		Body:         "This meeting could have been an email",
		ToAddress:    "jane@doe.com",
	})
	type_switches.TestSwitch(type_switches.Email{
		IsSubscribed: false,
		Body:         "Wanna catch up later?",
		ToAddress:    "elon@doe.com",
	})
	type_switches.TestSwitch(type_switches.Sms{
		IsSubscribed:  false,
		Body:          "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars",
		ToPhoneNumber: "+155555509832",
	})
	type_switches.TestSwitch(type_switches.Sms{
		IsSubscribed:  false,
		Body:          "I don't need this",
		ToPhoneNumber: "+155555504444",
	})
	type_switches.TestSwitch(type_switches.Invalid{})

	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"

}

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}
