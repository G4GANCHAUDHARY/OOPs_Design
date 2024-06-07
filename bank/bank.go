//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//var (
//	CustomerId         = 0
//	TellerId           = 0
//	BankHeadQuartersId = 0
//	TransactionId      = 0
//	AccountNumber      = 0
//)
//
//type Customers struct {
//	Id            int
//	Name          string
//	AccountNumber int
//	tellerId      int
//	currentAmount int
//}
//
//func NewCustomer(name string, tellerId int) *Customers {
//	CustomerId += 1
//	return &Customers{
//		Id:            CustomerId,
//		Name:          name,
//		tellerId:      tellerId,
//		currentAmount: 0,
//	}
//}
//
//func (c *Customers) GetId() int {
//	return c.Id
//}
//
//func (c *Customers) GetName() string {
//	return c.Name
//}
//
//func (c *Customers) GetAccountNumber() int {
//	return c.AccountNumber
//}
//
//func (c *Customers) GetTellerId() int {
//	return c.tellerId
//}
//
//func (c *Customers) GetCurrentAmount() int {
//	return c.currentAmount
//}
//
//func (c *Customers) SetId(id int) {
//	c.Id = id
//}
//
//func (c *Customers) SetName(name string) {
//	c.Name = name
//}
//
//func (c *Customers) SetAccountNumber(acc int) {
//	c.AccountNumber = acc
//}
//
//func (c *Customers) SetTellerId(tellerId int) {
//	c.tellerId = tellerId
//}
//
//func (c *Customers) SetCurrentAmount(amount int) {
//	c.currentAmount = amount
//}
//
//type BankHeadQuarters struct {
//	Id                int
//	availableAmount   int
//	transactionRecord map[int]map[int][]int
//}
//
//func NewBankHeadQuarters() *BankHeadQuarters {
//	BankHeadQuartersId += 1
//	return &BankHeadQuarters{
//		Id:                BankHeadQuartersId,
//		availableAmount:   0,
//		transactionRecord: make(map[int]map[int][]int),
//	}
//}
//
//func (b *BankHeadQuarters) GetId() int {
//	return b.Id
//}
//
//func (b *BankHeadQuarters) GetAvailableAmount() int {
//	return b.availableAmount
//}
//
//func (b *BankHeadQuarters) GetTransactionRecord() map[int]map[int][]int {
//	return b.transactionRecord
//}
//
//func (b *BankHeadQuarters) SetId(id int) {
//	b.Id = id
//}
//
//func (b *BankHeadQuarters) SetAvailableAmount(amount int) {
//	b.availableAmount = amount
//}
//
//func (b *BankHeadQuarters) SetTransactionRecord(record map[int]map[int][]int) {
//	b.transactionRecord = record
//}
//
//type Tellers struct {
//	Id                int
//	Name              string
//	Location          string
//	availableAmount   int
//	transactionRecord map[int][]int
//	accountRecord     map[int]int
//}
//
//func NewTeller(name string, location string) *Tellers {
//	TellerId += 1
//	return &Tellers{
//		Id:                TellerId,
//		Name:              name,
//		Location:          location,
//		availableAmount:   0,
//		transactionRecord: make(map[int][]int),
//		accountRecord:     make(map[int]int),
//	}
//}
//
//func (t *Tellers) OpenAccount(customer *Customers) {
//	AccountNumber += 1
//	accountRecord := t.GetAccountRecord()
//	accountRecord[customer.GetId()] = AccountNumber
//	customer.SetAccountNumber(AccountNumber)
//}
//
//func (t *Tellers) WithDrawMoney(customer *Customers, money int) bool {
//	if customer.GetCurrentAmount()-money < 0 {
//		return false
//	}
//
//	transaction := NewTransaction(t.GetId(), customer.GetId(), -money)
//
//	customer.SetCurrentAmount(customer.GetCurrentAmount() + transaction.GetAmount())
//	t.SetAvailableAmount(t.GetAvailableAmount() + transaction.GetAmount())
//
//	transactionRecord := t.GetTransactionRecord()
//	if _, found := transactionRecord[customer.GetId()]; !found {
//		transactionRecord[customer.GetId()] = []int{transaction.GetId()}
//	} else {
//		transactionRecord[customer.GetId()] = append(transactionRecord[customer.GetId()], transaction.GetId())
//	}
//	t.SetTransactionRecord(transactionRecord)
//
//	return true
//}
//
//func (t *Tellers) DepositMoney(customer *Customers, money int) {
//	transaction := NewTransaction(t.GetId(), customer.GetId(), money)
//
//	customer.SetCurrentAmount(customer.GetCurrentAmount() + transaction.GetAmount())
//	t.SetAvailableAmount(t.GetAvailableAmount() + transaction.GetAmount())
//
//	transactionRecord := t.GetTransactionRecord()
//	if _, found := transactionRecord[customer.GetId()]; !found {
//		transactionRecord[customer.GetId()] = []int{transaction.GetId()}
//	} else {
//		transactionRecord[customer.GetId()] = append(transactionRecord[customer.GetId()], transaction.GetId())
//	}
//	t.SetTransactionRecord(transactionRecord)
//}
//
//func (t *Tellers) MoneyTransactionWithHeadQuarters(bank *BankHeadQuarters) {
//	for {
//		time.Sleep(24 * time.Second)
//
//		currentAmount := t.GetAvailableAmount()
//		t.SetAvailableAmount(0)
//
//		bank.SetAvailableAmount(bank.GetAvailableAmount() + currentAmount)
//
//		break
//	}
//}
//
//func (t *Tellers) GetId() int {
//	return t.Id
//}
//
//func (t *Tellers) GetName() string {
//	return t.Name
//}
//
//func (t *Tellers) GetLocation() string {
//	return t.Location
//}
//
//func (t *Tellers) GetTransactionRecord() map[int][]int {
//	return t.transactionRecord
//}
//
//func (t *Tellers) GetAccountRecord() map[int]int {
//	return t.accountRecord
//}
//
//func (t *Tellers) GetAvailableAmount() int {
//	return t.availableAmount
//}
//
//func (t *Tellers) SetAvailableAmount(amount int) {
//	t.availableAmount = amount
//}
//
//func (t *Tellers) SetId(id int) {
//	t.Id = id
//}
//
//func (t *Tellers) SetName(name string) {
//	t.Name = name
//}
//
//func (t *Tellers) SetLocation(loc string) {
//	t.Location = loc
//}
//
//func (t *Tellers) SetTransactionRecord(record map[int][]int) {
//	t.transactionRecord = record
//}
//
//func (t *Tellers) SetAccountRecord(record map[int]int) {
//	t.accountRecord = record
//}
//
//type Transactions struct {
//	Id         int
//	Amount     int
//	TellerId   int
//	CustomerId int
//}
//
//func NewTransaction(TellerId int, CustomerId int, amount int) *Transactions {
//	TransactionId += 1
//	return &Transactions{
//		Id:         TransactionId,
//		Amount:     amount,
//		TellerId:   TellerId,
//		CustomerId: CustomerId,
//	}
//}
//
//func (tr *Transactions) GetId() int {
//	return tr.Id
//}
//
//func (tr *Transactions) GetAmount() int {
//	return tr.Amount
//}
//
//func (tr *Transactions) GetTellerId() int {
//	return tr.TellerId
//}
//
//func (tr *Transactions) GetCustomerId() int {
//	return tr.CustomerId
//}
//
//func (tr *Transactions) SetId(id int) {
//	tr.Id = id
//}
//
//func (tr *Transactions) SetAmount(amount int) {
//	tr.Amount = amount
//}
//
//func (tr *Transactions) SetTellerId(id int) {
//	tr.TellerId = id
//}
//
//func (tr *Transactions) SetCustomerId(id int) {
//	tr.CustomerId = id
//}
//
//func main() {
//	bank := NewBankHeadQuarters()
//
//	tellerOne := NewTeller("HDFC bank, Bengaluru", "IndraNagar")
//
//	customerOne := NewCustomer("Pagan", tellerOne.GetId())
//	tellerOne.DepositMoney(customerOne, 100)
//	tellerOne.DepositMoney(customerOne, 1500)
//
//	if !tellerOne.WithDrawMoney(customerOne, 200) {
//		fmt.Print("Not enough money to withdraw\n")
//	}
//
//	fmt.Printf("Customer balance : %v\n", customerOne.GetCurrentAmount())
//	fmt.Printf("Teller balance : %v\n", tellerOne.GetAvailableAmount())
//	fmt.Printf("Bank Balance : %v\n", bank.GetAvailableAmount())
//
//	tellerOne.MoneyTransactionWithHeadQuarters(bank)
//
//	fmt.Printf("Customer balance : %v\n", customerOne.GetCurrentAmount())
//	fmt.Printf("Teller balance : %v\n", tellerOne.GetAvailableAmount())
//	fmt.Printf("Bank Balance : %v\n", bank.GetAvailableAmount())
//}

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	customerID         = 0
	tellerID           = 0
	bankHeadQuartersID = 0
	transactionID      = 0
	accountNumber      = 0
	mu                 sync.Mutex
)

type Customers struct {
	id            int
	name          string
	accountNumber int
	tellerID      int
	currentAmount int
}

func NewCustomer(name string, tellerID int) *Customers {
	mu.Lock()
	defer mu.Unlock()
	customerID++
	return &Customers{
		id:            customerID,
		name:          name,
		tellerID:      tellerID,
		currentAmount: 0,
	}
}

func (c *Customers) GetID() int {
	return c.id
}

func (c *Customers) GetName() string {
	return c.name
}

func (c *Customers) GetAccountNumber() int {
	return c.accountNumber
}

func (c *Customers) GetTellerID() int {
	return c.tellerID
}

func (c *Customers) GetCurrentAmount() int {
	return c.currentAmount
}

func (c *Customers) SetAccountNumber(acc int) {
	c.accountNumber = acc
}

func (c *Customers) SetCurrentAmount(amount int) {
	c.currentAmount = amount
}

type BankHeadQuarters struct {
	id                int
	availableAmount   int
	transactionRecord map[int]map[int][]int
	mu                sync.Mutex
}

func NewBankHeadQuarters() *BankHeadQuarters {
	mu.Lock()
	defer mu.Unlock()
	bankHeadQuartersID++
	return &BankHeadQuarters{
		id:                bankHeadQuartersID,
		availableAmount:   0,
		transactionRecord: make(map[int]map[int][]int),
	}
}

func (b *BankHeadQuarters) GetID() int {
	return b.id
}

func (b *BankHeadQuarters) GetAvailableAmount() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.availableAmount
}

func (b *BankHeadQuarters) SetAvailableAmount(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.availableAmount = amount
}

func (b *BankHeadQuarters) AddTransactionRecord(tellerID, customerID, transactionID int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, found := b.transactionRecord[tellerID]; !found {
		b.transactionRecord[tellerID] = make(map[int][]int)
	}
	b.transactionRecord[tellerID][customerID] = append(b.transactionRecord[tellerID][customerID], transactionID)
}

type Tellers struct {
	id                int
	name              string
	location          string
	availableAmount   int
	transactionRecord map[int][]int
	accountRecord     map[int]int
	mu                sync.Mutex
}

func NewTeller(name, location string) *Tellers {
	mu.Lock()
	defer mu.Unlock()
	tellerID++
	return &Tellers{
		id:                tellerID,
		name:              name,
		location:          location,
		availableAmount:   0,
		transactionRecord: make(map[int][]int),
		accountRecord:     make(map[int]int),
	}
}

func (t *Tellers) OpenAccount(customer *Customers) {
	mu.Lock()
	defer mu.Unlock()
	accountNumber++
	t.accountRecord[customer.GetID()] = accountNumber
	customer.SetAccountNumber(accountNumber)
}

func (t *Tellers) WithDrawMoney(customer *Customers, money int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if customer.GetCurrentAmount() < money {
		return fmt.Errorf("insufficient funds")
	}

	transaction := NewTransaction(t.id, customer.GetID(), -money)
	customer.SetCurrentAmount(customer.GetCurrentAmount() + transaction.GetAmount())
	t.availableAmount += transaction.GetAmount()

	if _, found := t.transactionRecord[customer.GetID()]; !found {
		t.transactionRecord[customer.GetID()] = []int{transaction.GetID()}
	} else {
		t.transactionRecord[customer.GetID()] = append(t.transactionRecord[customer.GetID()], transaction.GetID())
	}

	return nil
}

func (t *Tellers) DepositMoney(customer *Customers, money int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	transaction := NewTransaction(t.id, customer.GetID(), money)
	customer.SetCurrentAmount(customer.GetCurrentAmount() + transaction.GetAmount())
	t.availableAmount += transaction.GetAmount()

	if _, found := t.transactionRecord[customer.GetID()]; !found {
		t.transactionRecord[customer.GetID()] = []int{transaction.GetID()}
	} else {
		t.transactionRecord[customer.GetID()] = append(t.transactionRecord[customer.GetID()], transaction.GetID())
	}
}

func (t *Tellers) MoneyTransactionWithHeadQuarters(bank *BankHeadQuarters) {
	for {
		time.Sleep(24 * time.Second)

		t.mu.Lock()
		currentAmount := t.availableAmount
		t.availableAmount = 0
		t.mu.Unlock()

		bank.SetAvailableAmount(bank.GetAvailableAmount() + currentAmount)
		break
	}
}

type Transactions struct {
	id         int
	amount     int
	tellerID   int
	customerID int
}

func NewTransaction(tellerID, customerID, amount int) *Transactions {
	mu.Lock()
	defer mu.Unlock()
	transactionID++
	return &Transactions{
		id:         transactionID,
		amount:     amount,
		tellerID:   tellerID,
		customerID: customerID,
	}
}

func (tr *Transactions) GetID() int {
	return tr.id
}

func (tr *Transactions) GetAmount() int {
	return tr.amount
}

func main() {
	bank := NewBankHeadQuarters()

	tellerOne := NewTeller("HDFC bank", "Bengaluru")

	customerOne := NewCustomer("Pagan", tellerOne.id)
	tellerOne.OpenAccount(customerOne)
	tellerOne.DepositMoney(customerOne, 100)
	tellerOne.DepositMoney(customerOne, 1500)

	if err := tellerOne.WithDrawMoney(customerOne, 200); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Customer balance: %v\n", customerOne.GetCurrentAmount())
	fmt.Printf("Teller balance: %v\n", tellerOne.availableAmount)
	fmt.Printf("Bank balance: %v\n", bank.GetAvailableAmount())

	tellerOne.MoneyTransactionWithHeadQuarters(bank)

	fmt.Printf("Customer balance: %v\n", customerOne.GetCurrentAmount())
	fmt.Printf("Teller balance: %v\n", tellerOne.availableAmount)
	fmt.Printf("Bank balance: %v\n", bank.GetAvailableAmount())
}
