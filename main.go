package main

import (
	"log"
	"myapi/controllers"
	"myapi/routers"
	"myapi/services"
	"myapi/storage"
	"net/http"
)

func main() {
    userStorage := storage.NewUserStorage()

	userService := &services.UserService{Storage: userStorage}

    controllers.InitUserService(userService)

    router := routers.SetupRouter(userStorage)
    log.Fatal(http.ListenAndServe(":8080", router))
}
