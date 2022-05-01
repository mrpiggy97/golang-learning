package designPatterns

import "fmt"

type Observer interface {
	GetId() string
	UpdateValue(string)
}

type Topic interface {
	Register(oberver Observer)
	BroadCast()
}

type Item struct {
	Observers []Observer
	Name      string
	Available bool
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (item *Item) UpdateAvailable() {
	fmt.Printf("item %s is now available\n", item.Name)
	item.Available = true
	item.BroadCast()
}

func (item *Item) BroadCast() {
	for _, obs := range item.Observers {
		obs.UpdateValue(item.Name)
	}
}

func (item *Item) Register(obs Observer) {
	item.Observers = append(item.Observers, obs)
}

type EmailClient struct {
	Id string
}

func (emailClient *EmailClient) GetId() string {
	return emailClient.Id
}

func (emailClient *EmailClient) UpdateValue(value string) {
	fmt.Printf("sending email %s from client %s\n", value, emailClient.Id)
}
