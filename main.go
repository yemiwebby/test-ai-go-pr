package main

import (
	"fmt"
)

func updateNumber(n *int) int {
	*n = 10
	return *n
}

func main() {
	number := 20
	update := fmt.Sprintf("the update value is: %d", updateNumber(&number))
	fmt.Println("update", update)
}

// func CountDown() {
// 	start := 3
// 	for {
// 		if start == 0 {
// 			fmt.Println("Go!")
// 			return
// 		}
// 		fmt.Println(start)
// 		time.Sleep(time.Second * 1)
// 		start--
// 	}
// }
