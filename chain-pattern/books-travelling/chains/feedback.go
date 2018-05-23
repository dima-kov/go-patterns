package chains

type readerFeedback struct {
	next BookTransferring
	book interface{}
}

func NewReaderFeedback(book interface{}) BookTransferring {
	return readerFeedback{book: book}
}

func (r readerFeedback) execute() error {
	return nil
}
