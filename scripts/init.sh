#!/bin/bash

echo "Creating channel"
peer channel create -o orderer1.wayne-entreprises.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile $CA_FILE

peers=( peer0.asd.wayne-entreprises.com:7051 peer1.asd.wayne-entreprises.com:7051 peer2.asd.wayne-entreprises.com:7051 )
for peer in "${peers[@]}"; do
  echo "Initializing peer $peer"
  CORE_PEER_ADDRESS=$peer peer channel join -b $CHANNEL_NAME.block
  CORE_PEER_ADDRESS=$peer peer chaincode install -n user -v 1.0 -p chaincode/user
done

peer chaincode instantiate -o orderer1.wayne-entreprises.com:7050 -C $CHANNEL_NAME -n user -v 1.0 -c '{"Args":["init"]}' --tls --cafile $CA_FILE
