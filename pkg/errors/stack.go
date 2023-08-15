package errors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

func (e *Wrapper) GetErrorLine() string {
	var fun, file, stack string
	var line int

	if e.stack != nil {
		fun, file, line = e.getFuncFileLine()
		stack = fmt.Sprintf("[%s]%s:%d", fun, file, line)
	}
	return trimRootPath(stack)
}

// StackTrace get error stack trace.
func (e *Wrapper) StackTrace() errors.StackTrace {
	f := make([]errors.Frame, len(*e.stack))
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame((*e.stack)[i])
	}
	return f
}

func (e *Wrapper) getFuncFileLine() (string, string, int) {
	st := e.StackTrace()
	f := st[0]
	pc := uintptr(f) - 1
	fn := runtime.FuncForPC(pc)
	file, line := fn.FileLine(pc)
	return funcname(fn.Name()), file, line
}

func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

func trimRootPath(file string) string {
	lastBin := strings.LastIndex(file, rootPath)
	if (lastBin+len(rootPath)) > len(file) || lastBin == -1 {
		return file
	}
	return file[lastBin+len(rootPath):]
}

func callers(pos int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[pos:n]
	return &st
}

func merge(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
