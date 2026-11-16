package export

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"pos/internal/domain"
	"strconv"
)

type CSVExporter struct{}

func (e CSVExporter) Export(student domain.Student, result domain.Result, path string) error {
	rows := [][]string{
		{"student_id", "student_name", "percent", "completed", "total"},
		{student.ID().String(), student.Name(), strconv.FormatFloat(result.Percent(), 'f', 2, 64), strconv.Itoa(result.Completed()), strconv.Itoa(result.Total())},
	}
	var data bytes.Buffer
	if err := csv.NewWriter(&data).WriteAll(rows); err != nil {
		return fmt.Errorf("не удалось записать CSV: %w", err)
	}
	if err := os.WriteFile(path, data.Bytes(), 0644); err != nil {
		return fmt.Errorf("не удалось создать CSV-файл %q: %w", path, err)
	}
	return nil
}
