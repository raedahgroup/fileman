package main

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/raedahgroup/fileman/config"
	ctl "github.com/raedahgroup/fileman/http"
	"github.com/raedahgroup/fileman/storage/bolt"
	"log"
	"net"
	"net/http"
)

func main() {
	// CONFIG
	err := config.Load("config/config.yaml")
	if err != nil {
		fmt.Println("Error: Failed to load configuration")
	}
	checkErr(err);
	db, err := storm.Open(config.State.DatabasePath)
	checkErr(err)
	defer db.Close()
	store, err := bolt.NewStorage(db)
	adr := "127.0.0.1:"   + config.State.Port
	var listener net.Listener
	listener, err = net.Listen("tcp", adr)
	checkErr(err)
	handler, err := ctl.NewHandler(store, config.State)
	log.Println("Listening on", listener.Addr().String())
	if err := http.Serve(listener, handler); err != nil {
		log.Fatal(err)
	}
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
