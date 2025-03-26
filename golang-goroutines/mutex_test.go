package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex	// seperti queue,go akan antri untuk mengeksekusi variable

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {

				mutex.Lock()	// 
				x = x + 1	// x akan ditambahkan 1 setiap perulangan / manipulasi variable
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter ", x)
}


// =========== //
// RWMutex (Read & Write mutex)
// =========== //
type BankAccount struct {
	RWMutex sync.RWMutex
	Balanced int
}

func (account *BankAccount) AddBalance(amount int) {

	// Mutex WRITE
	account.RWMutex.Lock()
	account.Balanced = account.Balanced + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {

	// Mutex READ
	account.RWMutex.RLock()
	balance := account.Balanced
	account.RWMutex.RUnlock()

	return balance
}

// testing
func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {

		// goroutines
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)	// akan ditambahkan 1 terus 
				fmt.Println("Jumlah Account", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
}



// ============== //
// Deadlock
// ============== //
type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()	// diambil dari UserBalance mutex
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {

	// user 1
	user1.Lock()
	fmt.Println("Lock User1, ", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	// user 2
	user2.Lock()
	fmt.Println("Lock User2, ", user2.Name)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance {
		Name: "Rizky",
		Balance: 1000000,
	}

	user2 := UserBalance {
		Name: "Cahyono",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)


	fmt.Println("User: ", user1.Name, "Balanced: ", user1.Balance)
	fmt.Println("User: ", user2.Name, "Balanced: ", user2.Balance)
}