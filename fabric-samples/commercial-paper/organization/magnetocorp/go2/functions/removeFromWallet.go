package functions

import (
	"fmt"
	"log"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func RemoveFromWallet(userId string) error {

	path := "../identity/user/" + userId
	path2 := path + "/wallet"
	fmt.Println("Wallet path: ", path2)

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		return fmt.Errorf("error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet(path2)
	if err != nil {
		log.Fatal(err)
	}
	if wallet.Exists(userId) {
		wallet.Remove(userId)
		err = os.RemoveAll(path)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		return fmt.Errorf("wallet %v does't exist! ", err)
	}
	return err
}
