package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading env file %v\n", err)
	}
	apiKey := os.Getenv("API_KEY")
	fmt.Println(apiKey)

}
