package options

import (
	"os"
	"path/filepath"
	"Gjvm/cmdline"
)

var (
	VerboseClass    bool
	ThreadStackSize uint
	AbsJavaHome     string
	AbsJreLib       string
)

func InitOptions(cmdOption *cmdline.Options) {
	VerboseClass = cmdOption.VerboseClass()
	ThreadStackSize = uint(cmdOption.Xss)
	initJavaHome(cmdOption.XuseJavaHome)
}
func initJavaHome(useOsEnv bool) {
	jh := "./jre"
	if useOsEnv {
		jh = os.Getenv("JAVA_HOME")
		if jh == "" {
			panic("$JAVA_HOME not set")
		}
	}

	if absJh, err := filepath.Abs(jh); err == nil {
		AbsJavaHome = absJh
		AbsJreLib = filepath.Join(absJh, "lib")
	} else {
		panic(err)
	}
}
