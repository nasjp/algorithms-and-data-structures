package chapter1

import (
	"math"
)

func Practice1_1(x int) bool {
	return x%2 == 0
}

func Practice1_2(x int) bool {
	return x%5 == 0
}

func Practice1_3(x []int) int {
	if len(x) == 0 {
		return 0
	}

	max := x[0]

	for _, v := range x {
		if max < v {
			max = v
		}
	}

	return max
}

func Practice1_4Mean(x []int) float64 {
	if len(x) == 0 {
		return 0
	}

	var sum float64

	for _, xi := range x {
		sum += float64(xi)
	}

	return sum / float64(len(x))
}

func Practice1_4Variance(x []int) float64 {
	if len(x) == 0 {
		return 0
	}

	_x := Practice1_4Mean(x)

	var diffSum float64

	for _, xi := range x {
		diff := (float64(xi) - _x)
		diffSum += diff * diff
	}

	return diffSum / float64(len(x))
}

type Point struct {
	X, Y float64
}

func Practice1_5(a, b *Point) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func Practice1_6(a, b, c *Point) float64 {
	return math.Abs((a.X-c.X)*(b.Y-c.Y)-(b.X-c.X)*(a.Y-c.Y)) / 2
}

func Practice1_7(x int) bool {
	if x <= 0 {
		return false
	}

	for i := 1; i*i <= x; i++ {
		if i*i == x {
			return true
		}
	}

	return false
}

func Practice1_8(a, b int) int {
	if a < b {
		b, a = a, b
	}

	prod := a * b

	for r := a % b; r > 0; r = a % b {
		a = b
		b = r
	}

	return prod / b
}

func Practice1_10(x, y int) (int, int) {
	const (
		craneFeet  = 2
		turtleFeet = 4
	)

	for t := 0; t <= y; t++ {
		if c := y - t; c*craneFeet+t*turtleFeet == x {
			return c, t
		}
	}

	return 0, 0
}

func Practice1_13(x ...int) int {
	if len(x) == 0 {
		return 0
	}

	max := x[0]

	for _, v := range x {
		if max < v {
			max = v
		}
	}

	return max
}

func Practice1_14(x int) []int {
	var factors []int
	for x%2 == 0 {
		factors = append(factors, 2)
		x /= 2
	}

	for i := 3; i <= x; i += 2 {
		if x%i != 0 {
			continue
		}

		factors = append(factors, i)
		x /= i
	}

	return factors
}

func Practice1_15(n uint32) uint32 {
	c := uint32(1)

	for i := uint32(1); i <= n-1; i++ {
		c = c * (n + i - 1) / i
	}

	return c / (n)
}
