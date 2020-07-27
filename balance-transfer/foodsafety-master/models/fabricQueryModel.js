var fabricQueryModel = {};
'use strict';

var hfc = require('fabric-client');
var path = require('path');
var sdkUtils = require('fabric-client/lib/utils')
var fs = require('fs');


//SDK调用链码，查询数据
fabricQueryModel.findData = function(options, inputfcn, inputargs, callback) {
    //回调
    var resData = {
        code: 0,
        data: {}
    }
    //inputchaincode = "myccfoodsafety"
  /*  var options = {
        user_id: 'Admin@org1.example.com',
        msp_id: 'Org1MSP',
        channel_id: 'mychannel',
        chaincode_id: inputchaincode,
        network_url: 'grpcs://192.168.1.103:7051', //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        privateKeyFolder: '/home/liukai/go/src/github.com/hyperledger/fabric/examples/food_cli/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore',
        signedCert: '/home/liukai/go/src/github.com/hyperledger/fabric/examples/food_cli/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem',
        tls_cacerts: '/home/liukai/go/src/github.com/hyperledger/fabric/examples/food_cli/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt',
        server_hostname: "peer0.org1.example.com"
    }; */

    var channel = {};
    var client = null;
    const getKeyFilesInDir = (dir) => {
        //该函数用于找到keystore目录下的私钥文件的路径 
        var files = fs.readdirSync(dir)
        var keyFiles = []
        files.forEach((file_name) => {
            let filePath = path.join(dir, file_name)
            if (file_name.endsWith('_sk')) {
                keyFiles.push(filePath)
            }
        })
        return keyFiles
    }
    Promise.resolve().then(() => {
        //指定了当前用户的私钥，证书等基本信息 
        console.log("Load privateKey and signedCert");
        client = new hfc();
        var createUserOpt = {
            username: options.user_id,
            mspid: options.msp_id,
            cryptoContent: {
                privateKey: getKeyFilesInDir(options.privateKeyFolder)[0],
                signedCert: options.signedCert
            }
        }

        return sdkUtils.newKeyValueStore({
            path: "/tmp/fabric-client-stateStore/"
        }).then((store) => {
            client.setStateStore(store)
            return client.createUser(createUserOpt)
        })
    }).then((user) => {
        //因为启用了TLS，所以需要指定Peer的TLS的CA证书 
        channel = client.newChannel(options.channel_id);
        let data = fs.readFileSync(options.tls_cacerts);
        let peer = client.newPeer(options.network_url, {
            pem: Buffer.from(data).toString(),
            'ssl-target-name-override': options.server_hostname
        });
        peer.setName("peer0");
        channel.addPeer(peer);
        return;
    }).then(() => {
        console.log("Make query");
        var transaction_id = client.newTransactionID();
        console.log("Assigning transaction_id: ", transaction_id._transaction_id);
        //构造查询request参数 
        const request = {
            chaincodeId: options.chaincode_id,
            txId: transaction_id,
            fcn: inputfcn,
            args: inputargs
        };
        return channel.queryByChaincode(request);
    }).then((query_responses) => {
        console.log("returned from query");
        if (!query_responses.length) {
            console.log("No payloads were returned from query");
        } else {
            console.log("Query result count = ", query_responses.length)
        }
        if (query_responses[0] instanceof Error) {
            console.error("error from query = ", query_responses[0]);
        }
        console.log("Response is ", query_responses[0].toString('utf8')); //打印返回的结果 

        var oo = JSON.parse(query_responses[0].toString('utf8')); //字符串转成JSON
        //console.log(oo); 
        var srcArr = [];
        if (oo instanceof Array) { //判断对象是否是Array的实例
            oo.forEach(function(v, i) {
                srcArr[i] = v.Record;
            });
            resData.code = 1;
            resData.data = srcArr;
            console.log('Data: ', srcArr);
        } else {
            resData.code = 1;
            resData.data = oo;
            console.log('Data: ', oo);
        }

        callback(resData);

    }).catch((err) => {
        console.error("Caught Error", err);
        callback(resData);
    });

}

module.exports = fabricQueryModel;