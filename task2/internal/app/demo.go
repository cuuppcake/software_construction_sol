package app

import (
	"fmt"
	"io"

	"github.com/google/uuid"
)

func RunDemo(service *Service, out io.Writer, csvPath string) error {
	annaID, err := service.AddStudent("Анна Иванова")
	if err != nil {
		return err
	}
	borisID, err := service.AddStudent("Борис Петров")
	if err != nil {
		return err
	}
	workID, err := service.AddCriterion("Практические работы", 10, 50)
	if err != nil {
		return err
	}
	examID, err := service.AddCriterion("Экзамен", 20, 30)
	if err != nil {
		return err
	}
	activityID, err := service.AddCriterion("Активность", 5, 20)
	if err != nil {
		return err
	}
	if err := service.Grade(annaID, workID, 8); err != nil {
		return err
	}
	if err := service.Grade(annaID, examID, 18); err != nil {
		return err
	}
	if err := service.Grade(annaID, activityID, 5); err != nil {
		return err
	}
	if err := service.Grade(borisID, workID, 7); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(out, "Студенты и критерии добавлены, оценки выставлены."); err != nil {
		return err
	}
	if err := service.Grade(annaID, workID, 12); err != nil {
		if _, writeErr := fmt.Fprintln(out, "Оценка 12 из 10 отклонена:", err); writeErr != nil {
			return writeErr
		}
	}
	if err := service.Grade(annaID, uuid.Nil, 3); err != nil {
		if _, writeErr := fmt.Fprintln(out, "Неизвестный критерий отклонён:", err); writeErr != nil {
			return writeErr
		}
	}
	result, err := service.Result(annaID)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "Итог Анны: %.2f%% (%d/%d критериев)\n", result.Percent(), result.Completed(), result.Total()); err != nil {
		return err
	}
	result, err = service.Result(borisID)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "Итог Бориса: %.2f%% (%d/%d критериев)\n", result.Percent(), result.Completed(), result.Total()); err != nil {
		return err
	}
	if err := service.ExportResult(annaID, csvPath); err != nil {
		return err
	}
	_, err = fmt.Fprintln(out, "Итог Анны экспортирован в", csvPath)
	return err
}
