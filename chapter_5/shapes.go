package chapter_5

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Square struct {
}
