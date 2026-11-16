package app_test

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"pos/internal/app"
	"pos/internal/bootstrap"
	"testing"
)

func testService(t *testing.T) *app.Service {
	t.Helper()
	container, err := bootstrap.NewContainer("weighted")
	if err != nil {
		t.Fatal(err)
	}
	return container.Service
}

func TestGradeCorrectionAndCancellation(t *testing.T) {
	service := testService(t)
	if err := service.AddStudent("s1", "Anna"); err != nil {
		t.Fatal(err)
	}
	if err := service.AddCriterion("work", "Work", 10, 1); err != nil {
		t.Fatal(err)
	}
	if err := service.Grade("s1", "work", 8); err != nil {
		t.Fatal(err)
	}
	if err := service.Grade("s1", "work", 12); err == nil {
		t.Fatal("недопустимое исправление оценки было принято")
	}
	result, err := service.Result("s1")
	if err != nil || result.GetPercent() != 80 {
		t.Fatal("отклонённое исправление изменило итог")
	}
	if err := service.Grade("s1", "work", 9); err != nil {
		t.Fatal(err)
	}
	result, err = service.Result("s1")
	if err != nil {
		t.Fatal(err)
	}
	if result.GetPercent() != 90 {
		t.Fatal("исправление не заменило прежнюю оценку")
	}
	if err := service.CancelGrade("s1", "work"); err != nil {
		t.Fatal(err)
	}
	result, err = service.Result("s1")
	if err != nil {
		t.Fatal(err)
	}
	if result.GetPercent() != 0 || result.GetCompleted() != 0 {
		t.Fatal("отменённая оценка всё ещё учитывается")
	}
}

func TestCriterionChangesProtectExistingGrades(t *testing.T) {
	service := testService(t)
	if err := service.AddStudent("s1", "Anna"); err != nil {
		t.Fatal(err)
	}
	if err := service.AddCriterion("work", "Work", 10, 1); err != nil {
		t.Fatal(err)
	}
	if err := service.Grade("s1", "work", 8); err != nil {
		t.Fatal(err)
	}
	if err := service.UpdateCriterion("work", "Work", 5, 1); err == nil {
		t.Fatal("максимум уменьшен ниже уже выставленной оценки")
	}
	if err := service.RemoveCriterion("work"); err == nil {
		t.Fatal("критерий с оценками был удалён")
	}
	if service.Criteria()[0].GetMaxScore() != 10 {
		t.Fatal("отклонённое изменение поменяло рубрикатор")
	}
	if err := service.CancelGrade("s1", "work"); err != nil {
		t.Fatal(err)
	}
	if err := service.RemoveCriterion("work"); err != nil {
		t.Fatal(err)
	}
	if len(service.Criteria()) != 0 {
		t.Fatal("критерий не был удалён")
	}
}

func TestExportResult(t *testing.T) {
	service := testService(t)
	if err := service.AddStudent("s1", "Анна"); err != nil {
		t.Fatal(err)
	}
	if err := service.AddCriterion("work", "Работа", 10, 1); err != nil {
		t.Fatal(err)
	}
	if err := service.Grade("s1", "work", 8); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "result.csv")
	if err := service.ExportResult("s1", path); err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 || rows[1][0] != "s1" || rows[1][1] != "Анна" || rows[1][2] != "80.00" {
		t.Fatalf("неожиданные строки CSV: %v", rows)
	}
}

func TestRemoveStudentDeletesGrades(t *testing.T) {
	service := testService(t)
	if err := service.AddStudent("s1", "Анна"); err != nil {
		t.Fatal(err)
	}
	if err := service.AddCriterion("work", "Работа", 10, 1); err != nil {
		t.Fatal(err)
	}
	if err := service.Grade("s1", "work", 8); err != nil {
		t.Fatal(err)
	}
	if err := service.RemoveStudent("s1"); err != nil {
		t.Fatal(err)
	}
	if err := service.AddStudent("s1", "Анна"); err != nil {
		t.Fatal(err)
	}
	result, err := service.Result("s1")
	if err != nil {
		t.Fatal(err)
	}
	if result.GetCompleted() != 0 || result.GetPercent() != 0 {
		t.Fatal("после повторного добавления студента сохранились прежние оценки")
	}
}
