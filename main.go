package main

import (
	"fmt"
	"my_account/config"
	"my_account/handler"
	"my_account/repository"
	"my_account/service"
	"net/http"
)

func main() {
	config := config.Get()

	repository := repository.NewAccountRepositoryMap()

	service := service.NewAccountService(config.InitialBalanceAmount, config.MinumumBalanceAmount, repository)
	handler := handler.NewAccountHandler(service)

	err := http.ListenAndServe("127.0.0.1:8090", handler)
	if err != nil {
		fmt.Println(err)
	}
}
