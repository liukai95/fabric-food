var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var seedling2plantModel = {};

//遍历seedling2plant表
seedling2plantModel.findAllSeedling2plant = function(callback) {
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
    var fcn= "readAll";//invoke操作方法
    var args= ["SEEDLINGPLANT000","SEEDLINGPLANT999"];//参数

    fabricQueryModel.findData(options, fcn, args, callback);
}
//改进查询所有seedling2plant数据
seedling2plantModel.findAllSeedling2plant2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["SEEDLINGPLANT000","SEEDLINGPLANT999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_ONE_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}
// 根据种子id查询秧苗id
seedling2plantModel.findPlantIdBySeedlingId = function(seedlingid, cb) {
	/*var sql = 'select plantid from seedling2plant where seedlingid = "' + seedlingid + '"';
	connection.query(sql, function (error, results, fields) {
		var resData = {
			code: 0,
			data: null
		}
		if (error) {
			throw error;
		}else {
			resData.code = 1;
			resData.data = results;
			console.log('plantid: ', results);
		}
		cb(resData);
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
    var fcn= "readOneById";//invoke操作方法
    var args= ["XSEEDLINGPLANT" + seedid];//参数

   // fabricQueryModel.findData(options, fcn, args, cb);
    fabricQueryModel.findData(options, fcn, args2, function(redata) {
    	resData.code = 1;
		resData.data = redata.data.plantid;;
        cb(resData);
    });

}

// 根据秧苗id查询种子id
seedling2plantModel.findSeedlingIdByPlantId = function(cb) {
	
}
//插入秧苗种植信息
seedling2plantModel.insertSeedling2plant = function(inputseedlingid, inputplantid, inputpersonid, callback) {
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
    var fcn = "initSeedling2plant"; //invoke操作方法
    var args = [inputseedlingid, inputplantid, inputpersonid]; //参数

    fabricInvokeModel.findData(invokeoptions, fcn, args, callback);
}

module.exports = seedling2plantModel;