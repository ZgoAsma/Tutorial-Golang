package chapters_1_4

import (
	"errors"
	"fmt"
)

//functions

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func Test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

// functions syntaxic sugar
func AddToDatabase(hp, damage int, name string, level int) {
	fmt.Printf("hp=%v added to db!\n", hp)
	fmt.Printf("damage=%v added to db!\n", damage)
	fmt.Printf("name=%s added to db!\n", name)
	fmt.Printf("level=%v added to db!\n", level)
}

//passing variables by values

func Sending() {
	sendsSoFar := 430
	const sendsToAdd = 25
	incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println(`you've sent`, sendsSoFar, "messages. \nIncorrect!")
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages.")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

// ignoring return values
func Names() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

func getNames() (string, string) {
	return "John", "Doe"
}

// implicit vs explicit
func GetCoords() (x, y int) {
	// x and y are initialized with zero values

	return // automatically returns x and y (implicit)
}

func GetCoordsAlt() (int, int) {
	var x int
	var y int
	return x, y //(explicit)
}

func GetCoordsAlt2() (x, y int) {
	return 5, 6 // this is explicit, x and y are NOT returned
}

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func yearsUntilEventsAlt(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return 0, 0, 0
}

func yearsUntilEventsAltImplicit(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}
func UntilAge(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEventsAlt(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====")
}

func UntilAgeAlt(age int) {
	fmt.Println("Explicit+constant. \nAge:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEventsAlt(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("===")
}

func UntilAgeAltImplicit(age int) {
	fmt.Println("Implicit.\nAge:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEventsAltImplicit(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("===")
}

// guard clause
func Divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New(`Can't divide by zero!`)
	}
	return dividend / divisor, nil
}

// structs
func TestStruct(m MessageToSendSimple) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.Message, m.PhoneNumber)
	fmt.Println("====================================")
}

func canSendMessage(mToSend MessageToSend) bool {
	if mToSend.Sender.Name == "" {
		return false
	}
	if mToSend.Sender.Number == 0 {
		return false
	}
	if mToSend.Recipient.Name == "" {
		return false
	}
	if mToSend.Recipient.Number == 0 {
		return false
	}
	return true
}
func TestNestedStruct(mToSend MessageToSend) {
	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
		mToSend.Message,
		mToSend.Sender.Name,
		mToSend.Sender.Number,
		mToSend.Recipient.Name,
		mToSend.Recipient.Number,
	)
	fmt.Println()
	if canSendMessage(mToSend) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("...can't send message")
	}
	fmt.Println("====================================")
}

func TestEmb(s Sender) {
	fmt.Println("Sender name:", s.Name)
	fmt.Println("Sender number:", s.Number)
	fmt.Println("Sender rateLimit:", s.RateLimit)
	fmt.Println("====================================")
}

// this is a method : area has a receiver of (r rect)
func (r Rect) Area() int {
	return r.Width * r.Height
}

func TestAuth(authInfo AuthenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}
