package errs

import "fmt"

type Error struct {
	Code  int
	Msg   string
	cause error
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.cause != nil {
		return fmt.Sprintf("code:%d msg:%s caused by %s", e.Code, e.Msg, e.cause)
	}
	return fmt.Sprintf("code:%d msg:%s", e.Code, e.Msg)
}

func New(code int, msg string) error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func Newf(code int, format string, args ...interface{}) error {
	return &Error{
		Code: code,
		Msg:  fmt.Sprintf(format, args...),
	}
}

func Wrap(err error, code int, msg string) error {
	if err == nil {
		return nil
	}
	wrapError := &Error{
		Code:  code,
		Msg:   msg,
		cause: err,
	}
	return wrapError
}

func Wrapf(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	wrapError := &Error{
		Code:  code,
		Msg:   fmt.Sprintf(format, args...),
		cause: err,
	}
	return wrapError
}

func Code(e error) int {
	if e == nil {
		return 0
	}
	err, ok := e.(*Error)
	if !ok {
		return 999
	}
	if err == (*Error)(nil) {
		return 0
	}
	return err.Code
}

func Msg(e error) string {
	if e == nil {
		return ""
	}
	err, ok := e.(*Error)
	if !ok {
		return e.Error()
	}
	if err == (*Error)(nil) {
		return ""
	}
	if err.cause != nil {
		return err.Error()
	}
	return err.Msg
}
