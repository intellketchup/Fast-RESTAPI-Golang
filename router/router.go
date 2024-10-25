// router/router.go
package router

import (
    "github.com/gorilla/mux"
    "RESTAPI/handlers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/items", handlers.GetItems).Methods("GET")
    r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
    r.HandleFunc("/items/{id}", handlers.GetItemByID).Methods("GET")
    r.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

    return r
}
