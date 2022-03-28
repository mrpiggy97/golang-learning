package employees

import "fmt"

type IEmployee interface {
	GetSalary() int64
	GetName() string
}

type Employee struct {
	Name   string
	Salary int64
}

func (employeeInstance *Employee) GetName() string {
	return employeeInstance.Name
}

func (employeeInstance *Employee) GetSalary() int64 {
	return employeeInstance.Salary
}

func GetMessage(employee IEmployee) string {
	return fmt.Sprintf("hello my name is %v and i make %v", employee.GetName(), employee.GetSalary())
}
