package console_test

import (
	"bytes"
	"pos/internal/bootstrap"
	"pos/internal/console"
	"strings"
	"testing"
)

func TestMainScenario(t *testing.T) {
	container, err := bootstrap.NewContainer("weighted")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.NewReader(strings.Join([]string{
		"2", "6",
		"1", "s1", "Анна",
		"5", "work", "Работа", "10", "1",
		"9", "s1", "work", "8",
		"11", "s1",
		"12", "s1",
		"1", "s1", "Анна",
		"0",
	}, "\n") + "\n")
	var output bytes.Buffer
	if err := console.New(container.Service, input, &output).Run(); err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{
		"Студентов пока нет.", "Критериев пока нет.",
		"Студент добавлен.", "Критерий добавлен.", "Оценка сохранена.",
		"work — 8", "Итог: 80.00%", "студент \"s1\" уже существует",
	} {
		if !strings.Contains(output.String(), expected) {
			t.Fatalf("в выводе нет строки %q", expected)
		}
	}
}
