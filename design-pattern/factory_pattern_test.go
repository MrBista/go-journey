package designpattern

import (
	"fmt"
	"testing"
)

/*
The factory pattern provides an interface for creating objects in a super clean and abstract way.

Use Case:
Imagine you have different notification channels (Email, SMS, Push) and want a unified interface to create them.

Benefits:
Separation of concerns
Easier to test and mock different implementations

*/

type Notifier interface {
	Send(msg string)
}

type EmailNotifier struct {
}

func (e *EmailNotifier) Send(msg string) {
	fmt.Println("Hello send message from Email Notifier with message: ", msg)
}

type SmsNotifier struct{}

func (e *SmsNotifier) Send(msg string) {
	fmt.Println("Hello send message from Sms Notifier with message: ", msg)

}

func GetNotifier(channel string) Notifier {
	switch channel {
	case "email":
		return &EmailNotifier{}
	case "sms":
		return &SmsNotifier{}

	default:
		return nil
	}
}

func TestFactoryPattern(t *testing.T) {
	notifer := GetNotifier("email")
	notifer.Send("Hallo abang-abang jago sekalian")
}
