echo "Setting up the network.."

echo "Creating channel genesis block.."

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=BankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Manufacturer.ccd.com/users/Admin@Manufacturer.ccd.com/msp" -e "CORE_PEER_ADDRESS=peer0.Manufacturer.ccd.com:7051" cli peer channel create -o orderer.ccd.com:7050 -c tfbcchannel -f /etc/hyperledger/configtx/tfbcchannel.tx


sleep 5

echo "Channel genesis block created."

echo "peer0.Manufacturer.ccd.com joining the channel..."
# Join peer0.Manufacturer.ccd.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=BankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Manufacturer.ccd.com/users/Admin@Manufacturer.ccd.com/msp" -e "CORE_PEER_ADDRESS=peer0.Manufacturer.ccd.com:7051" cli peer channel join -b tfbcchannel.block

echo "peer0.Manufacturer.ccd.com joined the channel"

echo "peer0.Distributor.ccd.com joining the channel..."

# Join peer0.Distributor.ccd.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=BuyerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Distributor.ccd.com/users/Admin@Distributor.ccd.com/msp" -e "CORE_PEER_ADDRESS=peer0.Distributor.ccd.com:7051" cli peer channel join -b tfbcchannel.block

echo "peer0.Distributor.ccd.com joined the channel"

echo "peer0.Dealer.ccd.com joining the channel..."
# Join peer0.Dealer.ccd.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=SellerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Dealer.ccd.com/users/Admin@Dealer.ccd.com/msp" -e "CORE_PEER_ADDRESS=peer0.Dealer.ccd.com:7051" cli peer channel join -b tfbcchannel.block
sleep 5

echo "peer0.Dealer.ccd.com joined the channel"

echo "Installing ccd chaincode to peer0.Manufacturer.ccd.com..."

# install chaincode
# Install code on Manufacturer peer
docker exec -e "CORE_PEER_LOCALMSPID=BankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Manufacturer.ccd.com/users/Admin@Manufacturer.ccd.com/msp" -e "CORE_PEER_ADDRESS=peer0.Manufacturer.ccd.com:7051" cli peer chaincode install -n tfbccc -v 1.0 -p github.com/ccd/go -l golang

echo "Installed ccd chaincode to peer0.Manufacturer.ccd.com"

echo "Installing ccd chaincode to peer0.Distributor.ccd.com...."

# Install code on Distributor peer
docker exec -e "CORE_PEER_LOCALMSPID=BuyerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Distributor.ccd.com/users/Admin@Distributor.ccd.com/msp" -e "CORE_PEER_ADDRESS=peer0.Distributor.ccd.com:7051" cli peer chaincode install -n tfbccc -v 1.0 -p github.com/ccd/go -l golang

echo "Installed ccd chaincode to peer0.Distributor.ccd.com"

echo "Installing ccd chaincode to peer0.Dealer.ccd.com..."
# Install code on Dealer peer
docker exec -e "CORE_PEER_LOCALMSPID=SellerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Dealer.ccd.com/users/Admin@Dealer.ccd.com/msp" -e "CORE_PEER_ADDRESS=peer0.Dealer.ccd.com:7051" cli peer chaincode install -n tfbccc -v 1.0 -p github.com/ccd/go -l golang

sleep 5

echo "Installed ccd chaincode to peer0.Distributor.ccd.com"

echo "Instantiating ccd chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=BankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/Manufacturer.ccd.com/users/Admin@Manufacturer.ccd.com/msp" -e "CORE_PEER_ADDRESS=peer0.Manufacturer.ccd.com:7051" cli peer chaincode instantiate -o orderer.ccd.com:7050 -C tfbcchannel -n tfbccc -l golang -v 1.0 -c '{"Args":[""]}' -P "OR ('BankMSP.member','BuyerMSP.member','SellerMSP.member')"

echo "Instantiated ccd chaincode."

echo "Following is the docker network....."

docker ps