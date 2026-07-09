package services_test

import (
	"gotest/services"
	"testing"
)

func TestCheckGradeA(t *testing.T) {
	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "a", score: 80, expected: "A"},
		{name: "b", score: 70, expected: "B"},
		{name: "c", score: 60, expected: "C"},
		{name: "d", score: 50, expected: "D"},
		{name: "f", score: 0, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)

			if grade != c.expected {
				t.Errorf("got %v, expected %v", grade, c.expected)
			}
		})
	}

}
