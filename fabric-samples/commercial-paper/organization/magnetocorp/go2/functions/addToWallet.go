package functions

import (
	"fmt"
	"log"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func AddToWallet(userId string) error {

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
	if !wallet.Exists(userId) {
		err := populateWallet(wallet, userId)
		if err != nil {
			return fmt.Errorf("failed to populate wallet contents: %v", err)
		}
	}
	_, err = wallet.Get(userId)
	if err != nil {
		return fmt.Errorf("failed to get wallet contents: %v", err)
	}

	return err

}

// func populateWallet(wallet *gateway.Wallet, id string) error {
// 	log.Println("============ Populating wallet ============")
// 	credPath := filepath.Join(
// 		"..",
// 		"..",
// 		"..",
// 		"..",
// 		"test-network",
// 		"organizations",
// 		"peerOrganizations",
// 		"org1.example.com",
// 		"users",
// 		"User1@org1.example.com",
// 		"msp",
// 	)

// 	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
// 	// read the certificate pem
// 	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
// 	if err != nil {
// 		return err
// 	}

// 	keyDir := filepath.Join(credPath, "keystore")
// 	// there's a single file in this dir containing the private key
// 	files, err := ioutil.ReadDir(keyDir)
// 	if err != nil {
// 		return err
// 	}
// 	// for _, f := range files {
// 	// 	fmt.Println(f.Name())
// 	// }
// 	// fmt.Println(len(files))
// 	// fmt.Println(files)
// 	// fmt.Println(files[0].Name())

// 	if len(files) != 1 {
// 		return fmt.Errorf("keystore folder should have contain one file")
// 	}
// 	keyPath := filepath.Join(keyDir, files[0].Name())
// 	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
// 	if err != nil {
// 		return err
// 	}

// 	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

// 	return wallet.Put(id, identity)
// }
