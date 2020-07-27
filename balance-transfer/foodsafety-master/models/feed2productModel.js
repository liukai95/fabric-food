var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var feed2productModel = {};

//遍历feed2product表
feed2productModel.findAllfeed2product = function(callback) {
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
    var args= ["FEEDPRODUCT000","FEEDPRODUCT999"];//参数

    fabricQueryModel.findData(options, fcn, args, callback);
}
//改进查询所有feed2product数据
feed2productModel.findAllFeed2product2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["FEEDPRODUCT000","FEEDPRODUCT999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_TWO_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}
//插入feed2product信息
feed2productModel.insertFeed2product = function(feedid, productid, personid, callback) {
    //链码和证书等参数
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
    var fcn = "initFeed2product"; //invoke操作方法
    var args = [feedid, productid, personid]; //参数

    fabricInvokeModel.findData(invokeoptions, fcn, args, callback);
}

module.exports = feed2productModel;
