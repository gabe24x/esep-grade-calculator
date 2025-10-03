package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected := "A"
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 100, Assignment)
	gc.AddGrade("exam", 100, Exam)
	gc.AddGrade("essay", 100, Essay)
	if got := gc.GetFinalGrade(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestGetGradeB(t *testing.T) {
	expected := "B"
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 80, Assignment)
	gc.AddGrade("exam", 81, Exam)
	gc.AddGrade("essay", 85, Essay)
	if got := gc.GetFinalGrade(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestGetGradeHighScores(t *testing.T) {
	expected := "A" // FIXED: should be A, not F
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 100, Assignment)
	gc.AddGrade("exam", 95, Exam)
	gc.AddGrade("essay", 91, Essay)
	if got := gc.GetFinalGrade(); got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
