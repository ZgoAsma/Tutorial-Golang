package chapters_1_4

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type BirthdayMessage struct {
	BirthdayTime  time.Time
	RecipientName string
}

func (bm BirthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.RecipientName, bm.BirthdayTime.Format(time.RFC3339))
}

type SendingReport struct {
	ReportName    string
	NumberOfSends int
}

func (sr SendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.ReportName, sr.NumberOfSends)
}

func TestRec(m message) {
	sendMessage(m)
	fmt.Println("====================================\n")
}
