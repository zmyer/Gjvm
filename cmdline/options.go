package cmdline

import (
	"strconv"
	"strings"
)

const (
	_1K = 1024
	_1M = _1K * _1K
	_1G = _1M * _1M
)

type Options struct {
	classpath    string
	verboseClass bool
	Xss          int
	Xcpuprofile  string
	XuseJavaHome bool
}

func (self*Options) Classpath() string {
	return self.classpath
}

func (self*Options) VerboseClass() bool {
	return self.verboseClass
}

func parseOption(argReader *ArgReader) *Options {
	options := &Options{
		Xss: 16 * _1K,
	}
	for argReader.hasMoreOptions() {
		optionName := argReader.removeFirst()
		switch optionName {
		case "-cp", "-classpath":
			options.classpath = argReader.removeFirst();
		case "-verbose", "verbose:class":
			options.verboseClass = true
		case "-Xcpuprofile":
			options.Xcpuprofile = argReader.removeFirst()
		case "-XuseJavaHome":
			options.XuseJavaHome = true
		default:
			if strings.HasPrefix(optionName, "-Xss") {
				options.Xss = parseXss(optionName)
			} else {
				panic("Unrecoginied option:" + optionName)
			}
		}
	}
	return options
}

func parseXss(optionName string) int {
	size := optionName[4:]
	switch size[len(size) - 1] {
	case 'g', 'G':
		return _1G * parseInt(size[:len(size) - 1])
	case 'm', 'M':
		return _1M * parseInt(size[:len(size) - 1])
	case 'k', 'K':
		return _1K * parseInt(size[:len(size) - 1])
	default:
		return parseInt(size)
	}

}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err.Error())
	}
	return i

}
