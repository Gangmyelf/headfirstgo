// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"log"
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

// Average got array []float64
// And write the average of all num.
func Average(numbers []float64) {
	var sum float64 = 0
	for _, value := range numbers {
		sum += value
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n ", sum/sampleCount)
}

// ReadFile got the name of text file and
// return []float64 slice
func ReadFile(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	numbers := []float64{}
	count := 0
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[count], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		count++
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return numbers, err
}
