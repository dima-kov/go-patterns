package states

import "github.com/dima-kov/go-patterns/state/book-travelling/models"

type ConfirmedByOwner struct {
	Book models.Book
}

func (c ConfirmedByOwner) RequestByUser(user interface{}) error {
	return nil
}

func (c ConfirmedByOwner) ConfirmByOwner(user interface{}) error {
	return nil
}

func (c ConfirmedByOwner) WriteFeedback(user interface{}) error {
	return nil
}
