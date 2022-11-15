/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package functions

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func populateWallet(wallet *gateway.Wallet, id string) error {
	log.Println("============== Populating wallet ==============")
	credPath := filepath.Join(
		"..",
		"..",
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)
	log.Println("============== Read the signcert ==============")
	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}
	log.Println("============== Read the keystore ==============")

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}

	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put(id, identity)
}
