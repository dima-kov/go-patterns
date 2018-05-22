package main

import (
	"fmt"
	"github.com/dima-kov/go-patterns/state/book-travelling/models"
	"github.com/dima-kov/go-patterns/state/book-travelling/states"
)

func main() {
	user := models.User{"Max"}
	state := states.WaitingForOwner{}
	book := models.Book{state, "415 degree", user}
	fmt.Println(book)
}
