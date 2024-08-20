package algorithms

import "sort"

// RussianPeasantMultiplication
// The Russian Peasant Multiplication method is an ancient algorithm for multiplying
// two integers. It works by repeatedly halving one number and doubling the other,
// then summing certain values based on whether the halved numbers are odd.
func RussianPeasantMultiplication(a, b int) int {
	result := 0

	// While 'a' > 0
	for a > 0 {
		// If 'a' is odd, add 'b' to the result
		if a%2 != 0 {
			result += b
		}

		// Halve 'a' and double 'b'
		a /= 2
		b *= 2
	}

	return result
}

// FindElementIn2DArray
// Finds the position of Q in a 2D matrix A
// Returns the indices (i, j) if found, or (-1, -1) if not found
//
//	A - Input 2D Array
//	N - Depth/Size of array dimension
//	Q - Query value for 2D array
//
// TimeComplexity: O(N)
//   - Worst Case: Algorithm moves from top-right to bottom-left corner.
//   - This means 'i' or 'j' is incremented/decremented each iteration
func FindElementIn2DArray(A [][]int, N int, Q int) (int, int) {
	i, j := 0, N-1

	for i < N && j >= 0 {
		if A[i][j] == Q {
			return i, j
		} else if A[i][j] > Q {
			j-- // Move left
		} else {
			i++ // Move down
		}
	}

	return -1, -1 // Element not found
}

// FindFixedPoint
// Binary search algorithm where the idea is to leverage the fact that the
// array is sorted and contains distinct integers to efficiently narrow
// down the search.
// FixedPoint: A[i] = i
func FindFixedPoint(A []int, low, high int) int {
	// If the search space is exhausted (low > high),
	// return -1 indicating that no such index exists.
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if A[mid] == mid {
		return mid
	} else if A[mid] > mid {
		return FindFixedPoint(A, low, mid-1) // Search left subarray
	} else {
		return FindFixedPoint(A, mid+1, high) // Search right subarray
	}
}

// FindHighestDegreeVertex
// Finds the vertex with the highest degree in the adjacency matrix A
func FindHighestDegreeVertex(A [][]int) (int, int) {
	maxDegree := -1 // Initialize the maximum degree to an invalid value
	maxVertex := -1 // Initialize the vertex with the highest degree to an invalid value

	// Iterate over each vertex in the graph
	for i := 0; i < len(A); i++ {
		degree := 0 // Initialize the degree of the current vertex to 0

		// Calculate the degree by summing the edges in the adjacency matrix row
		for j := 0; j < len(A[i]); j++ {
			degree += A[i][j] // A[i][j] is 1 if there's an edge, 0 otherwise
		}

		// Check if the current vertex has a higher degree than the maxDegree found so far
		if degree > maxDegree {
			maxDegree = degree // Update maxDegree with the current vertex's degree
			maxVertex = i      // Update maxVertex to the current vertex index
		}
	}

	return maxVertex, maxDegree // Return the vertex with the highest degree and its degree
}

// FindMaxDistance finds the largest distance between any pair of adjacent points in the array A
func FindMaxDistance(A []int) int {
	maxDistance := 0 // Initialize the maximum distance to 0

	// Iterate over the array to find the largest distance between adjacent points
	for i := 1; i < len(A); i++ {
		distance := A[i] - A[i-1] // Calculate the distance between adjacent points

		// Update maxDistance if the current distance is greater
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance // Return the largest distance found
}

// HasMajorityElement
// Function to check if there is an element that occurs more than N/2 times
func HasMajorityElement(A []int) bool {
	N := len(A)

	// Step 1: Sort the array
	sort.Ints(A) // Sort the array A in increasing order

	// Step 2: Check the middle element
	candidate := A[N/2] // The candidate is the middle element in the sorted array

	// Step 3: Count the occurrences of the candidate element
	count := 0
	for i := 0; i < N; i++ {
		if A[i] == candidate {
			count++
		}
	}

	// Step 4: Check if the count is greater than N/2
	if count > N/2 {
		return true // Majority element exists
	} else {
		return false // No majority element
	}
}
