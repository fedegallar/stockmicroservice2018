package errors

import "fmt"

type ErrCustom struct {
	Msg string `json:"error"`
}
type Custom interface {
	Error() string
}

func NewCustom(msg string) Custom {
	return &ErrCustom{
		Msg: msg,
	}
}

func (e *ErrCustom) Error() string {
	return fmt.Sprintf(e.Msg)
}

//400
var BadRequest = NewCustom("Parameters Needed")

//401
var Unauthorized = NewCustom("Unauthorized")

//404
var NotFound = NewCustom("Article not found")

//418
var Teapot = NewCustom("I'm a Teapot")

//500
var InternalError = NewCustom("Internal Error")
