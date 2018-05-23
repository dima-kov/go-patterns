package chains

type ownerConfirmTransferring struct {
	next BookTransferring
	book interface{}
}

func NewOwnerConfirmTransferring(book interface{}) BookTransferring {
	return ownerConfirmTransferring{next:NewReaderFeedback(book), book:book}
}

func (o ownerConfirmTransferring) execute() error {
	return nil
}