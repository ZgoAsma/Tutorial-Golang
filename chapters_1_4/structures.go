package chapters_1_4

import "fmt"

type car_simple struct {
	Make   string
	Model  string
	Height int
	Width  int
}

type MessageToSendSimple struct {
	PhoneNumber int
	Message     string
}

type Car struct {
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

type MessageToSend struct {
	Message   string
	Sender    User
	Recipient User
}

type User struct {
	Name   string
	Number int
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

type Truck struct {
	// "Car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	Car
	BedSize int
}

type Sender struct {
	RateLimit int
	User_emb
}

type User_emb struct {
	Name   string
	Number int
}

type Rect struct {
	Width  int
	Height int
}

type AuthenticationInfo struct {
	Username string
	Password string
}

func (i AuthenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", i.Username, i.Password)
}
