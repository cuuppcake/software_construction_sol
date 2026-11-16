package domain_test

import (
	"pos/internal/domain"
	"testing"

	"github.com/google/uuid"
)

func TestNewStudent(t *testing.T) {
	tests := []struct {
		scenario string
		name     string
		valid    bool
	}{
		{
			scenario: "пустое имя",
			name:     "",
		},
		{
			scenario: "только пробелы",
			name:     "   ",
		},
		{
			scenario: "пробелы по краям убираются",
			name:     " Анна ",
			valid:    true,
		},
	}
	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			student, err := domain.NewStudent(test.name)
			if (err == nil) != test.valid {
				t.Fatalf("имя %q: неожиданная ошибка %v", test.name, err)
			}
			if test.valid && (student.ID() == uuid.Nil || student.Name() != "Анна") {
				t.Fatal("студент создан с неверными данными")
			}
		})
	}
}

func TestStudentIDsAreDifferent(t *testing.T) {
	first, err := domain.NewStudent("Анна")
	if err != nil {
		t.Fatal(err)
	}
	second, err := domain.NewStudent("Анна")
	if err != nil {
		t.Fatal(err)
	}
	if first.ID() == second.ID() {
		t.Fatal("два студента получили одинаковый ID")
	}
}
