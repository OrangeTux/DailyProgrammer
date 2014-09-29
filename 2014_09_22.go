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

// Find intersection of 2 linear functions.
func find_intersection(f_1, f_2 string) (x, y int) {
	//
	f_1_a, f_1_b := parse_equation(f_1)
	f_2_a, f_2_b := parse_equation(f_2)

	y = find_y_intersect(f_1_b, f_2_b)
	x = find_x_intersect(f_1_a, f_2_a, y)

	return x, y
}

// Parse a string containg equation in format `ax + b` and return a and b.
func parse_equation(equation string) (a, b int) {
	re := regexp.MustCompile("[0-9]*")
	result := re.FindAllString(equation, -1)

	//fmt.Printf("%v", len(result))
	if len(result) != 4 {
		fmt.Printf("%s does not match format `ax+b`.\n", equation)
		os.Exit(1)
	}

	a, _ = strconv.Atoi(result[0])
	b, _ = strconv.Atoi(result[2])

	//fmt.Printf("%s: a = %v and b = %v.\n", equation, a, b)
	return a, b
}

// Return y part of coordinate.
func find_y_intersect(b_1, b_2 int) (y int) {
	return b_1 - b_2
}

// Return x part of coordinate
func find_x_intersect(a_1, a_2, y_intersect int) (x int) {
	return y_intersect / (a_2 - a_1)
}

func main() {
	equation_1 := prompt("Enter equation in format 'ax + b': ")
	equation_2 := prompt("Enter another equation: ")
	x, y := find_intersection(equation_1, equation_2)
	fmt.Printf("(%d, %d)\n", x, y)
}
