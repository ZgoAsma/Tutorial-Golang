package chapter_6

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	m1, err := sendSMS(msgToCustomer)
	if err != nil {
		fmt.Println("couldnt send msg to customer")
		//fmt.Println(err.Error())
		return 0, err
	}

	m2, err := sendSMS(msgToSpouse)
	if err != nil {
		fmt.Println("couldn't send msg to spouse")
		//fmt.Println(err.Error())
		return 0, err
	}

	return m1 + m2, nil
}

// don't edit below this line

func sendSMS(message string) (float64, error) {
	//const maxTextLen = 25
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func Test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}
