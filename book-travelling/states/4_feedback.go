package states

import "github.com/dima-kov/patterns/book-travelling/models"

type ReadFeedback struct {
	Book models.Book
}

func (rf *ReadFeedback) RequestByUser(user interface{}) error {
	return nil
}

func (rf *ReadFeedback) ConfirmByOwner(user interface{}) error {
	return nil
}

func (rf *ReadFeedback) WriteFeedback(user interface{}) error {
	return nil
}
