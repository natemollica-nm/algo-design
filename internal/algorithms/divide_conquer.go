package algorithms

import "math"

type Point struct {
	x float64
}

// Helper function to calculate the distance between two points
func distance(p1, p2 Point) float64 {
	return math.Abs(p1.x - p2.x)
}

// Merge function to merge two halves in sorted order
func merge(left, right []Point) []Point {
	result := make([]Point, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].x < right[j].x {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// MergeSort function that sorts the array of points
func mergeSort(points []Point) []Point {
	if len(points) <= 1 {
		return points
	}

	// Divide the array into two halves
	mid := len(points) / 2
	left := mergeSort(points[:mid])
	right := mergeSort(points[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

// ClosestPair function implements the divide and conquer algorithm
func ClosestPair(points []Point) float64 {
	n := len(points)
	if n <= 1 {
		return math.Inf(1)
	}
	if n == 2 {
		return distance(points[0], points[1])
	}

	// Sort the points by X-coordinate using our Merge Sort implementation
	points = mergeSort(points)

	// Divide
	mid := n / 2
	leftPoints := points[:mid]
	rightPoints := points[mid:]

	// Conquer
	d1 := ClosestPair(leftPoints)
	d2 := ClosestPair(rightPoints)

	// Combine
	return math.Min(d1, d2)
}

// CountOccurrences is a helper function that implements the divide and conquer approach
func CountOccurrences(A []int, low, high, Q int) int {
	// Base case: if the array has only one element
	if low == high {
		if A[low] == Q {
			return 1
		} else {
			return 0
		}
	}

	// Calculate the middle index
	mid := (low + high) / 2

	// Recursively count occurrences in the left and right halves
	leftCount := CountOccurrences(A, low, mid, Q)
	rightCount := CountOccurrences(A, mid+1, high, Q)

	// Combine the results from both halves
	return leftCount + rightCount
}

// CountOccurrencesInArray is the main function that initiates the divide and conquer process
//
// Base Case: If the array (or subarray) contains only one element, we directly check if it equals Q. If it does, we return 1 (indicating one occurrence); otherwise, we return 0.
//
// Divide:    The array is divided into two halves by calculating the middle index mid.
//
// Conquer:   The function recursively counts the occurrences of Q in both the left and right halves.
//
// Combine:   The results from the left and right halves are summed to get the total number of occurrences of Q in the current subarray.
func CountOccurrencesInArray(A []int, Q int) int {
	N := len(A)
	if N == 0 {
		return 0
	}

	// Start the divide and conquer process from the entire array
	return CountOccurrences(A, 0, N-1, Q)
}
