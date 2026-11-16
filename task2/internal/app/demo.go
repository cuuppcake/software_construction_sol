package app

import (
	"fmt"
	"io"
)

func RunDemo(service *Service, out io.Writer, csvPath string) error {
	if err := service.AddStudent("s1", "Анна Иванова"); err != nil {
		return err
	}
	if err := service.AddStudent("s2", "Борис Петров"); err != nil {
		return err
	}
	if err := service.AddCriterion("work", "Практические работы", 10, 0.5); err != nil {
		return err
	}
	if err := service.AddCriterion("exam", "Экзамен", 20, 0.3); err != nil {
		return err
	}
	if err := service.AddCriterion("activity", "Активность", 5, 0.2); err != nil {
		return err
	}
	if err := service.Grade("s1", "work", 8); err != nil {
		return err
	}
	if err := service.Grade("s1", "exam", 18); err != nil {
		return err
	}
	if err := service.Grade("s1", "activity", 5); err != nil {
		return err
	}
	if err := service.Grade("s2", "work", 7); err != nil {
		return err
	}

	fmt.Fprintln(out, "Студенты и критерии добавлены, оценки выставлены.")
	if err := service.Grade("s1", "work", 12); err != nil {
		fmt.Fprintln(out, "Оценка 12 из 10 отклонена:", err)
	}
	if err := service.Grade("s1", "unknown", 3); err != nil {
		fmt.Fprintln(out, "Неизвестный критерий отклонён:", err)
	}
	result, err := service.Result("s1")
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Итог Анны: %.2f%% (%d/%d критериев)\n", result.GetPercent(), result.GetCompleted(), result.GetTotal())
	result, err = service.Result("s2")
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Итог Бориса: %.2f%% (%d/%d критериев)\n", result.GetPercent(), result.GetCompleted(), result.GetTotal())
	if err := service.ExportResult("s1", csvPath); err != nil {
		return err
	}
	fmt.Fprintln(out, "Итог Анны экспортирован в", csvPath)
	return nil
}
