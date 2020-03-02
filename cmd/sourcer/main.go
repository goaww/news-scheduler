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
	service := scheduler.NewUrlItemServiceImpl(scheduler.NewConf())
	service.Add(*scheduler.NewItem(s))
	log.Println("insert url successful", *s)
}

func getUrl() {
	service := scheduler.NewUrlItemServiceImpl(scheduler.NewConf())
	for item := range service.Get() {
		log.Println(item.V.(*scheduler.Item).Url)
	}
}
