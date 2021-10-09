package ErrorHandling

import (
	"errors"
	"fmt"
)

func recoverFunc() {
	//recover from panic
	panicState := recover()
	fmt.Printf("%v", panicState)
	if panicState != nil {
		fmt.Printf("error provoked a panic, now its recovered\n")
	}
}

func Something() {
	//there is nothing special here,this is just so we can
	//call a panic
	defer fmt.Println("closed something")
	defer recoverFunc()
	var newErr error = nil
	for i := 0; i < 5; i++ {
		fmt.Printf("%v called\n", i)
		if i >= 2 {
			newErr = errors.New("2 was called ")
			panic(newErr)
		}
	}
}
