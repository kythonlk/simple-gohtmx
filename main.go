package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"property-listing-api/api"

	"github.com/gorilla/mux"
)

func main() {
	api.InitDB()
	migrate := flag.Bool("migrate", false, "migrate database")
	flag.Parse()
	if *migrate {
		api.RunMigrations()
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/register", api.RegisterUser).Methods("POST")
	r.HandleFunc("/login", api.LoginUser).Methods("POST")
	r.HandleFunc("/properties", api.GetProperties).Methods("GET")
	r.Handle("/properties", api.JwtMiddleware(http.HandlerFunc(api.CreateProperty))).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
