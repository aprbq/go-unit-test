package services_test

import (
	"gotest/services"
	"testing"
)

func TestCheckGradeA(t *testing.T) {

	t.Run("success grade a", func(t *testing.T) {
		grade := services.CheckGrade(80)
		expected := "A"

		if grade != expected {
			t.Errorf("got %v, expected %v", grade, expected)
		}
	})

	t.Run("B", func(t *testing.T) {
		grade := services.CheckGrade(70)
		expected := "B"

		if grade != expected {
			t.Errorf("got %v, expected %v", grade, expected)
		}
	})

	t.Run("C", func(t *testing.T) {
		grade := services.CheckGrade(60)
		expected := "C"

		if grade != expected {
			t.Errorf("got %v, expected %v", grade, expected)
		}
	})
}
