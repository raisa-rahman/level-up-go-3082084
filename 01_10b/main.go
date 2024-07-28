package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
// Returns a pointer to float64 and an error if parsing fails
func parseOperand(op string) (*float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)
	if err != nil {
		// Return nil and an error formatted with details of the parsing error
		return nil, fmt.Errorf("cannot parse: %v", err)
	}

	// Return a pointer to the parsed float64 and a nil error if parsing is successful
	return &parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
// Returns a pointer to float64 and an error if calculation fails
func calculate(expr string) (*float64, error) {
	// Split the expression into operands and operator
	ops := strings.Fields(expr)
	nops := len(ops)
	
	// Check if the number of operands and operator is exactly 3 (two operands and one operator)
	if nops != 3 {
		// Return nil and an error if the expression does not contain exactly 3 parts
		return nil, fmt.Errorf("cannot calculate: need 3 ops, got %d", nops)
	}

	// Parse the left operand
	left, err := parseOperand(ops[0])
	if err != nil {
		// Return nil and the error if parsing the left operand fails
		return nil, err
	}

	// Parse the right operand
	right, err := parseOperand(ops[2])
	if err != nil {
		// Return nil and the error if parsing the right operand fails
		return nil, err
	}

	// Retrieve the function for the operator from the operators map
	f, ok := operators[ops[1]]
	if !ok {
		// Return nil and an error if the operator is not found in the map
		return nil, fmt.Errorf("cannot calculate: %s is unknown", ops[1])
	}

	// Compute the result using the operator function
	result := f(*left, *right)
	
	// Return a pointer to the result and a nil error
	return &result, nil
}

func main() {
	// Define the command line flag for the expression
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")
	flag.Parse()

	// Call the calculate function with the provided expression
	result, err := calculate(*expr)
	if err != nil {
		// Print the error and exit if there is a calculation error
		log.Fatal(err)
	}

	// Print the result of the calculation, formatted to two decimal places
	log.Printf("%s = %.2f\n", *expr, *result)
}