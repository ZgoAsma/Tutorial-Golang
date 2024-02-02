package main

import (
	"example/hello/chapter_5"
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

}
