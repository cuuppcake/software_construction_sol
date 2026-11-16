package bootstrap

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

type options struct {
	demo   bool
	policy string
	out    string
	help   bool
}

func parseOptions(args []string, output io.Writer) (options, error) {
	flags := flag.NewFlagSet("pos", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	var usageErr error
	flags.Usage = func() {
		_, usageErr = fmt.Fprint(output, `Как запустить:
  go run ./cmd/pos         открыть меню
  go run ./cmd/pos -demo   посмотреть готовый пример

Дополнительные параметры:
  -policy weighted|equal  считать по весам или поровну (по умолчанию weighted)
  -out файл.csv           куда сохранить итог при -demo
`)
	}
	demo := flags.Bool("demo", false, "")
	policy := flags.String("policy", "weighted", "")
	out := flags.String("out", "demo_result.csv", "")
	if err := flags.Parse(args); err != nil {
		if usageErr != nil {
			return options{}, fmt.Errorf("не удалось вывести справку: %w", usageErr)
		}
		if errors.Is(err, flag.ErrHelp) {
			return options{help: true}, nil
		}
		return options{}, errors.New("неверные параметры запуска, используйте -h для справки")
	}
	if len(flags.Args()) != 0 {
		return options{}, errors.New("лишние аргументы запуска, используйте -h для справки")
	}
	return options{
		demo:   *demo,
		policy: *policy,
		out:    *out,
	}, nil
}
