var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var inputModel = {};

//遍历input表
inputModel.findAllInput = function(callback) {
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
    var args= ["INPUT000","INPUT999"];//参数

    fabricQueryModel.findData(options, fcn, args, callback);
}
//改进查询所有input数据
inputModel.findAllInput2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["INPUT000","INPUT999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_TWO_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}

//入库查询
inputModel.findInputInfoById = function(id, callback) {
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
}
//插入入库信息
inputModel.insertInput = function(inputid, harvestdate, quantity, inputdate, personid, callback) {
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
    var fcn = "initInput"; //invoke操作方法
    var args = [inputid, harvestdate, quantity, inputdate, personid]; //参数

    fabricInvokeModel.findData(invokeoptions, fcn, args, callback);
}
module.exports = inputModel;
