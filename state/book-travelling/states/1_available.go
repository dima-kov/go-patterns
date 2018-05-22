package states

import (
	"errors"
	"github.com/dima-kov/go-patterns/state/book-travelling/models"
)

type Available struct {
	Book models.Book
}

func (av Available) RequestByUser(user interface{}, fake bool) error {
	requestUser := user.(models.User)
	if requestUser.HasBook() {
		return errors.New("you have one book now")
	}
	if !requestUser.HasEnoughMoney() {
		return errors.New("make money first to buy it")
	}
	if !fake {
		//send_notification_for_book_owner()
		av.Book.SetState(WaitingForOwner{av.Book})
	}
	return nil
}

func (av Available) ConfirmByOwner(user interface{}) error {
	return errors.New("you can`t confirm to read book now, book is not booked yet")
}

func (av Available) WriteFeedback(user interface{}) error {
	return errors.New("you can`t write feedback to book now, book is not booked yet")
}
