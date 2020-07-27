1卸载已有的容器等内容
cd ~/fabric-food/balance-transfer/food_cli
./networkdown.sh 
或
./networkdown.sh couchdb
 
2启动环境
cd ~/fabric-food/balance-transfer/food_cli
docker-compose -f docker-compose-cli.yaml up -d
或启动couchdb
docker-compose -f docker-compose-cli-couch.yaml up -d

3查看cli日志
docker logs -f cli

4访问couchdb
http://192.168.1.103:5984/_utils/

5启动node.js SDK
cd ~/fabric-food/balance-transfer
sudo PORT=4000 node app

5打开新的命令行窗口，启动应用层服务
cd ~/fabric-food/balance-transfer/foodsafety-master
sudo npm start

7打开新的命令行窗口，测试node.js SDK

#CA注册

echo "POST request Enroll on Org1  ..."
echo
ORG1_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Jim&orgName=org1')
echo $ORG1_TOKEN
ORG1_TOKEN=$(echo $ORG1_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "ORG1 token is $ORG1_TOKEN"
echo
echo "POST request Enroll on Org2 ..."
echo
ORG2_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Barry&orgName=org2')
echo $ORG2_TOKEN
ORG2_TOKEN=$(echo $ORG2_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "ORG2 token is $ORG2_TOKEN"
echo
echo
echo "POST request Create channel  ..."
echo
curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"channelName":"mychannel",
	"channelConfigPath":"../food_cli//channel-artifacts/channel.tx"
}'
echo

#加入通道

echo "POST request Join channel on Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer1","peer2"]
}'
echo
echo

echo "POST request Join channel on Org2"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer1","peer2"]
}'
echo
echo

#安装链码
echo "POST Install chaincode on Org1"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer1", "peer2"],
	"chaincodeName":"myccfoodall",
	"chaincodePath":"food_safetyall",
	"chaincodeVersion":"v0"
}'
echo
echo


echo "POST Install chaincode on Org2"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer1","peer2"],
	"chaincodeName":"myccfoodall",
	"chaincodePath":"food_safetyall",
	"chaincodeVersion":"v0"
}'
echo
echo

#实例化链码
echo "POST instantiate chaincode on peer1 of Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"chaincodeName":"myccfoodall",
	"chaincodeVersion":"v0",
	"args":["init"]
}'
echo
echo

#安装链码
echo "POST Install chaincode1 on Org1"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer1", "peer2"],
	"chaincodeName":"myccfoodone",
	"chaincodePath":"food_safetysplit1",
	"chaincodeVersion":"v0"
}'
echo
echo


echo "POST Install chaincode2 on Org2"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer1","peer2"],
	"chaincodeName":"myccfoodtwo",
	"chaincodePath":"food_safetysplit2",
	"chaincodeVersion":"v0"
}'
echo
echo

#实例化链码
echo "POST instantiate chaincode on peer1 of Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"chaincodeName":"myccfoodone",
	"chaincodeVersion":"v0",
	"args":["init"]
}'
echo
echo
echo "POST instantiate chaincode2 on peer1 of Org2"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"chaincodeName":"myccfoodtwo",
	"chaincodeVersion":"v0",
	"args":["init"]
}'
echo
echo

#invoke链码
echo "POST invoke chaincode on peers of Org1 and Org2"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/myccfoodall/invoke \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"fcn":"initLedger",
	"args":["initLeger"]
}')
echo "Transacton ID is $TRX_ID"
echo
echo

echo "POST invoke chaincode on peers of Org1 and Org2"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/myccfoodone/invoke \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
    "peers":["peer1"],
	"fcn":"readAll",
	"args":["DRUG000","DRUG999"]
}'

echo
echo

echo "POST invoke chaincode on peers of Org1 and Org2"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/myccfoodall/invoke \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
    "peers":["peer1"],
	"fcn":"initPerson",
	"args":["FOOD888","张三","男","米业生产","监仓","123"]
}')
echo "Transacton ID is $TRX_ID"
echo
echo

#query链码
echo "POST query chaincode on peers of Org1 and Org2"
echo
curl -s -X POST\
  http://localhost:4000/channels/mychannel/chaincodes/myccfoodall/query \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
    "peers":["peer1"],
	"fcn":"readAll",
	"args":["PERSON000","PERSON999"]
}'
echo "Transacton ID is $TRX_ID"
echo

echo "POST query chaincode on peers of Org1 and Org2"
echo
curl -s -X POST\
  http://localhost:4000/channels/mychannel/chaincodes/myccfoodone/query \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
    "peers":["peer2"],
	"fcn":"readAll",
	"args":["DRUG000","DRUG999"]
}'

echo "Transacton ID is $TRX_ID"
echo

["PERSON000","PERSON999"]
echo
curl -s -X POST\
  http://localhost:4000/channels/mychannel/chaincodes/myccfoodtwo/query \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
    "peers":["peer2"],
	"fcn":"readAll",
	"args":["PRODUCT000","PRODUCT999"]
}'
echo "Transacton ID is $TRX_ID"
echo

#查询区块

echo "GET query Block by blockNumber"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/blocks/1?peer=peer1" \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "GET query Transaction by TransactionID"
echo
curl -s -X GET http://localhost:4000/channels/mychannel/transactions/$TRX_ID?peer=peer1 \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

############################################################################
### TODO: What to pass to fetch the Block information
############################################################################
#echo "GET query Block by Hash"
#echo
#hash=????
#curl -s -X GET \
#  "http://localhost:4000/channels/mychannel/blocks?hash=$hash&peer=peer1" \
#  -H "authorization: Bearer $ORG1_TOKEN" \
#  -H "cache-control: no-cache" \
#  -H "content-type: application/json" \
#  -H "x-access-token: $ORG1_TOKEN"
#echo
#echo

echo "GET query ChainInfo"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel?peer=peer1" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "GET query Installed chaincodes"
echo
curl -s -X GET \
  "http://localhost:4000/chaincodes?peer=peer1&type=installed" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "GET query Instantiated chaincodes"
echo
curl -s -X GET \
  "http://localhost:4000/chaincodes?peer=peer1&type=instantiated" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "GET query Channels"
echo
curl -s -X GET \
  "http://localhost:4000/channels?peer=peer1" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo



