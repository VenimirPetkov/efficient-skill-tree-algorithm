Skill Tree
The Solution function takes two slices of integers as input:

T: a slice of N integers representing the skill tree. The K-th skill can be learned only if the T[K]-th skill has already been learned. Skill 0 is the root and T[0] = 0.
A: a slice of M integers representing the skills to be acquired.
The function returns the minimum number of skills which have to be learned to acquire all of the skills from the array A in the skill tree T.

Example
T := []int{0, 0, 1, 1}
A := []int{2}
result := Solution(T, A)
fmt.Println(result) // Output: 3

In this example, the minimum number of skills to be learned to acquire skill 2 is 3, as explained in the prompt. The function returns 3 as the output.

Test Cases
Here are the test cases that verify the correctness of the Solution function:

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