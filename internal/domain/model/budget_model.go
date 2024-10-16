package model

import (
	decimal "fintrack/internal/pkg/decimal_clone"
	"time"
)

type Budget struct {
	Balance       Amount `json:"balance"`
	Daily         Amount `json:"daily"`
	RemainingDays int    `json:"remainingDays"`
}

type Amount = decimal.Decimal

var Time = func() time.Time {
	return time.Now()
}

// ComputeBalance computes the overall balance over all transactions.
func ComputeBalance(transactions Transactions) Amount {
	amount := decimal.NewFromFloat(0)
	for _, t := range transactions {
		amount = amount.Add(t.Amount)
	}
	return amount
}

// ComputeBudget computes the remaining budget for the whole month and per day.
//
// The returned tuple contains (month, day) budget.
func ComputeBudget(transactions Transactions) Budget {
	balance := ComputeBalance(transactions).Round(2)
	remainingDays := getRemainingDays()
	dailyBudget := computeDailyBudget(balance, remainingDays).Round(2)
	return Budget{balance, dailyBudget, remainingDays}
}

// getRemainingDays computes the number of remaining days in the current month.
func getRemainingDays() int {
	//Ignore leap years for now.
	now := Time()
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, month, day := now.Date()
	location := now.Location()
	_, _, today := time.Date(year, month, day, 0, 0, 0, 0, location).Date()
	lastDay := days[month-1]
	remainingDays := lastDay - today + 1
	return remainingDays
}

func computeDailyBudget(balance Amount, days int) Amount {
	daysDecimal := decimal.NewFromFloat(float64(days))
	dailyAmount := balance.Div(daysDecimal)
	return dailyAmount
}
