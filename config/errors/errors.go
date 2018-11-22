package errors

import "fmt"

//ErrCustom ErrCustom struct
type ErrCustom struct {
	Msg string `json:"error"`
}

//Custom Custom interface
type Custom interface {
	Error() string
}

//NewCustom Shows a message error.
func NewCustom(msg string) Custom {
	return &ErrCustom{
		Msg: msg,
	}
}

func (e *ErrCustom) Error() string {
	return fmt.Sprintf(e.Msg)
}

//BadRequest 400
var BadRequest = NewCustom("Parameters Needed")

//Unauthorized 401
var Unauthorized = NewCustom("Unauthorized")

//NotFound 404
var NotFound = NewCustom("Article not found")

//InternalError 500
var InternalError = NewCustom("Internal Error")

//ConnectionError Rabbit Connection Error
var ConnectionError = NewCustom("There was a problem connecting with Rabbit")

//ChannelError Rabbit Channel Error
var ChannelError = NewCustom("Channel not Channel not initialized")

//ExchangeDeclareError Rabbitn Queue Declare Error
var ExchangeDeclareError = NewCustom("There was an error declaring a exchange")

//QueueDeclareError Rabbitn Queue Declare Error
var QueueDeclareError = NewCustom("There was an error declaring a queue")

//QueueBindError Rabbitn Queue Declare Error
var QueueBindError = NewCustom("There was an error binding a queue")

//ConsumeError Rabbitn Queue Declare Error
var ConsumeError = NewCustom("There was an error consuming a message")
