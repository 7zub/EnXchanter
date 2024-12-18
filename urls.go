package main

import (
	"enchainer/_dev"
	"enchainer/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", _dev.HomePage)
	myRouter.HandleFunc("/articles", _dev.ReturnAllArticles)
	myRouter.HandleFunc("/article", _dev.CreateNewArticle)   //.Methods("POST")
	myRouter.HandleFunc("/article/{id}", _dev.DeleteArticle) //.Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", _dev.ReturnSingleArticle)
	myRouter.HandleFunc("/hotels/", _dev.AllHotels)
	myRouter.HandleFunc("/hotels/export", _dev.ExportHotel)
	myRouter.HandleFunc("/kafka", _dev.Kafkatest)

	myRouter.HandleFunc("/book", views.BookControl)
	myRouter.HandleFunc("/addpair", views.AddPair)
	myRouter.HandleFunc("/deletepair", views.DeletePair)
	myRouter.HandleFunc("/onpair", views.OnPair)
	myRouter.HandleFunc("/offpair", views.OffPair)
	//myRouter.HandleFunc("/ws", views.Ws)
	myRouter.HandleFunc("/trade", views.TradeTaskControl)

	log.Fatal(http.ListenAndServe(":10", myRouter))
}
