package main

import "fmt"

func main() {
	//functions call
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")

	fmt.Println("-------------------")

	//functions (+syntax sugar) call
	addToDatabase(20, 9, "AsmaZ", 100)

	fmt.Println("-------------------")

	//call sending function (passing variables by values)
	sending()

	fmt.Println("-------------------")

	//ignoring return values
	names()

	fmt.Println("-------------------")

	//naked returns (explicit vs implicit)
	fmt.Println(getCoords())
	fmt.Println(getCoordsAlt())

	fmt.Println("-------------------")

	//assignement funcs
	untilAge(4)
	untilAge(10)
	untilAge(18)
	untilAge(22)
	untilAge(34)

	fmt.Println("-------------------")

	//explicit return
	fmt.Println(getCoordsAlt2())
	untilAgeAlt(4)
	untilAgeAlt(22)
	untilAgeAltImplicit(4)
	untilAgeAltImplicit(22)

	fmt.Println("-------------------")

	//guard clause
	fmt.Println(divide(89, 6))
	fmt.Println(divide(89, 0))

	fmt.Println("-------------------")

	//structs
	testStruct(messageToSendSimple{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	testStruct(messageToSendSimple{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	testStruct(messageToSendSimple{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
	myCar := car{}
	myCar.FrontWheel.Radius = 5
	println(myCar.BackWheel.Material)
	println(myCar.Height)

	//nested structs
	testNestedStruct(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	testNestedStruct(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 19035558973,
		},
	})
	testNestedStruct(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name: "Njorn Halafax",
		},
		recipient: user{
			name:   "Becky Sue",
			number: 19035558973,
		},
	})
	testNestedStruct(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 16545550987,
		},
		recipient: user{
			number: 19035558973,
		},
	})
	testNestedStruct(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Jason Bjorn",
			number: 16545550987,
		},
		recipient: user{
			name: "Jim Bond",
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
	var lanesTruck = truck{
		BedSize: 10,
		car: car{
			Make:  "toyota",
			Model: "camry",
		},
	}

	fmt.Println(lanesTruck.BedSize)

	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.Make
	fmt.Println(lanesTruck.Make)
	fmt.Println(lanesTruck.Model)

	testEmb(sender{
		rateLimit: 10000,
		user_emb: user_emb{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	testEmb(sender{
		rateLimit: 5000,
		user_emb: user_emb{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	testEmb(sender{
		rateLimit: 1000,
		user_emb: user_emb{
			name:   "Sally",
			number: 19055558790,
		},
	})

	// methods associated to structs

	var r = rect{
		width:  5,
		height: 10,
	}

	fmt.Println(r.area())
	// prints 50
	testAuth(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	testAuth(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	testAuth(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
