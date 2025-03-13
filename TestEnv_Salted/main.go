package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	r := rand.Text()
	fmt.Println(r)
}
