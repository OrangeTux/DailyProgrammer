package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Show user prompt with given text and return his input.
func prompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	equation, _ := reader.ReadString('\n')

	return equation
}

// Parse a string containg equation in format `ax + b` and return a and b.
func parse_equation(equation string) (a, b float64) {
	re := regexp.MustCompile("[+-]?[\\d]*\\.?[\\d]+")
	result := re.FindAllString(equation, -1)

	if len(result) != 1 && len(result) != 2 {
		fmt.Printf("%s does not match format `ax+b`.\n", equation)
		os.Exit(1)
	}

	a, _ = strconv.ParseFloat(result[0], 32)

	if len(result) == 2 {
		b, _ = strconv.ParseFloat(result[1], 32)
	}

	// When no b value has been geven we b is 0, because of Go's zero default.
	return a, b
}

// Find intersection of 2 linear functions.
func find_intersection(f_1, f_2 string) (x, y float64) {
	f_1_a, f_1_b := parse_equation(f_1)
	f_2_a, f_2_b := parse_equation(f_2)

	x = find_x_intersect(f_1_a, f_2_a, f_1_b, f_2_b)
	y = find_y_intersect(f_1_a, f_1_b, x)

	return x, y
}

// Return y part of coordinate.
func find_y_intersect(a, b, x float64) (y float64) {
	return (a * x) + b
}

// Return x part of coordinate
func find_x_intersect(a_1, a_2, b_1, b_2 float64) (x float64) {
	return (b_1 - b_2) / (a_2 - a_1)
}

func main() {
	equation_1 := prompt("Enter equation in format 'ax + b': ")
	equation_2 := prompt("Enter another equation: ")
	x, y := find_intersection(equation_1, equation_2)
	fmt.Printf("(%f, %f)\n", x, y)
}
