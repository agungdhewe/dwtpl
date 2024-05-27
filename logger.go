package dwtpl

import (
	"fmt"
	"io"
	"runtime"
)

const (
	colorFgRed   string = "\u001b[31m"
	colorFgWhite string = "\u001b[37m"
	colorReset   string = "\u001b[0m"
	colorFgGray  string = "\033[38;5;245m"
)

func report_log(msg string, args ...any) {
	if mgr.logger.Writer() == io.Discard {
		return
	}

	_, callerfile, lineno, ok := runtime.Caller(1)
	if ok {
		text := fmt.Sprintf(msg, args...)
		lineinfo := fmt.Sprintf("%s%s:%d", colorFgGray, callerfile, lineno)
		mgr.logger.Printf("%s%s %s%s", colorFgWhite, text, lineinfo, colorReset)
	}
}

func report_error(msg string, args ...any) {
	if mgr.logger.Writer() == io.Discard {
		return
	}

	_, callerfile, lineno, ok := runtime.Caller(1)
	if ok {
		text := fmt.Sprintf(msg, args...)
		lineinfo := fmt.Sprintf("%s%s:%d", colorFgGray, callerfile, lineno)
		mgr.logger.Printf("%s%s %s%s", colorFgRed, text, lineinfo, colorReset)
	}
}
