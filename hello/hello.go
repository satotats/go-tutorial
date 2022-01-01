package main

import (
	"fmt"

	"satotats/greetings"
)

func main() {
	message := greetings.Hello("コニチハ")
	fmt.Println(message)
}
