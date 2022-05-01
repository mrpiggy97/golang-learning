package designPatterns

import (
	"errors"
	"fmt"
)

// top interface every object will implement
type IProduct interface {
	SetStock(stock int)
	GetStock() int
	SetName(name string)
	GetName() string
}

// top level object that will implement IProduct
type Computer struct {
	Name  string
	Stock int
}

func (computerInstance *Computer) SetStock(stock int) {
	computerInstance.Stock = stock
}

func (computerInstance *Computer) GetStock() int {
	return computerInstance.Stock
}

func (computerInstance *Computer) SetName(name string) {
	computerInstance.Name = name
}

func (computerInstance *Computer) GetName() string {
	return computerInstance.Name
}

// Laptop and Desktop will inherit all methods from
// Computer

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			Name:  "dell",
			Stock: 12,
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			Name:  "hp",
			Stock: 43,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "Desktop":
		return NewDesktop(), nil
	case "Laptop":
		return NewLaptop(), nil
	default:
		return nil, errors.New("computerType is invalid")
	}
}

func PrintNameAndStock(product IProduct) {
	mssg := fmt.Sprintf(
		"product with name %v, and stock %d",
		product.GetName(),
		product.GetStock(),
	)
	fmt.Println(mssg)
}
