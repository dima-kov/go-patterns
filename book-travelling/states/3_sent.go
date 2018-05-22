package states

import "github.com/dima-kov/patterns/book-travelling/models"

type SentByPostAndReading struct {
	Book models.Book
}

func (rr *SentByPostAndReading) RequestByUser(user interface{}) error {
	return nil
}

func (rr *SentByPostAndReading) ConfirmByOwner(user interface{}) error {
	return nil
}

func (rr *SentByPostAndReading) WriteFeedback(user interface{}) error {
	return nil
}
