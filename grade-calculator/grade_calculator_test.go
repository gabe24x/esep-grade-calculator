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

// -----------------
// Part 4 extra tests
// -----------------

// Boundary thresholds
func TestLetterBoundaries(t *testing.T) {
	cases := []struct {
		assign, exam, essay int
		want                string
	}{
		{90, 90, 90, "A"},
		{80, 80, 80, "B"},
		{70, 70, 70, "C"},
		{60, 60, 60, "D"},
		{59, 59, 59, "F"},
	}
	for _, tc := range cases {
		gc := NewGradeCalculator()
		gc.AddGrade("a", tc.assign, Assignment)
		gc.AddGrade("e", tc.exam, Exam)
		gc.AddGrade("s", tc.essay, Essay)
		if got := gc.GetFinalGrade(); got != tc.want {
			t.Fatalf("(%d,%d,%d) expected %s, got %s", tc.assign, tc.exam, tc.essay, tc.want, got)
		}
	}
}

// Multiple items per category: averaging + weighting
func TestMultipleItemsAverage(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 100, Assignment)
	gc.AddGrade("a2", 90, Assignment) // avg = 95 → 47.5
	gc.AddGrade("x1", 80, Exam)
	gc.AddGrade("x2", 70, Exam)       // avg = 75 → 26.25
	gc.AddGrade("s1", 60, Essay)
	gc.AddGrade("s2", 100, Essay)     // avg = 80 → 12
	// total = 85.75 → floor = 85 = B
	if got := gc.GetFinalGrade(); got != "B" {
		t.Fatalf("expected B, got %s", got)
	}
}

// Empty categories treated as 0, no panic
func TestEmptyCategories(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 100, Assignment)
	gc.AddGrade("a2", 100, Assignment) // only assignments
	// weighted = 50 points only → F
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("expected F with only assignments, got %s", got)
	}
}

// GradeType.String() coverage
func TestGradeTypeStrings(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Assignment string wrong")
	}
	if Exam.String() != "exam" {
		t.Errorf("Exam string wrong")
	}
	if Essay.String() != "essay" {
		t.Errorf("Essay string wrong")
	}
}
