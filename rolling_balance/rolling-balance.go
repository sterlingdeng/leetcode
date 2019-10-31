package main

/*

	given an input json payload in the format of
	{
		balance: 100.98,
		transactions: [
			{
				amount: -2.34,
				transaction_date: "2002-10-02T10:00:00-05:00"
			},
			{
				amount: 10.13,
				transaction_date: "2002-10-04T10:00:00-05:00"
			},
			{
				amount: -20.14,
				transaction_date: "2002-10-13T10:00:00-05:00"
			}
		]
	}

	return the daily balances
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"time"
)

type RollingBalance struct {
	Balances []dailyBalance
}

type dailyBalance struct {
	Date   time.Time
	Amount float64
}

type TransactionData struct {
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Amount          float64    `json:"amount"`
	TransactionDate *time.Time `json:"transaction_date"`
}

// parses the json input and returns TransactionData with transactions sorted by date, from ealiest
// to latest
func getTransactionData(in string) TransactionData {
	bytes := []byte(in)
	td := TransactionData{}

	err := json.Unmarshal(bytes, &td)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(td.Transactions, func(i, j int) bool {
		return td.Transactions[i].TransactionDate.Before(*td.Transactions[j].TransactionDate)
	})

	return td
}

func GetRollingBalance(in string) RollingBalance {
	td := getTransactionData(in)

	// txn map is a hash map of dates and the sum of all the transactions incurred for that day
	txnMap := make(map[time.Time]float64)

	// modify the dates, so that the transactions occur at 00:00:00:00 UTC
	for _, txn := range td.Transactions {
		t := txn.TransactionDate
		// change the date
		*txn.TransactionDate = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
		// add transaction amount to the txnMap
		if _, exists := txnMap[*txn.TransactionDate]; !exists {
			txnMap[*txn.TransactionDate] = txn.Amount
		} else {
			txnMap[*txn.TransactionDate] += txn.Amount
		}
	}


	// instantiate variables first
	balance := td.Balance
	startDate := td.Transactions[0].TransactionDate.AddDate(0, 0, -1)
	endDate := td.Transactions[len(td.Transactions)-1].TransactionDate.AddDate(0, 0, 1)
	currDate := startDate
	output := RollingBalance{}
	for !currDate.Equal(endDate) {
		if amount, exist := txnMap[currDate]; exist {
			balance += amount
		}
		output.Balances = append(output.Balances, dailyBalance{
			Date:   currDate,
			Amount: balance,
		})
		currDate = currDate.AddDate(0, 0, 1)
	}

	return output
}

func main() {
	data := `{
		"balance": 100.98,
		"transactions": [
			{
				"amount": -2.34,
				"transaction_date": "2002-10-02T10:00:00-05:00"
			},
			{
				"amount": 10.13,
				"transaction_date": "2002-10-04T10:00:00-05:00"
			},
			{
				"amount": -20.14,
				"transaction_date": "2002-10-13T10:00:00-05:00"
			}
		]
	}`
	rb := GetRollingBalance(data)
	fmt.Println(rb)
	for _, db := range rb.Balances {
		fmt.Printf("Day: %v. Balance: %f\n", db.Date, db.Amount)
	}
}
