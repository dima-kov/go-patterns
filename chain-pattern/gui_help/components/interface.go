package components

type Component interface {
	ShowHelp()
	SetContainer(cont Container)
}

type Container interface {
	ShowHelp()
	AddItems(items ...Component)
}
