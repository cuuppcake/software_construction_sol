package bootstrap

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"pos/internal/app"
	"pos/internal/console"
)

func Run(args []string, input io.Reader, output io.Writer) error {
	flags := flag.NewFlagSet("pos", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	flags.Usage = func() {
		fmt.Fprintln(output, "Использование: go run ./cmd/pos [-demo] [-policy weighted|equal] [-out файл.csv]")
		fmt.Fprintln(output, "  -demo    запустить демонстрационный сценарий")
		fmt.Fprintln(output, "  -policy  правило расчёта итога: weighted или equal")
		fmt.Fprintln(output, "  -out     CSV-файл для демонстрации")
	}
	demo := flags.Bool("demo", false, "")
	policy := flags.String("policy", "weighted", "")
	out := flags.String("out", "demo_result.csv", "")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return errors.New("неверные параметры запуска, используйте -h для справки")
	}
	if len(flags.Args()) != 0 {
		return errors.New("лишние аргументы запуска, используйте -h для справки")
	}

	container, err := NewContainer(*policy)
	if err != nil {
		return err
	}
	if *demo {
		return app.RunDemo(container.Service, output, *out)
	}
	return console.New(container.Service, input, output).Run()
}
