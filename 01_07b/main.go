package main

import (
	"flag"
	"log"
)

// operatorType represents the type of operator in an expression
type operatorType int

const (
	openBracket operatorType = iota  // Open bracket type
	closedBracket                    // Closed bracket type
	otherOperator                    // Other type (non-bracket)
)

// bracketPairs is the map of legal bracket pairs
var bracketPairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// getOperatorType returns the operator type of the given operator.
func getOperatorType(op rune) operatorType {
	for ob, cb := range bracketPairs {
		switch op {
		case ob:
			return openBracket
		case cb:
			return closedBracket
		}
	}

	return otherOperator
}

// stack is a simple LIFO stack implementation using a slice.
type stack struct {
	elems []rune
}

// push adds a new element to the stack.
func (s *stack) push(e rune) {
	s.elems = append(s.elems, e)
}

// pop removes the last element from the stack.
func (s *stack) pop() *rune {
	if len(s.elems) == 0 {
		return nil
	}
	n := len(s.elems) - 1
	last := s.elems[n]
	s.elems = s.elems[:n]
	return &last
}

// isBalanced returns whether the given expression has balanced brackets.
func isBalanced(expr string) bool {
	s := stack{}  // Initialize an empty stack
	for _, e := range expr {
		switch getOperatorType(e) {
		case openBracket:
			s.push(e)  // Push open bracket onto the stack
		case closedBracket:
			last := s.pop()  // Pop the last bracket from the stack
			if last == nil || bracketPairs[*last] != e {
				return false  // If no matching open bracket, expression is not balanced
			}
		}
	}

	return len(s.elems) == 0  // Expression is balanced if stack is empty at the end
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))  // Check if the expression is balanced and print the result
}