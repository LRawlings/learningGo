package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		//fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s got %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		wallet.Withdraw(Bitcoin(6))

		got := wallet.Balance()

		want := Bitcoin(4)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
