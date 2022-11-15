#!/bin/bash
#
# SPDX-License-Identifier: Apache-2.0

peer lifecycle chaincode package cp.tar.gz --lang golang --path ./contract-go --label cp_0
peer lifecycle chaincode install cp.tar.gz

export PACKAGE_ID=$(peer lifecycle chaincode queryinstalled --output json | jq -r '.installed_chaincodes[0].package_id')
echo $PACKAGE_ID

peer lifecycle chaincode approveformyorg  --orderer localhost:7050 --ordererTLSHostnameOverride orderer.example.com \
                                          --channelID mychannel  \
                                          --name papercontract  \
                                          -v 0  \
                                          --package-id $PACKAGE_ID \
                                          --sequence 1  \
                                          --tls  \
                                          --cafile "$ORDERER_CA"

peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name papercontract -v 0 --sequence 1
