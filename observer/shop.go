package main

import (
	"errors"
	"git.vakoms.com/steakager/steakager-backend/app/core/emails"
)

type Shop struct {
	Name string

	Subscribers []*Subscriber
}

func (s *Shop) Subscribe(subscriber *Subscriber) error {
	if yes, _ := s.checkAlreadySubscribed(subscriber); !yes {
		s.Subscribers = append(s.Subscribers, subscriber)
		return nil
	}
	return errors.New("already subscribed")
}

func (s *Shop) RemoveSubscription(subscriber *Subscriber) error {
	if yes, i := s.checkAlreadySubscribed(subscriber); yes {
		s.Subscribers = append(s.Subscribers[:i], s.Subscribers[i+1:]...)
		return nil
	}
	return errors.New("no such subscriber")

}

func (s *Shop) NotifySubscribers(message string) {
	//emails.Send(s.Subscribers)
}

func (s *Shop) checkAlreadySubscribed(subscriber *Subscriber)  (bool, int) {
	for i := range s.Subscribers {
		if s.Subscribers[i] == subscriber {
			return true, i
		}
	}
	return false, 0
}