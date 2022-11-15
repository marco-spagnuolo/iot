package functions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func AssetExists(idUser string, id string) ([]byte, error) {

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		return nil, fmt.Errorf("error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	path := "../identity/user/" + idUser
	path2 := path + "/wallet"
	log.Println("Wallet path: ", path2)

	wallet, err := gateway.NewFileSystemWallet(path2)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v", err)
	}

	if !wallet.Exists(idUser) {
		log.Println("==============> Oh it's the fist time for you...")

		log.Println("==============> I'm building the wallet...")

		err := populateWallet(wallet, idUser)
		if err != nil {
			return nil, fmt.Errorf("failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org2.example.com",
		"connection-org2.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, idUser),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gateway: %v", err)
	}
	defer log.Printf("==============> Get %v program complete", id)
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		return nil, fmt.Errorf("failed to get network: %v", err)
	}

	contract := network.GetContract("papercontract")

	log.Println("--> Evaluate Transaction: AssetExists:")
	result, err := contract.EvaluateTransaction("AssetExists", id)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))
	return result, err

}
