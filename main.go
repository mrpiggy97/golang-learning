package main

import (
	"sync"

	"github.com/mrpiggy97/golang-learning/designPatterns"
)

// Main will apply all concepts learned.

func main() {
	//factory
	laptop, _ := designPatterns.GetComputerFactory("Laptop")
	desktop, _ := designPatterns.GetComputerFactory("Desktop")
	designPatterns.PrintNameAndStock(laptop)
	designPatterns.PrintNameAndStock(desktop)

	//singleton
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	for i := 0; i <= 5; i++ {
		waiter.Add(1)
		go func() {
			designPatterns.GetDatabaseInstance()
			defer waiter.Done()
		}()
	}
	waiter.Wait()

	//adapter
	var cash *designPatterns.CashPayment = &designPatterns.CashPayment{}
	designPatterns.ProcessPayment(cash)
	var bank *designPatterns.BankPayment = &designPatterns.BankPayment{}
	var adapter *designPatterns.BankPaymentAdapter = &designPatterns.BankPaymentAdapter{
		BankPayment: bank,
		BankAccount: 300,
	}
	designPatterns.ProcessPayment(adapter)

	//observer
	var graphicsCard *designPatterns.Item = designPatterns.NewItem("graphics card")
	var firstObserver designPatterns.Observer = &designPatterns.EmailClient{
		Id: "2va",
	}
	var secondObserver designPatterns.Observer = &designPatterns.EmailClient{
		Id: "44aaa",
	}
	graphicsCard.Register(firstObserver)
	graphicsCard.Register(secondObserver)
	graphicsCard.UpdateAvailable()

	//strategy
	var shaAlgorithm *designPatterns.SHA = &designPatterns.SHA{}
	var md5Algorithm *designPatterns.MD5 = &designPatterns.MD5{}
	var passwordProtector *designPatterns.PasswordProtector = designPatterns.NewPasswordProtector(
		"ragnar",
		"lothbrok123s!",
		shaAlgorithm,
	)
	passwordProtector.Hash()
	passwordProtector.SetAlgorithm(md5Algorithm)
	passwordProtector.Hash()
}
