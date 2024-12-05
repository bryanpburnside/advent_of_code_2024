package main

import (
	"fmt"
	"log"

	"github.com/bryanpburnside/advent_of_code_2024/utils"
)

func main() {
	input, err := utils.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Puzzle input:", input)
}
