package type_switches

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	switch exp := e.(type) {
	case Email:
		return exp.ToAddress, exp.cost()
	case Sms:
		return exp.ToPhoneNumber, exp.cost()
	default:
		return "", 0.0
	}

}

// don't touch below this line

func (e Email) cost() float64 {
	if !e.IsSubscribed {
		return float64(len(e.Body)) * .05
	}
	return float64(len(e.Body)) * .01
}

func (s Sms) cost() float64 {
	if !s.IsSubscribed {
		return float64(len(s.Body)) * .1
	}
	return float64(len(s.Body)) * .03
}

func (i Invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type Email struct {
	IsSubscribed bool
	Body         string
	ToAddress    string
}

type Sms struct {
	IsSubscribed  bool
	Body          string
	ToPhoneNumber string
}

type Invalid struct{}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func TestSwitch(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case Email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case Sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}
