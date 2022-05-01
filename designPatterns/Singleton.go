package designPatterns

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

var db *Database = nil
var locker *sync.Mutex = new(sync.Mutex)

func (dbInstance *Database) CreateSingleConnection() {
	fmt.Println("connecting to database")
	time.Sleep(time.Second * 2)
	fmt.Println("finished creating database")
}

func GetDatabaseInstance() *Database {
	locker.Lock()
	if db == nil {
		fmt.Println("creating database")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("data already created")
	}
	defer locker.Unlock()
	return db
}
