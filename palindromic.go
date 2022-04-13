package main

import (
	"flag"
	"math/big"
	"os"
	"strconv"
	"time"
)

func main() {

	number := flag.Int("number", 196, "number to be calculated")
	delay := flag.Int("delay", 1, "delay in milliseconds")
	iterations := flag.Int("iterations", 1000, "number of iterations")

	flag.Parse()

	bigint := big.NewInt(int64(*number))

	calculatePalindromicNumber(*bigint, *delay, *iterations, 0, time.Now())
}

func calculatePalindromicNumber(number big.Int, delay int, iterations int, currentIteration int, initialTime time.Time) int {

	numberString := number.String()

	reversedNumberString := reverseString(numberString)

	reversedNumber := new(big.Int)
	reversedNumber.SetString(reversedNumberString, 10)

	sum := new(big.Int)
	sum.Add(&number, reversedNumber)

	sumString := sum.String()

	println(numberString + " + " + reversedNumberString + " = " + sumString)

	os.Stdout.WriteString("\033]0;" + "Iteration " + strconv.Itoa(currentIteration) + " of " + strconv.Itoa(iterations) + "\007")

	if sumString == reverseString(sumString) {
		println("Number is palindromic in " + strconv.Itoa(currentIteration+1) + " iterations")
		println("Time elapsed: " + time.Since(initialTime).String())
	} else {

		if currentIteration < iterations {
			if delay > 0 {
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}
			calculatePalindromicNumber(*sum, delay, iterations, currentIteration+1, initialTime)
		} else {
			println("Number is not palindromic in " + strconv.Itoa(iterations) + " iterations")
			println("Time elapsed: " + time.Since(initialTime).String())
		}
	}
	return 0
}

func reverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
