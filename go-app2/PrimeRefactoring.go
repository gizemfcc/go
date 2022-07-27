package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func isPrimeNum(num int) bool {
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
		return true
	} else {
		return false
	}
}
func getFileName() string {
	fmt.Println("Enter the file name: ")
	var fileName string
	fmt.Scanln(&fileName)
	return fileName
}

func readFileInArray(fileName string) []int {
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
	return Nums
}

func main() {
	var fileName = getFileName()
	var Nums []int = readFileInArray(fileName)
	for _, num := range Nums {

		var isPrime *bool
		isPrime = new(bool)
		*isPrime = isPrimeNum(num)
		if *isPrime == true {
			fmt.Println(num, " is prime")
		} else {
			fmt.Println(num, " is not prime")
		}
	}
}
