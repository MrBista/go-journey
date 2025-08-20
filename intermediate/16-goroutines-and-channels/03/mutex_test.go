package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutexBasictest(t *testing.T) {
	var x = 0
	var count = 1000

	var mutex sync.Mutex

	for i := 0; i < count; i++ {
		go func() {
			for j := 0; j <= count/10; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Counter: ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}
func (account *BankAccount) GetBalance() int {

	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final balance", account.GetBalance())
}
