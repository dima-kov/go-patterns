package main

import (
	"fmt"
	"github.com/dima-kov/patterns/book-travelling/models"
	"github.com/dima-kov/patterns/book-travelling/states"
)

func main() {
	user := models.User{"Max"}
	state := states.WaitingForOwner{}
	book := models.Book{state, "415 degree", user}
	fmt.Println(book)
}
