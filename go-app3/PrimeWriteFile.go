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

	if *isPrime {
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

func readFile(fileName string) []byte {
	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	return content

}

func readFileInArray(fileName string) []int {
	var content []byte = readFile(fileName)
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

func Shuffle(array []int) {
	rand.Shuffle(len(array), func(i, j int) {
		array[i], array[j] = array[j], array[i]
	})
}

func main() {
	var fileName = getFileName()
	var Nums []int = readFileInArray(fileName)
	Shuffle(Nums)
	f, err := os.Create("results.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, num := range Nums {
		if isPrimeNum(num) {
			fmt.Println(num, " is prime")
		} else {
			fmt.Println(num, " is not prime")
		}
		writeFile(f, num)
	}
	defer f.Close()
}

func writeFile(f *os.File, num int) {
	if isPrimeNum(num) {
		var msg string
		msg = strconv.Itoa(num) + " => PRIME" + "\n"
		_, err2 := f.WriteString(msg)
		if err2 != nil {
			log.Fatal(err2)
		}
	} else {
		var msg string
		msg = strconv.Itoa(num) + " => NOTPRIME" + "\n"
		_, err2 := f.WriteString(msg)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
