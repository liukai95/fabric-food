//

var config = {

	APP_PORT: 3000,

	DB_HOST: '192.168.1.105',
	DB_PORT: 3306,
	DB_USER: 'root',
	DB_PASSWD: '123456',
	DB_NAME: 'foodsafety',
	
    USER_ID_ADMIN_ORG1: 'Admin@org1.example.com',
    USER_ID_USER1_ORG1: 'User1@org1.example.com',
    MSP_ID_ORG1: 'Org1MSP',
    USER_ID_ADMIN_ORG2: 'Admin@org2.example.com',
    USER_ID_USER1_ORG2: 'User1@org2.example.com',
    MSP_ID_ORG2: 'Org2MSP',
    CHANNEL_ID: 'mychannel',
    NETWOEK_URL_ORG1_PEER0: 'grpcs://192.168.1.103:7051', 
    NETWOEK_URL_ORG1_PEER1: 'grpcs://192.168.1.103:8051',
    NETWOEK_URL_ORG2_PEER0: 'grpcs://192.168.1.103:9051',
    NETWOEK_URL_ORG2_PEER1: 'grpcs://192.168.1.103:10051', 
    PRIVATEKEYFOLDER_ADMIN_ORG1: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore',
    PRIVATEKEYFOLDER_ADMIN_ORG2: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp/keystore',
    PRIVATEKEYFOLDER_USER1_ORG1: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore',
    PRIVATEKEYFOLDER_USER1_ORG2: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/keystore',
    SIGNEDCERT_ADMIN_ORG1: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem',
    SIGNEDCERT_ADMIN_ORG2: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp/signcerts/Admin@org2.example.com-cert.pem',
    SIGNEDCERT_USER1_ORG1: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem',
    SIGNEDCERT_USER1_ORG2: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/signcerts/User1@org2.example.com-cert.pem',
	TLS_CACERTS_ORG1_PEER0: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt',
    TLS_CACERTS_ORG1_PEER1: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt',
	TLS_CACERTS_ORG2_PEER0: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt',
	TLS_CACERTS_ORG2_PEER1: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt',
	SERVER_HOSTNAME_ORG1_PEER0: "peer0.org1.example.com",
	SERVER_HOSTNAME_ORG1_PEER1: "peer1.org1.example.com",
	SERVER_HOSTNAME_ORG2_PEER0: "peer0.org2.example.com",
	SERVER_HOSTNAME_ORG2_PEER1: "peer1.org2.example.com",
    ORG1_PEER0_EVENT_URL: 'grpcs://192.168.1.103:7053',  
    ORG1_PEER1_EVENT_URL: 'grpcs://192.168.1.103:8053', 
    ORG2_PEER0_EVENT_URL: 'grpcs://192.168.1.103:9053', 
    ORG2_PEER1_EVENT_URL: 'grpcs://192.168.1.103:10053',
    ORDERER_URL: 'grpcs://192.168.1.103:7050',   
    TLS_CACERTS_ORDERER: '/home/liukai/fabric-food/balance-transfer/food_cli/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt', 
	CHAINCODE_ALL_ID: 'myccfoodall',
	CHAINCODE_ONE_ID: 'myccfoodone',
	CHAINCODE_TWO_ID: 'myccfoodtwo'
}

module.exports = config

