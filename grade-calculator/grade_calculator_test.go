package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 45, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment 1", 70, Assignment)
	gradeCalculator.AddGrade("assignment 2", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 72, Exam)
	gradeCalculator.AddGrade("essay 1", 68, Essay)
	gradeCalculator.AddGrade("essay 2", 72, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment 1", 60, Assignment)
	gradeCalculator.AddGrade("assignment 2", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 62, Exam)
	gradeCalculator.AddGrade("essay on philosophy", 58, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeBoundaryA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	// Test exactly 90 (boundary for A)
	gradeCalculator.AddGrade("assignment 1", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("essay 1", 90, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeBoundaryB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	// Test exactly 80 (boundary for B)
	gradeCalculator.AddGrade("assignment 1", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("essay 1", 80, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeBoundaryC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	// Test exactly 70 (boundary for C)
	gradeCalculator.AddGrade("assignment 1", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay 1", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeBoundaryD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	// Test exactly 60 (boundary for D)
	gradeCalculator.AddGrade("assignment 1", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay 1", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	// Test String() method for GradeType enum
	if Assignment.String() != "assignment" {
		t.Errorf("Expected Assignment.String() to return 'assignment'; got '%s' instead", Assignment.String())
	}

	if Exam.String() != "exam" {
		t.Errorf("Expected Exam.String() to return 'exam'; got '%s' instead", Exam.String())
	}

	if Essay.String() != "essay" {
		t.Errorf("Expected Essay.String() to return 'essay'; got '%s' instead", Essay.String())
	}
}

func TestEmptyGradeList(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	// Don't add any grades - all averages should be 0, resulting in F
	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
