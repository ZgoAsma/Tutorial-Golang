package main

import (
	"example/hello/chapters_1_4"
	"fmt"
	"time"
)

func main() {
	//functions call
	chapters_1_4.Test("Lane,", " happy birthday!")
	chapters_1_4.Test("Elon,", " hope that Tesla thing works out")
	chapters_1_4.Test("Go", " is fantastic")

	fmt.Println("-------------------")

	//functions (+syntax sugar) call
	chapters_1_4.AddToDatabase(20, 9, "AsmaZ", 100)

	fmt.Println("-------------------")

	//call sending function (passing variables by values)
	chapters_1_4.Sending()

	fmt.Println("-------------------")

	//ignoring return values
	chapters_1_4.Names()

	fmt.Println("-------------------")

	//naked returns (explicit vs implicit)
	fmt.Println(chapters_1_4.GetCoords())
	fmt.Println(chapters_1_4.GetCoordsAlt())

	fmt.Println("-------------------")

	//assignement funcs
	chapters_1_4.UntilAge(4)
	chapters_1_4.UntilAge(10)
	chapters_1_4.UntilAge(18)
	chapters_1_4.UntilAge(22)
	chapters_1_4.UntilAge(34)

	fmt.Println("-------------------")

	//explicit return
	fmt.Println(chapters_1_4.GetCoordsAlt2())
	chapters_1_4.UntilAgeAlt(4)
	chapters_1_4.UntilAgeAlt(22)
	chapters_1_4.UntilAgeAltImplicit(4)
	chapters_1_4.UntilAgeAltImplicit(22)

	fmt.Println("-------------------")

	//guard clause
	fmt.Println(chapters_1_4.Divide(89, 6))
	fmt.Println(chapters_1_4.Divide(89, 0))

	fmt.Println("-------------------")

	//structs
	chapters_1_4.TestStruct(chapters_1_4.MessageToSendSimple{
		PhoneNumber: 148255510981,
		Message:     "Thanks for signing up",
	})
	chapters_1_4.TestStruct(chapters_1_4.MessageToSendSimple{
		PhoneNumber: 148255510982,
		Message:     "Love to have you aboard!",
	})
	chapters_1_4.TestStruct(chapters_1_4.MessageToSendSimple{
		PhoneNumber: 148255510983,
		Message:     "We're so excited to have you",
	})
	myCar := chapters_1_4.Car{}
	myCar.FrontWheel.Radius = 5
	println(myCar.BackWheel.Material)
	println(myCar.Height)

	//nested structs
	chapters_1_4.TestNestedStruct(chapters_1_4.MessageToSend{
		Message: "you have an appointment tommorow",
		Sender: chapters_1_4.User{
			Name:   "Brenda Halafax",
			Number: 16545550987,
		},
		Recipient: chapters_1_4.User{
			Name:   "Sally Sue",
			Number: 19035558973,
		},
	})
	chapters_1_4.TestNestedStruct(chapters_1_4.MessageToSend{
		Message: "you have an event tomorrow",
		Sender: chapters_1_4.User{
			Number: 16545550987,
		},
		Recipient: chapters_1_4.User{
			Name:   "Suzie Sall",
			Number: 19035558973,
		},
	})
	chapters_1_4.TestNestedStruct(chapters_1_4.MessageToSend{
		Message: "you have an party tomorrow",
		Sender: chapters_1_4.User{
			Name: "Njorn Halafax",
		},
		Recipient: chapters_1_4.User{
			Name:   "Becky Sue",
			Number: 19035558973,
		},
	})
	chapters_1_4.TestNestedStruct(chapters_1_4.MessageToSend{
		Message: "you have a birthday tomorrow",
		Sender: chapters_1_4.User{
			Name:   "Eli Halafax",
			Number: 16545550987,
		},
		Recipient: chapters_1_4.User{
			Number: 19035558973,
		},
	})
	chapters_1_4.TestNestedStruct(chapters_1_4.MessageToSend{
		Message: "you have a birthday tomorrow",
		Sender: chapters_1_4.User{
			Name:   "Jason Bjorn",
			Number: 16545550987,
		},
		Recipient: chapters_1_4.User{
			Name: "Jim Bond",
		},
	})

	// ananymous struct
	var myCar2 = struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}
	println(myCar2.Model)

	// embedded truck
	var lanesTruck = chapters_1_4.Truck{
		BedSize: 10,
		Car: chapters_1_4.Car{
			Make:  "toyota",
			Model: "camry",
		},
	}

	fmt.Println(lanesTruck.BedSize)

	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.Make
	fmt.Println(lanesTruck.Make)
	fmt.Println(lanesTruck.Model)

	chapters_1_4.TestEmb(chapters_1_4.Sender{
		RateLimit: 10000,
		User_emb: chapters_1_4.User_emb{
			Name:   "Deborah",
			Number: 18055558790,
		},
	})
	chapters_1_4.TestEmb(chapters_1_4.Sender{
		RateLimit: 5000,
		User_emb: chapters_1_4.User_emb{
			Name:   "Sarah",
			Number: 19055558790,
		},
	})
	chapters_1_4.TestEmb(chapters_1_4.Sender{
		RateLimit: 1000,
		User_emb: chapters_1_4.User_emb{
			Name:   "Sally",
			Number: 19055558790,
		},
	})

	// methods associated to structs

	var r = chapters_1_4.Rect{
		Width:  5,
		Height: 10,
	}

	fmt.Println(r.Area())
	// prints 50
	chapters_1_4.TestAuth(chapters_1_4.AuthenticationInfo{
		Username: "Google",
		Password: "12345",
	})
	chapters_1_4.TestAuth(chapters_1_4.AuthenticationInfo{
		Username: "Bing",
		Password: "98765",
	})
	chapters_1_4.TestAuth(chapters_1_4.AuthenticationInfo{
		Username: "DDG",
		Password: "76921",
	})

	// interfaces
	chapters_1_4.TestRec(chapters_1_4.SendingReport{
		ReportName:    "First Report",
		NumberOfSends: 10,
	})
	chapters_1_4.TestRec(chapters_1_4.BirthdayMessage{
		RecipientName: "John Doe",
		BirthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	chapters_1_4.TestRec(chapters_1_4.SendingReport{
		ReportName:    "First Report",
		NumberOfSends: 10,
	})
	chapters_1_4.TestRec(chapters_1_4.BirthdayMessage{
		RecipientName: "Bill Deer",
		BirthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})

}
