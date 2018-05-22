package mediator

type Mediator interface {
	reactOnSubmit() error
	reactOnCheckbox() error
	reactOnTextFieldChanged() error
}