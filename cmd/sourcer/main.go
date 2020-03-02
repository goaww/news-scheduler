package main

import (
	"github.com/goaww/news-scheduler/internal/scheduler"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 3 && args[1] == "i" {
		insertUrl(&args[2])
	} else if len(args) == 2 && args[1] == "l" {
		getUrl()
	} else {
		log.Fatal("invalid param, use 'i' for insert new url and 'l' for list ")
	}
}

func insertUrl(s *string) {
	db := scheduler.NewDB(*scheduler.NewConf())
	db.Connect()
	defer db.Close()
	service := scheduler.NewUrlItemServiceImpl(db.Con)
	service.Add(*scheduler.NewItem(s))
	log.Println("insert url successful", *s)
}

func getUrl() {
	db := scheduler.NewDB(*scheduler.NewConf())
	db.Connect()
	defer db.Close()
	service := scheduler.NewUrlItemServiceImpl(db.Con)
	for item := range service.Get() {
		log.Println(item.V.(*scheduler.Item).Url)
	}
}
