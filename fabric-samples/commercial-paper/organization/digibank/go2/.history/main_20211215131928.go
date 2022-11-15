package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	httpsMux := http.NewServeMux()

	httpsMux.Handle("/", http.HandlerFunc(infoTeam))

	go http.ListenAndServeTLS(":8488", "cert.pem",
		"key.pem", httpsMux)

	http.Handle("/", r)

	f.Create(idUser, "ERICSSON", "A1234", "COSTUMID", "DESCR", "K", "TYPEP", "SWLT", "1-1-2020", "1-1-2021", "3", "1", "ERICSSON")
	f.QueryAll(idUser)
	f.GetItemFromId(idUser, "A1234")
	f.Transfer(idUser, "A1234", "TIM")
	f.QueryAll(idUser)

	fmt.Println("Main Page")
}
