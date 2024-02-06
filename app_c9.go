package main

import (
	"crypto/md5"
	"example/hello/chapter_9"
	"example/hello/chapter_9/counting"
	"example/hello/chapter_9/mutations"
	"fmt"
	"io"
	"sort"
)

func main() {

	fmt.Println("********************Trivial********************")
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // overwrites 24

	ages = map[string]int{
		"John": 37,
		"Mary": 21,
	}
	fmt.Println(len(ages)) // 2

	fmt.Println("********************App Maps********************")
	chapter_9.Test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	chapter_9.Test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	chapter_9.Test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)

	users := map[string]mutations.User{
		"john": {
			Name:                 "john",
			Number:               18965554631,
			ScheduledForDeletion: true,
		},
		"elon": {
			Name:                 "elon",
			Number:               19875556452,
			ScheduledForDeletion: true,
		},
		"breanna": {
			Name:                 "breanna",
			Number:               98575554231,
			ScheduledForDeletion: false,
		},
		"kade": {
			Name:                 "kade",
			Number:               10765557221,
			ScheduledForDeletion: false,
		},
	}
	mutations.Test(users, "john")
	mutations.Test(users, "musk")
	mutations.Test(users, "santa")
	mutations.Test(users, "kade")

	keys := []string{}
	for name := range users {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	fmt.Println("Final map keys:")
	for _, name := range keys {
		fmt.Println(" - ", name)
	}
	fmt.Println("********************Key types********************")
	fmt.Println("map of maps")
	hits1 := make(map[string]map[string]int)

	chapter_9.Add(hits1, "/doc/", "au")
	n := hits1["/doc/"]["au"]

	println(n)

	type Key struct {
		Path, Country string
	}
	hits := make(map[Key]int)
	hits[Key{"/", "vn"}]++
	n = hits[Key{"/ref/spec", "ch"}]

	println(n)

	names := map[string]int{}

	if _, ok := names["elon"]; !ok {
		// if the key doesn't exist yet,
		// initialize its value to 0
		names["elon"] = 0
	}
	names["asma"] = 0
	fmt.Println(names)
	userIDs := []string{}
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:2])
	}

	counting.Test(userIDs, []string{"00", "ff", "dd"})
	counting.Test(userIDs, []string{"aa", "12", "32"})
	counting.Test(userIDs, []string{"bb", "33", "bb"})

}
