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

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 79, Assignment)
	gradeCalculator.AddGrade("exam 1", 75, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 69, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 59, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestStringAssignment(t *testing.T) {
	expected_value := "assignment"

	actual_value := Assignment.String()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestExamAssignment(t *testing.T) {
	expected_value := "exam"

	actual_value := Exam.String()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestEssayAssignment(t *testing.T) {
	expected_value := "essay"

	actual_value := Essay.String()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestPassA(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 90, Essay)

	actual_value := gradeCalculator.GetFinalPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestPassB(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 89, Assignment)
	gradeCalculator.AddGrade("exam 1", 85, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 80, Essay)

	actual_value := gradeCalculator.GetFinalPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestPassC(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 79, Assignment)
	gradeCalculator.AddGrade("exam 1", 75, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestPassD(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 67, Assignment)
	gradeCalculator.AddGrade("exam 1", 64, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestPassF(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}