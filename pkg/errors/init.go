package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type StatusCode int32

// stack represents a stack of program counters.
type stack []uintptr

type Wrapper struct {
	httpStatus int
	code       StatusCode
	message    string
	devMessage error
	stack      *stack
	metadata   map[string]interface{}
}

// New creates wrapper error
func New(code StatusCode, err error) error {
	var msg string
	var httpStatus int

	if err == nil {
		msg = msgOK
		httpStatus = http.StatusOK
	}

	msg, ok := errMsg[code]
	if !ok {
		msg = msgUnclassified
	}

	httpStatus, ok = errHttpStatus[code]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}

	return &Wrapper{
		httpStatus: httpStatus,
		code:       code,
		message:    msg,
		devMessage: err,
		stack:      callers(0),
	}
}

// New creates wrapper error with optional message
func NewWithMessage(code StatusCode, message string) error {
	httpStatus, ok := errHttpStatus[code]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}

	return &Wrapper{
		httpStatus: httpStatus,
		code:       code,
		message:    message,
		devMessage: errors.New(message),
		stack:      callers(0),
	}
}

func (e *Wrapper) Error() string {
	err := fmt.Sprintf("error: %v %s", e.code, e.message)
	if e.devMessage == nil {
		return err
	}
	return fmt.Sprintf("%s %s", err, e.devMessage.Error())
}

func CastToErrorWrapper(err error) *Wrapper {
	if err == nil {
		return nil
	}

	errorsper, ok := err.(*Wrapper)
	if !ok || errorsper == nil {
		return nil
	}

	return errorsper
}

func AddMetadata(err error, metadata map[string]interface{}) {
	errorsper := CastToErrorWrapper(err)
	if errorsper != nil {
		errorsper.metadata = merge(errorsper.metadata, metadata)
	}
}

func (e *Wrapper) GetCode() StatusCode {
	return e.code
}

func (e *Wrapper) GetHttpStatus() int {
	return e.httpStatus
}

func (e *Wrapper) GetUserMessage() string {
	return e.message
}

func (e *Wrapper) GetDevMessage() error {
	return e.devMessage
}

func (e *Wrapper) GetMetadata() map[string]interface{} {
	return e.metadata
}

func (e *Wrapper) IsLoggable() bool {
	if e.httpStatus == http.StatusOK || e.devMessage == nil {
		return false
	}

	return true
}
