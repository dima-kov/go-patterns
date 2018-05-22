package chains

type BookTransferring interface {
	execute() error
}