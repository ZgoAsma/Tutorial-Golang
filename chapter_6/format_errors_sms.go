package chapter_6

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	s := fmt.Sprintf("SMS that costs $%.1f to be sent to '%s' can not be sent", cost, recipient)
	//fmt.Printf(s)
	return s
}

// don't edit below this line

func TestFormatting(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}
