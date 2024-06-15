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

// report_log logs a message with optional arguments and additional information about the caller.
//
// Parameters:
// - msg: a string representing the message to be logged.
// - args: optional arguments to be formatted into the message.
//
// Returns: None.
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

// report_error logs an error message with optional arguments and additional information about the caller.
//
// Parameters:
// - msg: a string representing the error message to be logged.
// - args: optional arguments to be formatted into the error message.
//
// The function logs the error message along with the file name and line number of the caller.
// If the logger's writer is set to io.Discard, the error message will not be logged.
// The error message is formatted with color codes to highlight the error.
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
