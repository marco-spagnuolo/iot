#!/bin/bash
#
# SPDX-License-Identifier: Apache-2.0


peer lifecycle chaincode commit -o localhost:7050 \
                                --peerAddresses localhost:7051 --tlsRootCertFiles "${PEER0_ORG1_CA}" \
                                --peerAddresses localhost:9051 --tlsRootCertFiles "${PEER0_ORG2_CA}" \
                                --ordererTLSHostnameOverride orderer.example.com \
                                --channelID mychannel --name papercontract -v 0 \
                                --sequence 1 \
                                --tls --cafile "$ORDERER_CA" --waitForEvent