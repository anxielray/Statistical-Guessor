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

// Function to guess the range for the next number
func guess_it_1(mean, stddev float64) (lowerLimit, upperLimit float64) {
	lowerLimit = mean - (3 * stddev)
	upperLimit = mean + (3 * stddev)
	return
}

// Window size for calculation
const window int = 8

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []float64

	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error converting to number:", err)
			os.Exit(1)
		}

		// Convert to float64 and append to numbers
		numbers = append(numbers, number)
		if len(numbers) > window {
			numbers = numbers[len(numbers)-window:]
		}

		// Provide a range for a single number
		// if len(numbers) == 1 {
		// 	lowerLimit := numbers[0] - 10 // Use a reasonable default range
		// 	upperLimit := numbers[0] + 10
		// 	fmt.Printf("%d %d\n", int(math.Round(lowerLimit)), int(math.Round(upperLimit)))
		// }
		if len(numbers) > 1 {
			mean, stddev := calculateStatistics(numbers)
			lowerLimit, upperLimit := guess_it_1(mean, stddev)
			fmt.Printf("%d %d\n", int(math.Round(lowerLimit)), int(math.Round(upperLimit)))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}

//Hey Reimund came up with this logic!
//Thic comment was staged and commited by the action of automation
