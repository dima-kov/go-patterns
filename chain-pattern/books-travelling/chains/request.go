package chains

type requestBookTransferring struct {
	next BookTransferring
	book interface{}
}

func NewRequestBookTransferring(book interface{}) BookTransferring{
	return requestBookTransferring{next: NewOwnerConfirmTransferring(book), book: book}
}

func (r requestBookTransferring) execute() error {
	// if !currentUser.hasEnoughMoney() {
	//	return errors.New("Not enough money")
	// }
	// err := send_email_to(r.book.owner, "Someone wants to read your book")
	//if err != nil {
	//	return err
	//}
	r.next.execute()
	return nil
}