package state_interface

type BookState interface {
	//GetMessage(requestUser interface{}) string

	RequestByUser(user interface{}) error
	ConfirmByOwner(user interface{}) error
	WriteFeedback(user interface{}) error
}
