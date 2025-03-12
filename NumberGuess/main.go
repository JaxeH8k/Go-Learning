package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 100
	num := rand.Intn(maxNum)
	//fmt.Println(num)
	guessCount := 1
	var guess int
	fmt.Println("Pick a number between 0 & 100")
	fmt.Scanln(&guess)
	for guess != num {
		if guessCount >= 10 {
			fmt.Println("LOSER!  YOU LOSE!")
			break
		}
		switch {
		case guess < num:
			fmt.Println("higher")
			fmt.Scanln(&guess)
		case guess > num:
			fmt.Println("lower")
			fmt.Scanln(&guess)
		}
		guessCount++
	}

	if guessCount <= 5 {
		fmt.Printf("Great stuff! It was %d\n", num)
	}
	if guessCount <= 10 && guessCount >= 6 && num == guess {
		fmt.Println("Got it, number was ", num)
	}
}
