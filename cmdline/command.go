package cmdline

import "fmt"

type Command struct {
	options *Options
	class   string
	args    []string
}

func (self*Command) Class() string {
	return self.class
}

func (self*Command) Args() []string {
	return self.args
}

func (self*Command) Options() *Options {
	return self.options
}

func ParseCommand(osArgs []string) (cmd *Command, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	argReader := &ArgReader{osArgs[1:]}
	cmd = &Command{
		options: parseOption(argReader),
		class:   argReader.removeFirst(),
		args:    argReader.args,
	}
	return
}

func PrintUsage() {
	fmt.Println("usage: gjvm [-option] class [args...]")
}
