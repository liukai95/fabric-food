var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var productModel = {};

//遍历product表
productModel.findAllProduct = function(callback) {
    var options = {
        user_id: config.USER_ID_ADMIN_ORG2,
        msp_id: config.MSP_ID_ORG2,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_TWO_ID,
        network_url: config.NETWOEK_URL_ORG2_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG2,
        signedCert: config.SIGNEDCERT_ADMIN_ORG2,
        tls_cacerts: config.TLS_CACERTS_ORG2_PEER0,
        server_hostname: config.SERVER_HOSTNAME_ORG2_PEER0
    };
    var fcn= "readAll";//invoke操作方法
    var args= ["PRODUCT000","PRODUCT999"];//参数

    fabricQueryModel.findData(options, fcn, args, callback);
}
//改进查询所有product数据
productModel.findAllProduct2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["PRODUCT000","PRODUCT999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_TWO_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}

//根据id查询
productModel.findProductInfoById = function(id, callback) {
    var options = {
        user_id: config.USER_ID_ADMIN_ORG2,
        msp_id: config.MSP_ID_ORG2,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_TWO_ID,
        network_url: config.NETWOEK_URL_ORG2_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG2,
        signedCert: config.SIGNEDCERT_ADMIN_ORG2,
        tls_cacerts: config.TLS_CACERTS_ORG2_PEER0,
        server_hostname: config.SERVER_HOSTNAME_ORG2_PEER0
    };
    var fcn= "readOneById";//invoke操作方法
    var args= [id];//参数

    fabricQueryModel.findData(options, fcn, args, callback);
    /*	var sql = "select productid, name, specification, flavor, date from product where productid = '" + productid + "'";
	console.log(sql)
	connection.query(sql, function(error, results) {
		if(error) {
			res.send("error")
		}else {
console.log(results);
			res.send(results);
		}
	});*/
}
//插入产品信息
productModel.insertProduct = function(productid, name, specification, flavor, date, personid, callback) {
    var invokeoptions = {
        user_id: config.USER_ID_ADMIN_ORG2,
        msp_id: config.MSP_ID_ORG2,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_TWO_ID,
        peer_url: config.NETWOEK_URL_ORG2_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        event_url: config.ORG2_PEER0_EVENT_URL,
        orderer_url: config.ORDERER_URL,
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG2,
        signedCert: config.SIGNEDCERT_ADMIN_ORG2,
        peer_tls_cacerts: config.TLS_CACERTS_ORG2_PEER0,
        orderer_tls_cacerts: config.TLS_CACERTS_ORDERER,
        server_hostname: config.SERVER_HOSTNAME_ORG2_PEER0
    };
    var fcn = "initProduct"; //invoke操作方法
    var args = [productid, name, specification, flavor, date, personid]; //参数

    fabricInvokeModel.findData(invokeoptions, fcn, args, callback);
}
module.exports = productModel;