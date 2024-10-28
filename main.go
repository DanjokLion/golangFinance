package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	Amount   float32
	Type     string
	Category string
	Date     time.Time
	Note string
}

type Tracker struct {
	Transactions []Transaction
	Balance       float32
}

func (t *Tracker) AddTransaction(amount float32, date time.Time, note, category, transType string) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount")
	}

	transaction := Transaction{
		Amount:   amount,
		Type:     transType,
		Category: category,
		Date:     date,
		Note:     note,
	}

	t.Transactions = append(t.Transactions, transaction)

	if transType == "income" {
		t.Balance += amount
	} else if transType == "expense" {
		t.Balance -= amount
	} else {
		return fmt.Errorf("invalid transaction type")
	}
	return nil
}

func (t *Tracker) GetBalance()  {
	fmt.Printf("Balance: %.2f\n", t.Balance)
}

func (t *Tracker) GetTransactions(category string) {
	fmt.Println("Transactions:")
	for _, trans := range t.Transactions {
		if category == "" || category == trans.Category {
			fmt.Printf("Type: %s, Amount: %.2f, Date: %s, Note: %s\n", trans.Type, trans.Amount, trans.Date.Format("2006-01-02"), trans.Note)
		}
	}
}

func addTransaction(tracker *Tracker) {
	var amount float32
	var transType, category, note string

	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)

	fmt.Print("Enter transaction type (income/expense): ")
	fmt.Scanln(&transType)

	fmt.Print("Enter category: ")
	fmt.Scanln(&category)

	fmt.Print("Enter note: ")
	fmt.Scanln(&note)

	err := tracker.AddTransaction(amount, time.Now(), note, category, transType)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Transaction added successfully")
	}
}

func main() {
	var input int
	tracker := Tracker{}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add transaction")
		fmt.Println("2. Get transactions")
		fmt.Println("3. Get balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scanln(&input)
			switch input {
				case 1:
					addTransaction(&tracker)
				case 2:
					fmt.Print("Enter category to filter by (or press Enter for all): ")
					var category string
					fmt.Scanln(&category)
					tracker.GetTransactions(category)
				case 3:
					tracker.GetBalance()
				case 4:
					fmt.Println("Exiting...")
					return
				default:
					fmt.Println("Invalid choice, please try again.")
			}

		 }
	}