package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/bryanpburnside/advent_of_code_2024/utils"
)

func main() {
	input, err := utils.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	output := getDifference(input)
	fmt.Println(output)
}

func getDifference(input []string) int {
	difference := int(0)
	leftList := []int{}
	rightList := []int{}

	for _, line := range input {
		split := strings.Split(line, "   ")
		numberLeft, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		numberRight, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		leftList = append(leftList, numberLeft)
		rightList = append(rightList, numberRight)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i, leftNumber := range leftList {
		rightNumber := rightList[i]
		difference += utils.Abs(leftNumber - rightNumber)
	}

	return difference
}
