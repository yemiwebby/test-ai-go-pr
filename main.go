package main

import "fmt"

func updateNumber(n *int) int {
	*n = 10
	return *n
}

func main() {
	number := 5
	update := fmt.Sprintf("the update value is: %d", updateNumber(&number))
	fmt.Println(update)
}
