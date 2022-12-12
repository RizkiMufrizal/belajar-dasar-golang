package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("counter", x)
}

func TestRaceConditionMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("counter", x)
}

type BankAccount struct {
	Mutex   sync.RWMutex
	Balance int
}

func (bankAccount *BankAccount) AddBalance(balance int) {
	bankAccount.Mutex.Lock()
	bankAccount.Balance = balance + bankAccount.Balance
	bankAccount.Mutex.Unlock()
}

func (bankAccount *BankAccount) GetBalance() int {
	bankAccount.Mutex.RLock()
	balance := bankAccount.Balance
	bankAccount.Mutex.RUnlock()
	return balance
}

func TestRwMutex(t *testing.T) {
	bankAccount := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				bankAccount.AddBalance(1)
				fmt.Println(bankAccount.Balance)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("balance", bankAccount.Balance)
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (userBalance *UserBalance) Lock() {
	userBalance.Mutex.Lock()
}

func (userBalance *UserBalance) Unlock() {
	userBalance.Mutex.Unlock()
}

func (userBalance *UserBalance) Change(amount int) {
	userBalance.Balance = userBalance.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Rizki",
		Balance: 100000,
	}
	user2 := UserBalance{
		Name:    "Mufrizal",
		Balance: 100000,
	}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 20000)

	time.Sleep(3 * time.Second)

	fmt.Println("user 1", user1.Name, "balance", user1.Balance)
	fmt.Println("user 2", user2.Name, "balance", user2.Balance)

}
