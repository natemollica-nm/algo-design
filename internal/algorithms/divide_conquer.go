package algorithms

import (
	"math"
	"sort"
)

type Point struct {
	x float64
}

// Helper function to calculate the distance between two points
func distance(p1, p2 Point) float64 {
	return math.Abs(p1.x - p2.x)
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

	// Sort the points by X-coordinate
	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})

	// Divide
	mid := n / 2
	leftPoints := points[:mid]
	rightPoints := points[mid:]

	// Conquer
	d1 := closestPair(leftPoints)
	d2 := closestPair(rightPoints)

	// Combine
	return math.Min(d1, d2)
}
