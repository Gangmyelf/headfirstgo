// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a floating-point number from the keyboard.
// It returns the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	finalInput, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return finalInput, nil
}

func Average(numbers []float64) {
	var sum float64 = 0
	for _, value := range numbers {
		sum += value
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n ", sum/sampleCount)
}
