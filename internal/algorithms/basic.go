package algorithms

import (
	"fmt"
	"strconv"
)

func BinaryResult(num int) string {
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

func LinearPatternMatch(text, pattern string) {
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

// CheckGraphProperties checks the three properties of the graph
func CheckGraphProperties(adjMatrix [][]int) (isComplete bool, hasSelfLoop bool, hasIsolatedVertex bool) {
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
