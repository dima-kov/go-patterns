package observer

type ShopNewsletter interface {
	Subscribe(subscriber Subscriber) error
	RemoveSubscription(subscriber Subscriber) error

	NotifySubscribers(message string)
}

type Subscriber interface {
	Notification(ShopNewsletter) error
}
