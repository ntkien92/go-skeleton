package errors

import (
	"encoding/json"
	"reflect"
	"runtime"
	"strconv"
)

type ErrorItem struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Caller  *CallerItem `json:"caller"`
	Func    string      `json:"func"`
}

type CallerItem struct {
	File string
	Line int
}

func (item *ErrorItem) Error() string {
	return "[" + strconv.Itoa(item.Code) + "]" + item.Message + " -> " + item.ToJSON()
}

func (item *ErrorItem) ToJSON() string {
	json, _ := json.Marshal(item)
	return string(json)
}

// constructor for new ErrorItem
func New(
	code int,
	message string,
	data interface{},
) error {

	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)

	err := ErrorItem{
		Code:    code,
		Message: message,
		Data:    data,
		Caller: &CallerItem{
			File: file,
			Line: line,
		},
		Func: f.Name(),
	}

	return &err
}

func IsErrorItem(err error) bool {
	if reflect.TypeOf(err).String() == "*errors.ErrorItem" {
		return true
	} else {
		return false
	}
}

func ToErrorItem(err error) *ErrorItem {
	if reflect.TypeOf(err).String() == "*errors.ErrorItem" {
		return err.(*ErrorItem)
	} else {
		return nil
	}
}
