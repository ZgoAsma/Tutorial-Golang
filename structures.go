package main

import "fmt"

type car_simple struct {
	Make   string
	Model  string
	Height int
	Width  int
}

type messageToSendSimple struct {
	phoneNumber int
	message     string
}

type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

type car2 struct {
	Make   string
	Model  string
	Height int
	Width  int
	// Wheel is a field containing an anonymous struct
	Wheel struct {
		Radius   int
		Material string
	}
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car
	BedSize int
}

type sender struct {
	rateLimit int
	user_emb
}

type user_emb struct {
	name   string
	number int
}

type rect struct {
	width  int
	height int
}

type authenticationInfo struct {
	username string
	password string
}

func (i authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", i.username, i.password)
}
