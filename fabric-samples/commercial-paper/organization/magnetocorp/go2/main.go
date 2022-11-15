package main

import (
	"encoding/json"
	"fmt"
	f "go2/functions"
	"log"
	"net/http"
	"os"
)

type UserPost struct {
	ID string `json:"id"`
}

type Asset struct {
	ID          string `json:"ID"`
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
	Owner       string `json:"owner"`
}

type idAsset struct {
	ID string `json:"ID"`
}

var idUser = os.Args[1]

func main() {
	initLedgerR()
	httpsMux := http.NewServeMux()
	httpsMux.Handle("/update", http.HandlerFunc(updateItem))
	httpsMux.Handle("/read", http.HandlerFunc(getIDItem))
	http.ListenAndServeTLS(":8081", "localhost.crt", "localhost.key", httpsMux)
	fmt.Println("Main Page")
}

func initLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := f.Inizialize(idUser)
	if err != nil {
		log.Println("errore:", err)
	}

	w.Write([]byte("Succesful"))
}
func initLedgerR() {
	_, err := f.Inizialize(idUser)
	if err != nil {
		log.Println("errore:", err)
	}

}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("UPDATE START")
	var appAsset Asset
	json.NewDecoder(r.Body).Decode(&appAsset)
	res, err := f.Update(idUser, appAsset.ID, appAsset.Temperature, appAsset.Humidity, appAsset.Owner)
	if err != nil {
		log.Println("errore:", err)
	}
	log.Println(string(res))

	w.Write(res)

}
func getIDItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("GET START")
	var newIdAsset idAsset
	json.NewDecoder(r.Body).Decode(&newIdAsset)
	res, err := f.GetItemFromId(idUser, newIdAsset.ID)
	if err != nil {
		log.Println("errore:", err)
	}
	log.Println(string(res))

	w.Write(res)
}

//--------------------------------------------------------------------------------------------
/*
func createItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	log.Println("CREATE START")

	var newAsset Asset

	json.NewDecoder(r.Body).Decode(&newAsset)
	//prendi valori dal post e passali alla funzione
	assets = append(assets, newAsset)
	str, _ := json.MarshalIndent(newAsset, "", "")
	log.Println("ResultJSON: ", string(str))
	json.NewEncoder(w).Encode(&newAsset)

	_, err := f.Create(os.Args[1], newAsset.ID, newAsset.Temperature, newAsset.Humidity, newAsset.Owner)
	if err != nil {
		log.Println("CREATE FAIL")
	}
	log.Println("CREATE OK")

}
*/

/*
//POST LOGIC ok
func transferItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("TRANSFER START")
	params := mux.Vars(r)
	for index, item := range assets {
		if item.ID == params["id"] {
			assets = append(assets[:index], assets[index+1:]...)

			var newAsset Asset
			_ = json.NewDecoder(r.Body).Decode(&newAsset)
			newAsset.ID = params["id"]

			_, err := f.Transfer(os.Args[1], newAsset.ID, newAsset.Owner)
			if err != nil {
				log.Println("TRANSFER FAIL")

			}
			str, _ := json.MarshalIndent(newAsset, "", "")
			log.Println("ResultJSON: ", string(str))

			assets = append(assets, newAsset)
			json.NewEncoder(w).Encode(&newAsset)
			log.Println("TRANSFER OK")
			return
		}
	}
	json.NewEncoder(w).Encode(assets)

}
*/

/*
//GET LOGIC
func getAllItem(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL START ")

	w.Header().Set("Content-Type", "application/json")

	f.QueryAll(os.Args[1])

	json.NewEncoder(w).Encode(assets)

	//log.Println("RES", string(result))

	str, _ := json.MarshalIndent(assets, "", "")
	log.Println("ResultJSON: ", string(str))

}
*/
/*
//POST LOGIC
func addUser(w http.ResponseWriter, r *http.Request) {

	log.Println("ADD USER START")
	w.Header().Set("Content-Type", "application/json")

	var userpost UserPost

	_ = json.NewDecoder(r.Body).Decode(&userpost)
	f.AddToWallet(userpost.ID)
	userposts = append(userposts, userpost)
	json.NewEncoder(w).Encode(&userposts)

	log.Printf("USER ID: %v ADDED", userpost.ID)

}

func removeUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("REMOVE USER START")

	params := mux.Vars(r)
	for index, item := range userposts {
		if item.ID == params["id"] {
			id := params["id"]
			userposts = append(userposts[:index], userposts[index+1:]...)
			f.RemoveFromWallet(id)
			log.Printf("USER ID: %v REMOVED", id)
			break
		}
	}
	json.NewEncoder(w).Encode(userposts)
}

func getAllWallet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	f.ListAll()
	json.NewEncoder(w).Encode(userposts)

	str, _ := json.MarshalIndent(userposts, "", "")
	log.Println("ResultJSON ID WALLET: ", string(str))
}
*/
