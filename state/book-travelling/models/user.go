package models

type User struct {
	Name string
}

func (u User) HasBook() bool {
	return true
}

func (u User) HasEnoughMoney() bool {
	return true
}
