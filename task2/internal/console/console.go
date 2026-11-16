package console

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"pos/internal/app"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Console struct {
	service *app.Service
	input   *bufio.Scanner
	out     io.Writer
}

const menu = `
ПОС — преподавательская оценка студентов
1. Добавить студента   2. Список студентов   3. Переименовать студента   4. Удалить студента
5. Добавить критерий  6. Список критериев  7. Изменить критерий       8. Удалить критерий
9. Выставить/исправить оценку  10. Отменить оценку  11. Оценки студента
12. Итог студента  13. Экспорт итога в CSV  0. Выход
`

func New(service *app.Service, input io.Reader, out io.Writer) *Console {
	return &Console{
		service: service,
		input:   bufio.NewScanner(input),
		out:     out,
	}
}

func (c *Console) println(args ...any) error {
	_, err := fmt.Fprintln(c.out, args...)
	return err
}

func (c *Console) printf(format string, args ...any) error {
	_, err := fmt.Fprintf(c.out, format, args...)
	return err
}

func (c *Console) prompt(label string) (string, error) {
	if _, err := fmt.Fprint(c.out, label, ": "); err != nil {
		return "", fmt.Errorf("не удалось вывести приглашение: %w", err)
	}
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

func (c *Console) id(label string) (uuid.UUID, error) {
	value, err := c.prompt(label)
	if err != nil {
		return uuid.Nil, err
	}
	result, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("некорректный ID %q", value)
	}
	return result, nil
}

func (c *Console) Run() error {
	for {
		if _, err := fmt.Fprint(c.out, menu); err != nil {
			return err
		}

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
			if writeErr := c.println("Ошибка:", err); writeErr != nil {
				return writeErr
			}
		} else if message := successMessage(choice); message != "" {
			if err := c.println(message); err != nil {
				return err
			}
		}
	}
}

func (c *Console) handle(choice string) error {
	switch choice {
	case "1":
		return c.addStudent()
	case "2":
		return c.showStudents()
	case "3":
		return c.renameStudent()
	case "4":
		return c.removeStudent()
	case "5":
		return c.saveCriterion(false)
	case "6":
		return c.showCriteria()
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
