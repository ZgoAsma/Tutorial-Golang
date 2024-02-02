package chapter_5

import (
	"fmt"
)

func (e Email) cost() float64 {
	if e.IsSubscribed {
		return float64(len(e.Body)) * 0.01
	}
	return float64(len(e.Body)) * 0.05
}

func (e Email) Print() {
	fmt.Println(e)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	Print()
}

type Email struct {
	IsSubscribed bool
	Body         string
}

func Print(p printer) {
	p.Print()
}

func TestPrint(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.Print()
	fmt.Println("====================================\n")
}

type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

type CopierAlt interface {
	Copy(string, string) int
}
