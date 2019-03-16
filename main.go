package main

import (
	"github.com/lee5i3/clutterfinder-auth-get-token/auth"
	"fmt"
)

func main() {
	val := auth.getToken()

	fmt.Println(val)
}
