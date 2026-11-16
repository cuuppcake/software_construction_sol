package bootstrap

import (
	"io"
	"pos/internal/app"
	"pos/internal/console"
)

func Run(args []string, input io.Reader, output io.Writer) error {
	opts, err := parseOptions(args, output)
	if err != nil || opts.help {
		return err
	}

	container, err := NewContainer(opts.policy)
	if err != nil {
		return err
	}
	if opts.demo {
		return app.RunDemo(container.Service, output, opts.out)
	}
	return console.New(container.Service, input, output).Run()
}
