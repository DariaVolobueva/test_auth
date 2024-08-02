// routers/router.go
package routers

import (
	"myapi/controllers"
	"myapi/pkg/authentication_check"
	"myapi/storage"
	"net/http"
)

func SetupRouter(userStorage *storage.UserStorage) *http.ServeMux {
    router := http.NewServeMux()
    authHandler := controllers.AuthHandler{UserStorage: userStorage}

    router.HandleFunc("/login", authHandler.Login)

    authenticatedRouter := http.NewServeMux()
    authenticatedRouter.HandleFunc("/users", controllers.HandleUsers)
    authenticatedRouter.HandleFunc("/users/", controllers.HandleUser)

    router.Handle("/users", authentication_check.Authenticate(authenticatedRouter))
    router.Handle("/users/", authentication_check.Authenticate(authenticatedRouter))

    return router
}
