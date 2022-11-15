#!/bin/bash
#
# SPDX-License-Identifier: Apache-2.0
rm /home/ant0/Desktop/IOTPORJ/fabric-samples/commercial-paper/organization/digibank/identity/user/id/wallet/id.id
rm /home/ant0/Desktop/IOTPORJ/fabric-samples/commercial-paper/organization/magnetocorp/identity/user/id/wallet/id.id


function _exit(){
    printf "Exiting:%s\n" "$1"
    exit -1
}

# Exit on first error, print all commands.
set -ev
set -o pipefail

# Where am I?
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export FABRIC_CFG_PATH="${DIR}/../config"

cd "${DIR}/../test-network/"

docker kill cliDigiBank cliMagnetoCorp logspout || true
./network.sh down

# remove any stopped containers
docker rm $(docker ps -aq)

