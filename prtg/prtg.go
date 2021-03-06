package prtg

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var output io.Writer = os.Stdout
var exit func(int) = os.Exit

func Out(_ ...interface{}) {
	fmt.Fprintln(output, "0:0:OK")
	exit(0)
}

func Err(args ...interface{}) {
	s, str := make([]string, 0, len(args)), ""
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			str = arg
		case fmt.Stringer:
			str = arg.String()
		case error:
			str = arg.Error()
		default:
			str = fmt.Sprintf("%v", arg)
		}
		s = append(s, strconv.Quote(str))
	}
	fmt.Fprintf(output, "2:1:%s\n", strings.Join(s, " "))
	exit(1)
}
