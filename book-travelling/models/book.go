package models

import (
	"github.com/dima-kov/patterns/book-travelling/states/interface"
)

type Book struct {
	State state_interface.BookState
	Name  string
	Owner User
}

func (b *Book) SetState(state state_interface.BookState) {
	b.State = state
}
