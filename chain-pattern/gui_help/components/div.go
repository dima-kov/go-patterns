package components

type Div struct {
	HelpMessage string

	container Container
	children []Component
}

func (d *Div) ShowHelp() {
	d.container.ShowHelp()
}

func (d *Div) SetContainer(cont Container) {
	d.container = cont
}

func (d *Div) AddItems(items ...Component) {
	d.children = append(d.children, items...)
	for _, item:= range items {
		item.SetContainer(d)
	}
}
