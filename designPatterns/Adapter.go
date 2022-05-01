package designPatterns

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (cashPaymentInstance *CashPayment) Pay() {
	fmt.Println("paying using cash")
}

type BankPayment struct{}

func (bankPayment *BankPayment) Pay(accountNumber int) {
	fmt.Printf("paying using bank account %v\n", accountNumber)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	BankAccount int
}

func (bankPaymentAdapter *BankPaymentAdapter) Pay() {
	bankPaymentAdapter.BankPayment.Pay(bankPaymentAdapter.BankAccount)
}

func ProcessPayment(payment Payment) {
	payment.Pay()
}
