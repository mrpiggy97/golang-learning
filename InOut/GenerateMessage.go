package InOut

import (
	"fmt"
	"sync"
	"time"
)

func SendMessage(channel chan<- map[string]string, waiter *sync.WaitGroup) {
	defer waiter.Done()
	waiter.Add(1)
	for i := 0; i <= 10; i++ {
		var funcName string = fmt.Sprintf("SayHello%v", i)
		time.Sleep(time.Second * 1)
		var code string = fmt.Sprintf(`
		func %v(){fmt.Printf("this is the salute %v")}
		`, funcName, i)
		var funcInfo map[string]string = make(map[string]string)
		funcInfo["funcName"] = funcName + ","
		funcInfo["func"] = code
		channel <- funcInfo
	}
	close(channel)
}
