package main

import (
	"math/rand"
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

func TestSolutionRandom(t *testing.T) {
	// Set up a random number generator with a fixed seed
	r := rand.New(rand.NewSource(0))

	for i := 0; i < 1000; i++ {
		// Generate a random value for N
		N := r.Intn(100000) + 1

		// Generate a random slice of skills for T
		T := make([]int, N)
		for j := 0; j < N; j++ {
			T[j] = r.Intn(N)
		}

		// Generate a random slice of skills for A
		M := r.Intn(N) + 1
		A := make([]int, M)
		for j := 0; j < M; j++ {
			A[j] = r.Intn(N)
		}

		// Check if the Solution function is correct
		result := Solution(T, A)
		expected := expectedResult(T, A)
		if result != expected {
			t.Errorf("Solution(%v, %v) = %d, expected %d", T, A, result, expected)
		}
	}
}

// expectedResult calculates the expected result for the given input data
func expectedResult(T, A []int) int {
	// Set up a set to store the skills that have been learned
	learned := map[int]bool{}

	// Set up a variable to store the number of skills learned
	skillsLearned := 0

	// For each skill in array A:
	for _, skill := range A {
		// Traverse the skill tree from the skill to the root
		for !learned[skill] {
			learned[skill] = true
			skill = T[skill]
			skillsLearned++
		}
	}

	// Return the number of skills learned
	return skillsLearned
}
