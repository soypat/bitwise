package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/cosmos72/gomacro/fast"
	"github.com/spf13/pflag"
)

func macro(toeval string) reflect.Value {
	interp := fast.New()
	for _, v := range imports {
		interp.Eval(fmt.Sprintf("import . \"%s\";", v))
	}

	vals, _ := interp.Eval(toeval)
	// for simplicity, only use the first returned value
	return vals[0].ReflectValue()
}

var (
	format  = "v"
	imports = []string{"math"}
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("bw:", err)
			os.Exit(1)
		}
	}()
	pflag.StringVarP(&format, "format", "f", format, "Result format. Golang format rules. Omit the leading % sign")
	pflag.StringArrayVarP(&imports, "imports", "", imports, "Standard library imports to include.")

	pflag.Parse()
	format = "%" + format
	args := pflag.Args()
	if len(args) == 0 {
		pflag.Usage()
		panic("no expression found")
	}
	expr := strings.Join(pflag.Args(), " ")

	fmt.Print(expr + " = ")
	fmt.Printf(format+"\n", macro(expr))
}
