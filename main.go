package main

import (
	"encoding/gob"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/shadowfaxenator/golangevc/agregate"
)

func init() {
	gob.Register(&WalletCreated{})
	gob.Register(&FundsAdded{})

}

func main() {
	//id := agregate.NewAgregateID()
	id := agregate.AgregateId(bson.ObjectIdHex("599c56591b6ecf61803396c5"))
	//fmt.Println(bson.ObjectId(id))
	ag := NewWalletAgregate(id)
	fmt.Println("Balance:", ag.Balance)
	// if err := ag.CreateNewWallet(10, "MyWal"); err != nil {
	// 	log.Fatalln(err)
	// }
	if err := ag.AddFunds(5); err != nil {
		log.Fatalln(err)
	}
	//	time.Sleep(20 * time.Second)
	if err := agregate.Commit(ag); err != nil {
		log.Fatalln(err)
	}

	// var e bson.M
	// i := col.Find(nil).Iter()
	// for i.Next(&e) {
	// 	fmt.Println(e)

	// }
	// if err := i.Close(); err != nil {
	// 	panic(err)
	// }
	// walletRepo := []agregate.Event{
	// 	NewWalletCreatedEvent(123, "www2", 100),
	// 	NewFundsAddedEvent(123, 5),
	// }
	// w := NewWalletAgregate(123, walletRepo)

	// fmt.Println(w)
	// w.AddFunds(60)
	// fmt.Println(w)

}
