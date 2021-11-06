package main

import (
	"log"

	"geek/errors/dao"
	"geek/errors/service"
)

func main() {
	err := dao.Init("localhost", 3306, "test", "root", "123456")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	account := "zsc"
	user, err := service.CreateUser(account, "zsc@yeah.net")
	if err != nil {
		log.Printf("create %s failed: %+v\n", account, err)
		return
	}
	log.Printf("%+v\n", user)
}
