package main

import (
	"log"
	"net/http"

	"github.com/chrispeterjeyaraj/track-my-wealth/server/db"
	"github.com/chrispeterjeyaraj/track-my-wealth/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/userpref", h.GetAllUserPrefs).Methods(http.MethodGet)
	// router.HandleFunc("/userpref/{id}", h.GetUserPref).Methods(http.MethodGet)
	// router.HandleFunc("/userpref", h.AddUserPref).Methods(http.MethodPost)
	// router.HandleFunc("/userpref/{id}", h.UpdateUserPref).Methods(http.MethodPut)
	// router.HandleFunc("/userpref/{id}", h.DeleteUserPref).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
