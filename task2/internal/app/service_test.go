package app_test

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"pos/internal/app"
	"pos/internal/bootstrap"
	"testing"

	"github.com/google/uuid"
)

func testService(t *testing.T) *app.Service {
	t.Helper()
	container, err := bootstrap.NewContainer("weighted")
	if err != nil {
		t.Fatal(err)
	}
	return container.Service
}

func addStudent(t *testing.T, service *app.Service) uuid.UUID {
	t.Helper()
	id, err := service.AddStudent("Анна")
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func addCriterion(t *testing.T, service *app.Service) uuid.UUID {
	t.Helper()
	id, err := service.AddCriterion("Работа", 10, 100)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func TestGradeCorrectionAndCancellation(t *testing.T) {
	service := testService(t)
	studentID := addStudent(t, service)
	criterionID := addCriterion(t, service)
	if err := service.Grade(studentID, criterionID, 8); err != nil {
		t.Fatal(err)
	}

	invalid := []struct {
		name        string
		criterionID uuid.UUID
		points      int
	}{
		{
			name:        "балл выше максимума",
			criterionID: criterionID,
			points:      12,
		},
		{
			name:        "нет критерия",
			criterionID: uuid.Nil,
			points:      3,
		},
	}
	for _, test := range invalid {
		t.Run(test.name, func(t *testing.T) {
			if err := service.Grade(studentID, test.criterionID, test.points); err == nil {
				t.Fatal("недопустимая оценка была принята")
			}
			result, err := service.Result(studentID)
			if err != nil {
				t.Fatal(err)
			}
			if result.Percent() != 80 || result.Completed() != 1 {
				t.Fatal("отклонённая оценка изменила итог")
			}
		})
	}

	if err := service.Grade(studentID, criterionID, 9); err != nil {
		t.Fatal(err)
	}
	result, err := service.Result(studentID)
	if err != nil {
		t.Fatal(err)
	}
	if result.Percent() != 90 {
		t.Fatal("исправление не заменило прежнюю оценку")
	}
	if err := service.CancelGrade(studentID, criterionID); err != nil {
		t.Fatal(err)
	}
	result, err = service.Result(studentID)
	if err != nil {
		t.Fatal(err)
	}
	if result.Percent() != 0 || result.Completed() != 0 {
		t.Fatal("отменённая оценка всё ещё учитывается")
	}
}

func TestCriterionChangesProtectGrades(t *testing.T) {
	service := testService(t)
	studentID := addStudent(t, service)
	criterionID := addCriterion(t, service)
	if err := service.Grade(studentID, criterionID, 8); err != nil {
		t.Fatal(err)
	}
	if err := service.UpdateCriterion(criterionID, "Работа", 5, 100); err == nil {
		t.Fatal("максимум уменьшен ниже уже выставленного балла")
	}
	if err := service.RemoveCriterion(criterionID); err == nil {
		t.Fatal("критерий с оценками был удалён")
	}
	if service.Criteria()[0].MaxScore() != 10 {
		t.Fatal("отклонённое изменение поменяло рубрикатор")
	}
	if err := service.CancelGrade(studentID, criterionID); err != nil {
		t.Fatal(err)
	}
	if err := service.RemoveCriterion(criterionID); err != nil {
		t.Fatal(err)
	}
}

func TestExportResult(t *testing.T) {
	service := testService(t)
	studentID := addStudent(t, service)
	criterionID := addCriterion(t, service)
	if err := service.Grade(studentID, criterionID, 8); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "result.csv")
	if err := service.ExportResult(studentID, path); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	rows, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 || rows[1][0] != studentID.String() || rows[1][1] != "Анна" || rows[1][2] != "80.00" {
		t.Fatalf("неожиданные строки CSV: %v", rows)
	}
}

func TestRemoveStudentDeletesGrades(t *testing.T) {
	service := testService(t)
	studentID := addStudent(t, service)
	criterionID := addCriterion(t, service)
	if err := service.Grade(studentID, criterionID, 8); err != nil {
		t.Fatal(err)
	}
	if err := service.RemoveStudent(studentID); err != nil {
		t.Fatal(err)
	}
	if _, err := service.Result(studentID); err == nil {
		t.Fatal("удалённый студент всё ещё найден")
	}
	if err := service.RemoveCriterion(criterionID); err != nil {
		t.Fatal("оценки удалённого студента остались в хранилище")
	}
}
