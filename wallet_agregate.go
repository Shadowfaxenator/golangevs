package main

import (
	"errors"
	"fmt"

	"log"

	"github.com/shadowfaxenator/golangevc/agregate"
)

type WalletAgregate struct {
	*agregate.BasicAgregate
	Balance int
	Name    string
	Status  string
}

func (w *WalletAgregate) String() string {
	format := `
	"Name: %s"
	"Balance: %d"
	"PendingEvents: %v",
	"Status: %v"
	`
	return fmt.Sprintf(format, w.Name, w.Balance, w.PendingEvents, w.Status)
}

func (state *WalletAgregate) AddFunds(amount int) error {
	if state.Status != "ACTIVE" {

		return errors.New("Can't Add Founds, status is not active")
	}
	agregate.AddChange(state, NewFundsAddedEvent(amount))
	return nil
}

func (state *WalletAgregate) CreateNewWallet(baseDeposit int, name string) error {
	e := NewWalletCreatedEvent(name, baseDeposit)
	if state.Status != "" {
		return errors.New("Wallet is already created")
	}
	agregate.AddChange(state, e)
	return nil
}

func NewWalletAgregate(id agregate.AgregateId) *WalletAgregate {

	ag := &WalletAgregate{}
	ag.BasicAgregate = agregate.NewBasicAgregate(id, agregate.AgregateType("Wallet"))

	err := agregate.Construct(ag)
	if err != nil {
		log.Fatalln(err)
	}
	return ag
}
