package embedding

import (
	"testing"
)

func TestPrintName(testCase *testing.T) {
	var student1 Student = Student{
		Person: Person{Id: 11222, Name: "fabian"},
		Class:  "4b",
	}
	var err error = student1.PrintName()
	if err != nil {
		testCase.Error(err)
	}
}

func TestGetGreeting(testCase *testing.T) {
	var student1 Student = Student{
		Person: Person{Name: "fabian", Id: 1212},
		Class:  "intro to c++",
	}
	var expectedInfo string = "hello my name is fabian and my id is 1212"
	if student1.ReturnGreeting() != expectedInfo {
		testCase.Errorf("expected student1.ReturnGreeting() to return %v", expectedInfo)
	}
}
