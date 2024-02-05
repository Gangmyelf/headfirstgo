// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"errors"
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

func GetInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	finalInput, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(finalInput), nil
}

func GetString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
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
	numbers := make([]float64, 3)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
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

func ReadArg(arrayOfFloat []string) (float64, error) {
	argument := arrayOfFloat[1:]
	var sum float64
	for _, value := range argument {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	if sum == 0 {
		err := errors.New("Вы ввели нулевое значение")
		log.Fatal(err)
	}
	return sum / float64(len(argument)), nil
}

func Maximum(num ...float64) (float64, error) {
	MaxNum := num[0]
	for _, number := range num {
		if number > MaxNum {
			MaxNum = number
		}
	}
	return MaxNum, nil
}

func InRange(min float64, max float64, rangeOther ...float64) []float64 {
	result := []float64{}
	for _, numbers := range rangeOther {
		if numbers <= max && numbers >= min {
			result = append(result, numbers)
		}
	}
	return result
}

func Strings(filename string) ([]string, error) {
	file, err := os.Open(filename)
	fullnames := []string{}
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		fullnames = append(fullnames, name)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
	return fullnames, nil
}

func CollectStrings(text ...string) ([]string, []int) {
	counts := []int{}
	names := []string{}

	for _, line := range text {
		matched := false
		for i, name := range names {
			if name == line {
				//fmt.Printf("The count is %d\n", counts[i])
				counts[i]++
				//fmt.Printf("The count is %d after ++\n", counts[i])
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	return names, counts
}

func CollectStringsWithMaps(text ...string) map[string]int {
	ranks := map[string]int{}

	for _, line := range text {
		ranks[line]++
	}
	return ranks
}

func ReadFile_string(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	strings := make([]string, 0)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		strings = append(strings, scanner.Text())
		//fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return strings, err
}
