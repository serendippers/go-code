package go_sync

import (
	"fmt"
	"sync"
	"time"
)

type user struct {
	mutex sync.Mutex
	Name  string
}

func GoMutex() {
	var user = user{Name: "zpp"}

	wait := sync.WaitGroup{}

	fmt.Println("Locked")
	user.mutex.Lock()

	for i := 0; i < 3; i++ {
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)
			user.Name = fmt.Sprintf("zpp:%d号", i)
			user.mutex.Lock()
			fmt.Println("Lock:", i)
			user.Name = fmt.Sprintf("zpp:%d号", i)
			fmt.Printf("Lock:%d,Name:%s\n", i, user.Name)
			time.Sleep(time.Second * 3)
			user.mutex.Unlock()
			fmt.Println("Unlock:", i)
			defer wait.Done()
		}(i)
	}
	user.Name = "zhangpp"
	time.Sleep(time.Second * 3)
	fmt.Printf("Unlocked, name:%s\n", user.Name)
	user.mutex.Unlock()
	wait.Wait()

}
