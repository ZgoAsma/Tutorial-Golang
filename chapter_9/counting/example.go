package counting

import (
	"fmt"
)

func getCounts(userIDs []string) map[string]int {
	res := map[string]int{}
	for _, user_id := range userIDs {
		_, ok := res[user_id]
		if !ok {
			res[user_id] = 1
		}
		if ok {
			res[user_id] += 1
		}
	}
	return res
}

// don't edit below this line

func Test(userIDs []string, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)
	}
	fmt.Println("=====================================")
}
