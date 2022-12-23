package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindSkillsToLearnCase1(t *testing.T) {
	T := []int{0, 0, 1, 1}
	A := []int{2}
	M := 4
	skillID := 2
	expected := []int{2, 1, 0}
	actual, err := FindSkillsToLearn(T, A, M, skillID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(expected) != len(actual) {
		t.Errorf("unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
			return
		}
	}
}

func TestFindSkillsToLearnCase2(t *testing.T) {
	T := []int{0, 0, 0, 0, 2, 3, 3}
	A := []int{2, 5, 6}
	M := 7
	skillID := 2

	expected := []int{2, 0}
	actual, err := FindSkillsToLearn(T, A, M, skillID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
	}

	skillID = 5
	expected = []int{5, 3, 0}
	actual, err = FindSkillsToLearn(T, A, M, skillID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
	}

	skillID = 6
	expected = []int{6, 3, 0}
	actual, err = FindSkillsToLearn(T, A, M, skillID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
	}

	skillID = 10
	expected = nil
	_, err = FindSkillsToLearn(T, A, M, skillID)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
func TestFindSkillsToLearnCase2Optimized(t *testing.T) {
	type testCase struct {
		T        []int
		A        []int
		M        int
		skillID  int
		expected []int
		err      error
	}

	testCases := []testCase{
		{
			T:        []int{0, 0, 0, 0, 2, 3, 3},
			A:        []int{2, 5, 6},
			M:        7,
			skillID:  2,
			expected: []int{2, 0},
			err:      nil,
		},
		{
			T:        []int{0, 0, 0, 0, 2, 3, 3},
			A:        []int{2, 5, 6},
			M:        7,
			skillID:  5,
			expected: []int{5, 3, 0},
			err:      nil,
		},
		{
			T:        []int{0, 0, 0, 0, 2, 3, 3},
			A:        []int{2, 5, 6},
			M:        7,
			skillID:  6,
			expected: []int{6, 3, 0},
			err:      nil,
		},
		{
			T:        []int{0, 0, 0, 0, 2, 3, 3},
			A:        []int{2, 5, 6},
			M:        7,
			skillID:  10,
			expected: nil,
			err:      fmt.Errorf("skill ID %d not found", 10),
		},
	}

	for _, tc := range testCases {
		skills, err := FindSkillsToLearn(tc.T, tc.A, tc.M, tc.skillID)
		if len(skills) != len(tc.expected) {
			t.Errorf("Unexpected result for skill %d. Expected %v, got %v", tc.skillID, tc.expected, skills)
		} else {
			for i := range skills {
				if skills[i] != tc.expected[i] {
					t.Errorf("Unexpected result for skill %d. Expected %v, got %v", tc.skillID, tc.expected, skills)
					break
				}
			}
		}
		if err == nil && tc.err != nil {
			t.Errorf("Expected error %v, got nil", tc.err)
		} else if err != nil && tc.err == nil {
			t.Errorf("Unexpected error: %v", err)
		} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Expected error %v, got %v", tc.err, err)
		}
	}
}
func TestFindSkillsToLearnCase3(t *testing.T) {
	T := []int{0, 0, 1, 2}
	A := []int{1, 2}
	M := 4
	skillID := 2
	expected := []int{2, 1, 0}
	actual, err := FindSkillsToLearn(T, A, M, skillID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(expected) != len(actual) {
		t.Errorf("unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
			return
		}
	}
}
func TestFindSkillsToLearnCase4(t *testing.T) {
	T := []int{0, 3, 0, 0, 5, 0, 1}
	A := []int{4, 2, 6, 1, 0}
	M := 7
	skillID := 4
	expected := []int{4, 5, 0, 3}
	actual, err := FindSkillsToLearn(T, A, M, skillID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(expected) != len(actual) {
		t.Errorf("unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("unexpected result for skill %d. Expected %v, got %v", skillID, expected, actual)
			return
		}
	}
}
