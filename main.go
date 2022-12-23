package main

func Solution(T []int, A []int) int {
	// Set up a set to store the skills that have been visited
	visited := map[int]bool{}

	// Set up a variable to store the minimum number of skills learned
	minSkillsLearned := 0

	// For each skill in array A:
	for _, skill := range A {
		// Traverse the skill tree from the skill to the root
		for !visited[skill] {
			visited[skill] = true
			skill = T[skill]
			minSkillsLearned++
		}
	}

	// Return the minimum number of skills learned
	return minSkillsLearned
}
