#!/bin/bash

docker-compose -f docker-compose-cli.yaml down
docker volume rm -f $(docker volume ls -q)

rm -rf channel-artifacts/*
rm -rf crypto-config

mkdir -p channel-artifacts

# generate certificate
./bin/cryptogen generate --config=crypto-config.yaml

# create first block
./bin/configtxgen -profile OrdererGenesis -outputBlock ./channel-artifacts/genesis.block

# create channel(s)
./bin/configtxgen -profile Channel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID asdchannel

cd crypto-config/peerOrganizations/asd.wayne-entreprises.com/ca/
export CERT_FILENAME=$(ls *_sk)
cd ../../../..

docker-compose -f docker-compose-cli.yaml up
