package main

import "example/hello/chapter_5"

func main() {
	e := chapter_5.Email{
		IsSubscribed: true,
		Body:         "hello there",
	}
	chapter_5.TestPrint(e, e)
	e = chapter_5.Email{
		IsSubscribed: false,
		Body:         "I want my money back",
	}
	chapter_5.TestPrint(e, e)
	e = chapter_5.Email{
		IsSubscribed: true,
		Body:         "Are you free for a chat?",
	}
	chapter_5.TestPrint(e, e)
	e = chapter_5.Email{
		IsSubscribed: false,
		Body:         "This meeting could have been an email",
	}
	chapter_5.TestPrint(e, e)
}
