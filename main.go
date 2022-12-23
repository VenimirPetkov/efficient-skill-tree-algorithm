package main

import "fmt"

// FindSkillsToLearn returns the skills that need to be learned in order to learn the skill with the given ID.
// The skills are returned in the order that they need to be learned, with the root skill (ID 0) first.
// The input skill tree is represented as two arrays, T and A, where T[i] is the ID of the parent skill of skill i,
// and A[i] is the ID of the child skill of skill i. The input M is the number of skills in the skill tree.
func FindSkillsToLearn(T []int, A []int, M int, skillID int) ([]int, error) {
	// Define a recursive function to find the skills that need to be learned
	var findSkillsToLearnRecursive func(int) []int
	findSkillsToLearnRecursive = func(id int) []int {
		skills := []int{id}
		if id == 0 {
			// Stop the recursion when reaching the root skill
			return skills
		}
		if id < 0 || id >= M {
			// Return an empty list if the skill ID is out of range
			return []int{}
		}
		// Add the skills that need to be learned to learn the parent skill
		parentID := T[id]
		skills = append(skills, findSkillsToLearnRecursive(parentID)...)
		return skills
	}

	if skillID < 0 || skillID >= M {
		// Return an error if the skill ID is out of range
		return nil, fmt.Errorf("skill ID %d not found", skillID)
	}
	return findSkillsToLearnRecursive(skillID), nil
}

func FindSkillsToLearnSimple(T []int, A []int, M int, skillID int) ([]int, error) {
	skills := []int{}

	// Loop through the parent skills until reaching the root skill (ID 0)
	for id := skillID; id != 0; id = T[id] {
		if id < 0 || id >= M {
			// Return an empty list if the skill ID is out of range
			return []int{}, fmt.Errorf("skill ID %d not found", id)
		}
		skills = append(skills, id)
	}
	skills = append(skills, 0) // Add the root skill (ID 0) to the list

	return skills, nil
}
