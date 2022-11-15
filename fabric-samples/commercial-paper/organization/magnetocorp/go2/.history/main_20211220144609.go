package main

import (
	"encoding/json"
	"fmt"
	f "go2/functions"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

var assets []Asset = []Asset{}
var userposts []UserPost

func main() {

	r := mux.NewRouter()

	httpsMux := http.NewServeMux()

	httpsMux.Handle("/", http.HandlerFunc(createItem))
	httpsMux.Handle("/prova", http.HandlerFunc(getAllItem))

	idUser := os.Args[1]
	res, err := f.Inizialize(idUser)
	if err != nil {
		log.Println("errore:", err)
	}
	log.Println(string(res))
	log.Println("INIT OK")
	//http.ListenAndServeTLS(":8081", "localhost.crt","localhost.key", httpsMux)
	http.ListenAndServe(":8081", httpsMux)
	http.Handle("/", r)
	/*


		log.Println(string(res))
		log.Println("INIT OK")
		res, err = f.GetItemFromId(idUser, "Asset2")
		if err != nil {
			log.Println("errore:", err)
		}
		log.Println(string(res))
		log.Println("Read OK")
		//f.QueryAll(idUser)
		//f.Create(idUser, "Asset2", "32.5", "80.5", "tagrotto")
		//

			f.Create(idUser, "ERICSSON", "A1234", "COSTUMID", "DESCR", "K", "TYPEP", "SWLT", "1-1-2020", "1-1-2021", "3", "1", "ERICSSON")
			f.QueryAll(idUser)
			f.GetItemFromId(idUser, "A1234")
			f.Transfer(idUser, "A1234", "TIM")

	*/
	fmt.Println("Main Page")
}

//--------------------------------------------------------------------------------------------

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

/*
//POST LOGIC ok
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("TRANSFER START")
	params := mux.Vars(r)
	for _, item := range assets {
		if item.ID == params["id"] {
			//assets = append(assets[:index], assets[index+1:]...)

			var newAsset Asset
			_ = json.NewDecoder(r.Body).Decode(&newAsset)
			newAsset.ID = params["id"]

			_, err := f.Update(os.Args[1], newAsset.Issuer, newAsset.ID,
				newAsset.CustomerID, newAsset.Description,
				newAsset.Key, newAsset.ProductType,
				newAsset.SwltID, newAsset.ActivationTime,
				newAsset.ExpirationTime, newAsset.TotalCapacity,
				newAsset.UsedCapacity, newAsset.Owner)

			if err != nil {
				log.Println("UPDATE FAIL")

			}
			str, _ := json.MarshalIndent(newAsset, "", "")
			log.Println("ResultJSON: ", string(str))

			assets = append(assets, newAsset)
			json.NewEncoder(w).Encode(&newAsset)
			log.Println("UPDATE OK")
			return
		}
	}
	json.NewEncoder(w).Encode(assets)

}

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
//GET LOGIC
func getIDItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("GET START")
	params := mux.Vars(r)
	for _, item := range assets {
		if item.ID == params["id"] {
			id := params["id"]

			f.GetItemFromId(os.Args[1], id)
			json.NewEncoder(w).Encode(item)
			str, _ := json.MarshalIndent(item, "", "")
			log.Println("GET ResultJSON: ", string(str))
			break
		}
		id := params["id"]
		log.Println("GET FAIL FOR ID :", id)
		return
	}
	json.NewEncoder(w).Encode(&Asset{})
}

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
