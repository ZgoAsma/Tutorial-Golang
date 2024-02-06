package ranges

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, bad := range badWords {
			if word == bad {
				return i
			}
		}
	}
	return -1
}

// don't touch below this line

func Test(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}
