package export

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"pos/internal/domain"
	"strconv"
)

type CSVExporter struct{}

func (e CSVExporter) Export(student domain.Student, result domain.Result, path string) error {
	file, err := os.Create(path)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return fmt.Errorf("не удалось создать CSV-файл %q: каталог не найден", path)
		case errors.Is(err, os.ErrPermission):
			return fmt.Errorf("не удалось создать CSV-файл %q: нет прав доступа", path)
		default:
			return fmt.Errorf("не удалось создать CSV-файл %q", path)
		}
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	if err := writer.Write([]string{"student_id", "student_name", "percent", "completed", "total"}); err != nil {
		return errors.New("не удалось записать заголовок CSV-файла")
	}
	if err := writer.Write([]string{student.GetID(), student.GetName(), strconv.FormatFloat(result.GetPercent(), 'f', 2, 64), strconv.Itoa(result.GetCompleted()), strconv.Itoa(result.GetTotal())}); err != nil {
		return errors.New("не удалось записать итог в CSV-файл")
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return errors.New("не удалось сохранить CSV-файл")
	}
	return nil
}
