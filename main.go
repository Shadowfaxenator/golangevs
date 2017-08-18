package main

import (
	"fmt"

	"github.com/shadowfaxenator/golangevc/agregate"
)

func main() {
	walletRepo := []agregate.Event{
		NewWalletCreatedEvent(123, "www2", 100),
		NewFundsAddedEvent(123, 5),
	}
	w := NewWalletAgregate(123, walletRepo)

	//w := NewWallet(walletAgregate)
	fmt.Println(w)
	w.AddFunds(60)
	fmt.Println(w)
}
