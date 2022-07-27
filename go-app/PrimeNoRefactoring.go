package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	fmt.Println("Enter the file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileLines := lineBreakRegExp.Split(string(content), -1)
	var Nums []int
	for _, line := range fileLines {
		intVar, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		Nums = append(Nums, intVar)
	}
	rand.Shuffle(len(Nums), func(i, j int) {
		Nums[i], Nums[j] = Nums[j], Nums[i]
	})
	for _, num := range Nums {
		if num < 2 {
			fmt.Println("Number must be greater than 2.")
		}
		sq_root := int(math.Sqrt(float64(num)))
		var isPrime *bool
		isPrime = new(bool)
		*isPrime = true
		for i := 2; i <= sq_root; i++ {
			if num%i == 0 {
				*isPrime = false
				break
			}
		}
		if *isPrime == true {
			fmt.Println(num, " is prime")
		} else {
			fmt.Println(num, " is not prime")
		}

	}

}
