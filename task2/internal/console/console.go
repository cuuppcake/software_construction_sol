package console

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"pos/internal/app"
	"strconv"
	"strings"
)

type Console struct {
	service *app.Service
	input   *bufio.Scanner
	out     io.Writer
}

func New(service *app.Service, input io.Reader, out io.Writer) *Console {
	return &Console{service: service, input: bufio.NewScanner(input), out: out}
}

func (c *Console) prompt(label string) (string, error) {
	fmt.Fprint(c.out, label, ": ")
	if !c.input.Scan() {
		if c.input.Err() != nil {
			return "", errors.New("не удалось прочитать ввод")
		}
		return "", io.EOF
	}
	return strings.TrimSpace(c.input.Text()), nil
}

func (c *Console) integer(label string) (int, error) {
	value, err := c.prompt(label)
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("нужно ввести целое число, получено %q", value)
	}
	return result, nil
}

func (c *Console) decimal(label string) (float64, error) {
	value, err := c.prompt(label)
	if err != nil {
		return 0, err
	}
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("нужно ввести число, получено %q", value)
	}
	return result, nil
}

func (c *Console) Run() error {
	for {
		fmt.Fprintln(c.out, "\nПОС — преподавательская оценка студентов")
		fmt.Fprintln(c.out, "1. Добавить студента   2. Список студентов   3. Переименовать студента   4. Удалить студента")
		fmt.Fprintln(c.out, "5. Добавить критерий  6. Список критериев  7. Изменить критерий       8. Удалить критерий")
		fmt.Fprintln(c.out, "9. Выставить/исправить оценку  10. Отменить оценку  11. Оценки студента")
		fmt.Fprintln(c.out, "12. Итог студента  13. Экспорт итога в CSV  0. Выход")

		choice, err := c.prompt("Выбор")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if choice == "0" {
			return nil
		}
		if err := c.handle(choice); err != nil {
			if err == io.EOF {
				return nil
			}
			fmt.Fprintln(c.out, "Ошибка:", err)
		} else if message := successMessage(choice); message != "" {
			fmt.Fprintln(c.out, message)
		}
	}
}

func (c *Console) handle(choice string) error {
	switch choice {
	case "1":
		return c.addStudent()
	case "2":
		c.showStudents()
	case "3":
		return c.renameStudent()
	case "4":
		return c.removeStudent()
	case "5":
		return c.saveCriterion(false)
	case "6":
		c.showCriteria()
	case "7":
		return c.saveCriterion(true)
	case "8":
		return c.removeCriterion()
	case "9":
		return c.grade()
	case "10":
		return c.cancelGrade()
	case "11":
		return c.showEvaluations()
	case "12":
		return c.showResult()
	case "13":
		return c.exportResult()
	default:
		return fmt.Errorf("неизвестная команда %q", choice)
	}
	return nil
}

func successMessage(choice string) string {
	switch choice {
	case "1":
		return "Студент добавлен."
	case "3":
		return "Имя студента изменено."
	case "4":
		return "Студент удалён вместе с его оценками."
	case "5":
		return "Критерий добавлен."
	case "7":
		return "Критерий изменён."
	case "8":
		return "Критерий удалён."
	case "9":
		return "Оценка сохранена."
	case "10":
		return "Оценка отменена."
	}
	return ""
}
