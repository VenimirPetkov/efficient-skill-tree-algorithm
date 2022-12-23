package main

import (
	"testing"
)

func TestSolutionCase1(t *testing.T) {
	// Test input
	T := []int{0, 0, 1, 1}
	A := []int{2}

	// Expected output
	expected := 3

	// Test the Solution function
	result := Solution(T, A)
	if result != expected {
		t.Errorf("Solution(%v, %v) = %v, expected %v", T, A, result, expected)
	}
}
func TestSolutionCase2(t *testing.T) {
	// Test input
	T := []int{0, 0, 0, 0, 2, 3, 3}
	A := []int{2, 5, 6}

	// Expected output
	expected := 5

	// Test the Solution function
	result := Solution(T, A)
	if result != expected {
		t.Errorf("Solution(%v, %v) = %v, expected %v", T, A, result, expected)
	}
}
func TestSolutionCase3(t *testing.T) {
	// Test input
	T := []int{0, 0, 1, 2}
	A := []int{1, 2}

	// Expected output
	expected := 3

	// Test the Solution function
	result := Solution(T, A)
	if result != expected {
		t.Errorf("Solution(%v, %v) = %v, expected %v", T, A, result, expected)
	}
}

func TestSolutionCase4(t *testing.T) {
	// Test input
	T := []int{0, 3, 0, 0, 5, 0, 5}
	A := []int{4, 2, 6, 1, 0}

	// Expected output
	expected := 7

	// Test the Solution function
	result := Solution(T, A)
	if result != expected {
		t.Errorf("Solution(%v, %v) = %v, expected %v", T, A, result, expected)
	}
}

func TestSolution(t *testing.T) {
	// Test cases
	testCases := []struct {
		T        []int
		A        []int
		expected int
	}{
		{
			T:        []int{0, 0, 1, 1},
			A:        []int{2},
			expected: 3,
		},
		{
			T:        []int{0, 0, 0, 0, 2, 3, 3},
			A:        []int{2, 5, 6},
			expected: 5,
		},
		{
			T:        []int{0, 0, 1, 2},
			A:        []int{1, 2},
			expected: 3,
		},
		{
			T:        []int{0, 3, 0, 0, 5, 0, 5},
			A:        []int{4, 2, 6, 1, 0},
			expected: 7,
		},
	}

	// Test the Solution function for each test case
	for _, tc := range testCases {
		result := Solution(tc.T, tc.A)
		if result != tc.expected {
			t.Errorf("Solution(%v, %v) = %v, expected %v", tc.T, tc.A, result, tc.expected)
		}
	}
}
