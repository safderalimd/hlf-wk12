rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile CCDOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile CCDOrgChannel -outputCreateChannelTx ./config/ccdchannel.tx -channelID ccdchannel
