package chapter_5

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type Contractor struct {
	Name         string
	HourlyPay    int
	HoursPerYear int
}

func (c Contractor) getName() string {
	return c.Name
}

func (c Contractor) getSalary() int {
	return c.HourlyPay * c.HoursPerYear
}

// ?

// don't touch below this line

type FullTime struct {
	Name   string
	Salary int
}

func (ft FullTime) getSalary() int {
	return ft.Salary
}

func (ft FullTime) getName() string {
	return ft.Name
}

func Test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}
