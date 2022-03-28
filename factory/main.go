package factory

import (
	"errors"
	"fmt"
)

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type SmsNotificationSender struct {
}

func (SmsNotificationSender) GetSenderMethod() string {
	return "Sms"
}

func (SmsNotificationSender) GetSenderChannel() string {
	return "twilio"
}

type SmsNotification struct {
}

func (SmsNotification) SendNotification() {
	fmt.Println("sendings notification via sms")
}

func (SmsNotification) GetSender() ISender {
	return SmsNotificationSender{}
}

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("this is an email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return ".amazon"
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "sms" {
		return SmsNotification{}, nil
	}
	if notificationType == "email" {
		return EmailNotification{}, nil
	}
	return nil, errors.New("not valid notificationType")
}
