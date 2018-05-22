package states

import (
	"errors"
	"github.com/dima-kov/patterns/book-travelling/models"
)

type WaitingForOwner struct {
	Book models.Book
}

func (wo WaitingForOwner) RequestByUser(user interface{}) error {
	return errors.New("another user is reading this book right now. Wait")
}

func (wo WaitingForOwner) ConfirmByOwner(user interface{}) error {
	requestUser := user.(models.User)
	if requestUser != wo.Book.Owner {
		return errors.New("wou, you are not the owner. We`ll report about u")
	}
	send_notigfication_that_book_will_arrive_soon(wo.Book.Owner) // TODO change Book.Owner to bookreading.user
	wo.Book.SetState(ConfirmedByOwner{})
	return nil
}

func (wo WaitingForOwner) WriteFeedback(user interface{}) error {
	// do nothing
	return nil
}
