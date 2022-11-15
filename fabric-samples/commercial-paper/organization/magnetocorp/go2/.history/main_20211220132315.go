package main

import (
	"fmt"
	f "go2/functions"
	"io/ioutil"
	"log"
	"os"

	//"github.com/gorilla/mux"

	"net/http"
)

func main() {
	/*
		r := mux.NewRouter()

		httpsMux := http.NewServeMux()

		httpsMux.Handle("/", http.HandlerFunc(infoTeam))
		httpsMux.Handle("/prova", http.HandlerFunc(prova))

		http.ListenAndServeTLS(":8081", "localhost.crt",
			"localhost.key", httpsMux)

		http.Handle("/", r)
	*/
	idUser := os.Args[1]
	result, err = f.Inizialize(idUser)
	f.QueryAll(idUser)
	f.Create(idUser, "Asset2", "32.5", "80.5", "tagrotto")
	f.GetItemFromId(idUser, "Asset2")
	/*
		f.Create(idUser, "ERICSSON", "A1234", "COSTUMID", "DESCR", "K", "TYPEP", "SWLT", "1-1-2020", "1-1-2021", "3", "1", "ERICSSON")
		f.QueryAll(idUser)
		f.GetItemFromId(idUser, "A1234")
		f.Transfer(idUser, "A1234", "TIM")

	*/
	fmt.Println("Main Page")
}

//--------------------------------------------------------------------------------------------

func infoTeam(rw http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(rw, "hello")

	fmt.Printf("Ho ricevuto\n")
	fmt.Printf("%s\n", reqBody)
	idUser := os.Args[1]
	f.Inizialize(idUser)
	f.QueryAll(idUser)
}

func prova(rw http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(rw, "hello")

	fmt.Printf("Ho ricevuto in prova\n")
	fmt.Printf("%s\n", reqBody)
}
