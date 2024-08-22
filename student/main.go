package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Calculates mean and standard deviation of a slice of floats
func calculateStatistics(numbers []float64) (mean, stddev float64) {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	mean = sum / float64(len(numbers))

	var variance float64
	for _, num := range numbers {
		variance += (num - mean) * (num - mean)
	}
	variance /= float64(len(numbers))
	stddev = math.Sqrt(variance)
	return
}

func guess_it_1(mean, stddev float64) (lowerLimit, upperLimit float64) {
	lowerLimit = mean - (3 * stddev)
	upperLimit = mean + (3 * stddev)
	return
}

const window int = 8

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []float64
	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("Error converting to number:", err)
			os.Exit(1)
		}

		// use a window slide to be more precise in the approximation
		numbers = append(numbers, float64(number))
		if len(numbers) > window {
			numbers = numbers[len(numbers)-window:]
		}

		// ignore the calculation if the number provided is only 1, should be 2 and above...
		if len(numbers) > 1 {
			mean, stddev := calculateStatistics(numbers)
			lowerLimit, upperLimit := guess_it_1(mean, stddev)
			fmt.Printf("%d %d", int(lowerLimit), int(upperLimit))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}
