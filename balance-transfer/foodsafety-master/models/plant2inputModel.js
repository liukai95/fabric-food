var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var plant2inputModel = {};

//遍历plant2input表
plant2inputModel.findAllPlant2input = function(callback) {
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
    var args= ["PLANTINPUT000","PLANTINPUT999"];//参数

    fabricQueryModel.findData(options, fcn, args, callback);
}
//改进查询所有plant2input数据
plant2inputModel.findAllPlant2input2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["PLANTINPUT000","PLANTINPUT999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_TWO_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}
//入库查询
plant2inputModel.findInputInfoById = function(id, callback) {
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
//插入信息
plant2inputModel.insertPlant2input = function(plantid, inputid, personid, callback) {
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
    var fcn = "initPlant2input"; //invoke操作方法
    var args = [plantid, inputid, personid]; //参数

    fabricInvokeModel.findData(options, fcn, args, callback);
}

module.exports = plant2inputModel;