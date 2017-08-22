package main

import (
	"fmt"

	"github.com/shadowfaxenator/golangevc/agregate"
)

type WalletCreated struct {
	agregate.BasicEvent `bson:",inline"`
	BaseDeposit         int
	Name                string
}

func (b *WalletCreated) String() string {

	return "WalletCreated"
}

type FundsAdded struct {
	agregate.BasicEvent `bson:",inline"`
	Amount              int
}

func (b *FundsAdded) String() string {

	return "FoundsAddded"
}

func (e *WalletCreated) Apply(state agregate.Agregate) {
	if state, ok := state.(*WalletAgregate); ok {

		state.Lock()
		defer state.Unlock()
		state.Name = e.Name
		state.Balance = e.BaseDeposit
		state.Status = "ACTIVE"
	}

}

func (e *FundsAdded) Apply(state agregate.Agregate) {
	if state, ok := state.(*WalletAgregate); ok {
		state.Lock()
		defer state.Unlock()
		state.Balance += e.Amount
	}

}

func NewFundsAddedEvent(amount int) *FundsAdded {
	event := &FundsAdded{
		Amount: amount,
	}
	t := agregate.EventType(fmt.Sprintf("%T\n", event))
	event.BasicEvent = agregate.NewBasicEvent(t)

	return event
}

func NewWalletCreatedEvent(name string, balance int) *WalletCreated {

	event := &WalletCreated{

		Name:        name,
		BaseDeposit: balance,
	}
	t := agregate.EventType(fmt.Sprintf("%T\n", event))
	event.BasicEvent = agregate.NewBasicEvent(t)
	return event
}
