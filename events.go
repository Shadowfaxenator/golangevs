package main

import (
	"github.com/shadowfaxenator/golangevc/agregate"
)

type WalletCreated struct {
	*agregate.BasicEvent
	BaseDeposit int
	Name        string
}

func (b *WalletCreated) String() string {

	return "WalletCreated"
}

type FundsAdded struct {
	*agregate.BasicEvent
	Amount int
}

func (b *FundsAdded) String() string {

	return "FoundsAddded"
}

func (e *WalletCreated) Apply(state agregate.Agregate) {
	state.(*WalletAgregate).Id = e.AgregateId
	state.(*WalletAgregate).Name = e.Name
	state.(*WalletAgregate).Balance = e.BaseDeposit
}

func (e *FundsAdded) Apply(state agregate.Agregate) {
	state.(*WalletAgregate).Balance += e.Amount
}

func NewFundsAddedEvent(id agregate.AgregateId, amount int) (event *FundsAdded) {
	event = &FundsAdded{}
	event.BasicEvent = &agregate.BasicEvent{}
	event.AgregateId = id
	event.Amount = amount
	return
}

func NewWalletCreatedEvent(id agregate.AgregateId, name string, balance int) (event *WalletCreated) {
	event = &WalletCreated{}
	event.BasicEvent = &agregate.BasicEvent{}
	event.AgregateId = id
	event.Name = name
	event.BaseDeposit = balance
	return
}
