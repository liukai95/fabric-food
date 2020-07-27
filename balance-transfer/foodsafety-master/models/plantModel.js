var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var plantModel = {};

//遍历plant表
plantModel.findAllPlant = function(callback) {
    var options = {
        user_id: config.USER_ID_ADMIN_ORG1,
        msp_id: config.MSP_ID_ORG1,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_ONE_ID,
        network_url: config.NETWOEK_URL_ORG1_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG1,
        signedCert: config.SIGNEDCERT_ADMIN_ORG1,
        tls_cacerts: config.TLS_CACERTS_ORG1_PEER0,
        server_hostname: config.SERVER_HOSTNAME_ORG1_PEER0
    };
    var fcn = "readAll"; //invoke操作方法
    var args = ["PLANT000", "PLANT999"]; //参数

    fabricQueryModel.findData(options, fcn, args, callback);
}
//改进查询所有plant数据
plantModel.findAllPlant2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["PLANT000", "PLANT999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_ONE_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}

//种植查询
plantModel.findPlantInfoById = function(id, callback) {
    var resData = {
        code: 1,
        data: null
    }
    
    // var sql = 'select * from seed where seedid = "' + id + '"';
    /*var sql = 'select * from plant,plantusedrug where plant.plantid = "' + id + '" and plantusedrug.plantid = "' + id + '"';
    console.log(sql);
    connection.query(sql, function(error, results) {
    	if(error) {
    		resData.code = 0;
    		console.log(error)
    	}else {
    		resData.data = results;
            console.log(results);
    	}
    	callback(resData);
    });*/
    var options = {
        user_id: config.USER_ID_ADMIN_ORG1,
        msp_id: config.MSP_ID_ORG1,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_ONE_ID,
        network_url: config.NETWOEK_URL_ORG1_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG1,
        signedCert: config.SIGNEDCERT_ADMIN_ORG1,
        tls_cacerts: config.TLS_CACERTS_ORG1_PEER0,
        server_hostname: config.SERVER_HOSTNAME_ORG1_PEER0
    };
    var fcn = "readOneById"; //invoke操作方法
    var args = [id]; //参数
    var args2 = ["XPLANTUSEDRUG" + id]; //参数
    fabricQueryModel.findData(options, fcn, args, function(redata) {
        fabricQueryModel.findData(options, fcn, args2, function(redata2) {
            console.log("plant"+redata.data);
            console.log("plantusedrug"+redata2.data);
            //合并两个json对象
            var resultJsonObject = {};
            for (var attr in redata.data) {
                console.log("aa"+attr);
                resultJsonObject[attr] = redata.data[attr];
            }
            for (var attr in redata2.data) {
                resultJsonObject[attr] = redata.data[attr];
            }
            resData.data = resultJsonObject;
            console.log("resultJsonObject"+resultJsonObject);
            callback(resData);
        });
    });
}
//插入种植信息
plantModel.insertPlant = function(inputplantid, inputplace, inputstartdate, inputpersonid, callback) {
    //链码和证书等参数
    var invokeoptions = {
        user_id: config.USER_ID_ADMIN_ORG1,
        msp_id: config.MSP_ID_ORG1,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_ONE_ID,
        peer_url: config.NETWOEK_URL_ORG1_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        event_url: config.ORG1_PEER0_EVENT_URL,
        orderer_url: config.ORDERER_URL,
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG1,
        signedCert: config.SIGNEDCERT_ADMIN_ORG1,
        peer_tls_cacerts: config.TLS_CACERTS_ORG1_PEER0,
        orderer_tls_cacerts: config.TLS_CACERTS_ORDERER,
        server_hostname: config.SERVER_HOSTNAME_ORG1_PEER0
    };
    var fcn = "initPlant"; //invoke操作方法
    var args = [inputplantid, inputplace, inputstartdate, inputpersonid]; //参数

    fabricInvokeModel.findData(invokeoptions, fcn, args, callback);
}

module.exports = plantModel;