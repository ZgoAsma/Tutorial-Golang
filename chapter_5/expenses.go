package chapter_5

import (
	"fmt"
)

func getExpenseReport(e Expense) (string, float64) {
	em, ok := e.(EmailAlt)
	if ok {
		return em.ToAddress, em.cost()
	}
	SMSe, ok := e.(SMS)

	if ok {
		return SMSe.ToPhoneNumber, SMSe.cost()
	}
	return "", 0
}

// don't touch below this line

func (e EmailAlt) cost() float64 {
	if !e.IsSubscribed {
		return float64(len(e.Body)) * .05
	}
	return float64(len(e.Body)) * .01
}

func (s SMS) cost() float64 {
	if !s.IsSubscribed {
		return float64(len(s.Body)) * .1
	}
	return float64(len(s.Body)) * .03
}

func (i Invalid) cost() float64 {
	return 0.0
}

type Expense interface {
	cost() float64
}

type EmailAlt struct {
	IsSubscribed bool
	Body         string
	ToAddress    string
}

type SMS struct {
	IsSubscribed  bool
	Body          string
	ToPhoneNumber string
}

type Invalid struct{}

func estimateYearlyCost(e Expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func TestPhones(e Expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case EmailAlt:
		fmt.Printf("Report: The Email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case SMS:
		fmt.Printf("Report: The SMS going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid Expense")
		fmt.Println("====================================")
	}
}
