package embedding

import "fmt"

type Person struct {
	Id   int64
	Name string
}

func (personInstance *Person) PrintName() error {
	fmt.Println(personInstance.Id, personInstance.Name)
	return nil
}

func (personInstance *Person) ReturnGreeting() string {
	return fmt.Sprintf("hello my name is %v and my id is %v", personInstance.Name, personInstance.Id)
}

type Student struct {
	Person
	Class string
}
