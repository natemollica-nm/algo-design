package main

import (
	"fmt"
	"strconv"
)

func binaryResult(num int) string {
	if num == 0 {
		return "0"
	}

	binary := ""

	for num > 0 {
		remainder := num % 2
		binary = strconv.Itoa(remainder) + binary
		num = num / 2
	}

	return binary
}

func linearPatternMatch(text, pattern string) {
	N := len(text)
	M := len(pattern)
	found := false

	for i := 0; i <= N-M; i++ {
		match := true
		for j := 0; j < M; j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			fmt.Printf("Pattern found at index %d\n", i)
			found = true
		}
	}

	if !found {
		fmt.Println("Pattern not found in the text.")
	}
}

// checkGraphProperties checks the three properties of the graph
func checkGraphProperties(adjMatrix [][]int) (isComplete bool, hasSelfLoop bool, hasIsolatedVertex bool) {
	N := len(adjMatrix)
	isComplete = true
	hasSelfLoop = false
	hasIsolatedVertex = false

	for i := 0; i < N; i++ {
		rowSum := 0
		for j := 0; j < N; j++ {
			if i == j && adjMatrix[i][j] == 1 {
				hasSelfLoop = true
			}
			if i != j && adjMatrix[i][j] == 0 {
				isComplete = false
			}
			rowSum += adjMatrix[i][j]
		}
		if rowSum == 0 {
			hasIsolatedVertex = true
		}
	}

	return
}

func main() {
	var num int
	fmt.Print("Enter a decimal number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	binary := binaryResult(num)
	fmt.Printf("The binary equivalent of %d is %s\n", num, binary)

	var text, pattern string

	// Prompt the user for input
	fmt.Print("Enter the text: ")
	fmt.Scanln(&text)
	fmt.Print("Enter the pattern to search for: ")
	fmt.Scanln(&pattern)

	// Perform the pattern search
	linearPatternMatch(text, pattern)

	// Example adjacency matrix
	adjMatrix := [][]int{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	// Check the properties of the graph
	isComplete, hasSelfLoop, hasIsolatedVertex := checkGraphProperties(adjMatrix)

	// Output the results
	fmt.Printf("Graph is complete: %v\n", isComplete)
	fmt.Printf("Graph has a self-loop: %v\n", hasSelfLoop)
	fmt.Printf("Graph has an isolated vertex: %v\n", hasIsolatedVertex)
}
