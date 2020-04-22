package errors

import (
	"errors"
	"fmt"
	"github.com/ltto/T/gobox/utils"
	"reflect"
	"runtime"
)

type E struct {
	file  string
	line  int
	msg   string
	e     error
	cause *E
}

func NewE(msg string, cause error) *E {
	_, file, line, _ := runtime.Caller(1)
	err := &E{msg: msg, cause: nil, file: file, line: line}
	err.e = err
	if cause != nil {
		e := &E{}
		if errors.As(cause, &e) {
			err.cause = e
		}
		err.cause = &E{msg: cause.Error(), e: cause}
	}
	return err
}
func (e *E) Error() string {
	title := fmt.Sprintf("Error in Gorouter %d", utils.GetGID())
	return fmt.Sprintf("%s %s%s", Red(title), e.String(), e.Cause())
}
func (e *E) String() string {
	f := e.file
	l := e.line
	if f == "" {
		var in interface{}
		in = e.e
		path := reflect.TypeOf(in).String()
		return fmt.Sprintf(Magenta("[%T]")+": \n\t at "+Blue("%s")+" :%s", e.e, path, Red(e.msg))
	} else {
		return fmt.Sprintf(Magenta("[%T]")+": \n\t at "+Blue("%s:%d")+" :%s", e.e, f, l, Red(e.msg))
	}
}
func (e *E) Cause() string {
	if e.cause != nil {
		title := fmt.Sprintf("\nCaused by: %d", utils.GetGID())
		c := fmt.Sprintf("%s %s", Red(title), e.cause.String())
		c += e.cause.Cause()
		return c
	}
	return ""
}

func a(skip int) {
	_, file, line, _ := runtime.Caller(skip)
	fmt.Printf("%s:%d\n", file, line)
}

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func Black(str string) string {
	return textColor(TextBlack, str)
}

func Red(str string) string {
	return textColor(TextRed, str)
}

func Green(str string) string {
	return textColor(TextGreen, str)
}

func Yellow(str string) string {
	return textColor(TextYellow, str)
}

func Blue(str string) string {
	return textColor(TextBlue, str)
}

func Magenta(str string) string {
	return textColor(TextMagenta, str)
}

func Cyan(str string) string {
	return textColor(TextCyan, str)
}

func White(str string) string {
	return textColor(TextWhite, str)
}

func textColor(color int, str string) string {
	if IsWindows() {
		return str
	}

	switch color {
	case TextBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
	case TextMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
	case TextCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}
