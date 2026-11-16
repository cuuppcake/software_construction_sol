package export_test

import (
	"errors"
	"os"
	"path/filepath"
	"pos/internal/domain"
	"pos/internal/export"
	"testing"
)

func TestExportKeepsFileError(t *testing.T) {
	student, err := domain.NewStudent("Анна")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "missing", "result.csv")
	err = (export.CSVExporter{}).Export(student, domain.NewResult(80, 1, 1), path)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("ожидалась причина ошибки об отсутствующем каталоге, получено %v", err)
	}
}
