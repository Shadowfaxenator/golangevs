package main

import (
	"fmt"
	"sync"

	"github.com/shadowfaxenator/golangevc/agregate"
)

type WalletAgregate struct {
	*agregate.BasicAgregate
	Balance int
	Name    string
}

func (w *WalletAgregate) String() string {
	format := `
	"Name: %s"
	"Balance: %d"
	"Version: %d"
	"ID: %v"
	"PendingEvents: %v"
	`
	return fmt.Sprintf(format, w.Name, w.Balance, w.ExpectedVersion, w.Id, w.Changes)
}

func (state *WalletAgregate) AddFunds(amount int) {

	agregate.AddChange(state, NewFundsAddedEvent(state.Id, amount))
}

func NewWalletAgregate(id agregate.AgregateId, e []agregate.Event) *WalletAgregate {

	ag := &WalletAgregate{BasicAgregate: &agregate.BasicAgregate{Id: id, Mu: new(sync.RWMutex)}}

	agregate.Construct(ag, e)
	return ag
}
