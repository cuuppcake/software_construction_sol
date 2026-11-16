package bootstrap

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseOptions(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantDemo   bool
		wantPolicy string
		wantOut    string
		wantHelp   bool
		wantError  bool
	}{
		{
			name:       "обычный запуск",
			wantPolicy: "weighted",
			wantOut:    "demo_result.csv",
		},
		{
			name:       "демо с другим правилом и файлом",
			args:       []string{"-demo", "-policy", "equal", "-out", "result.csv"},
			wantDemo:   true,
			wantPolicy: "equal",
			wantOut:    "result.csv",
		},
		{
			name:     "справка",
			args:     []string{"-h"},
			wantHelp: true,
		},
		{
			name:      "неизвестный параметр",
			args:      []string{"-unknown"},
			wantError: true,
		},
		{
			name:      "лишний аргумент",
			args:      []string{"extra"},
			wantError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			got, err := parseOptions(test.args, &output)
			if (err != nil) != test.wantError {
				t.Fatalf("ошибка: %v", err)
			}
			if test.wantError {
				return
			}
			if got.demo != test.wantDemo || got.policy != test.wantPolicy || got.out != test.wantOut || got.help != test.wantHelp {
				t.Fatalf("неверные параметры: %+v", got)
			}
			if test.wantHelp && !strings.Contains(output.String(), "Как запустить") {
				t.Fatal("справка не показана")
			}
		})
	}
}
