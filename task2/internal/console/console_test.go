package console_test

import (
	"bytes"
	"fmt"
	"pos/internal/bootstrap"
	"pos/internal/console"
	"strings"
	"testing"
)

func TestCreateShowsGeneratedIDs(t *testing.T) {
	container, err := bootstrap.NewContainer("weighted")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.NewReader("1\nАнна\n5\nРабота\n10\n50\n0\n")
	var output bytes.Buffer
	if err := console.New(container.Service, input, &output).Run(); err != nil {
		t.Fatal(err)
	}
	students := container.Service.Students()
	criteria := container.Service.Criteria()
	if len(students) != 1 || len(criteria) != 1 {
		t.Fatal("студент или критерий не были созданы")
	}
	if !strings.Contains(output.String(), students[0].ID().String()) || !strings.Contains(output.String(), criteria[0].ID().String()) {
		t.Fatal("сгенерированные ID не показаны в консоли")
	}
}

func TestGradeAndResult(t *testing.T) {
	container, err := bootstrap.NewContainer("weighted")
	if err != nil {
		t.Fatal(err)
	}
	studentID, err := container.Service.AddStudent("Анна")
	if err != nil {
		t.Fatal(err)
	}
	criterionID, err := container.Service.AddCriterion("Работа", 10, 100)
	if err != nil {
		t.Fatal(err)
	}
	input := strings.NewReader(fmt.Sprintf("9\n%s\n%s\n8\n12\n%s\n0\n", studentID, criterionID, studentID))
	var output bytes.Buffer
	if err := console.New(container.Service, input, &output).Run(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "Оценка сохранена.") || !strings.Contains(output.String(), "Итог: 80.00%") {
		t.Fatal("консоль не показала сохранение оценки и итог")
	}
}
