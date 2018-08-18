# Hyperledger core example

Inspired by https://github.com/hyperledger/fabric-test and https://github.com/hyperledger/fabric-samples

## Prerequisites

- Install [Docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/)


## Start hyperledger network

- `./start.sh`

## CLI

Connect to CLI with command: `docker-compose -f docker-compose-cli.yaml exec cli /bin/bash`

### Init channel and chaincode

- `./scripts/init.sh`

### Create user

- `peer chaincode invoke -o orderer1.wayne-entreprises.com:7050 -C $CHANNEL_NAME -n user -c '{"Args":["create", "2", "Bruce Wayne"]}' --tls --cafile $CA_FILE`

### Get user (using another peer and another orderer)

- `CORE_PEER_ADDRESS=peer1.asd.wayne-entreprises.com:7051 peer chaincode invoke -o orderer2.wayne-entreprises.com:7050 -C $CHANNEL_NAME -n user -c '{"Args":["get", "2"]}' --tls --cafile $CA_FILE`
