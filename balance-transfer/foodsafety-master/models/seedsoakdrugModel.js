var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var seedsoakdrugModel = {};

//遍历seedsoakdrug表
seedsoakdrugModel.findAllSeedsoakdrug = function(callback) {
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
    var args= ["SEEDSOAKDRUG000","SEEDSOAKDRUG999"];//参数

    fabricQueryModel.findData(options, fcn, args, callback);
}
//改进查询所有seedsoakdrug数据
seedsoakdrugModel.findAllSeedsoakdrug2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["SEEDSOAKDRUG000","SEEDSOAKDRUG999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_ONE_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}
//插入种子用药信息
seedsoakdrugModel.insertSeedsoakdrug = function(seedid, drugid, concentration, startdate, enddata, personid, callback) {
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
    var fcn = "initSeedsoakdrug"; //invoke操作方法
    var args = [seedid, drugid, concentration, startdate, enddata, personid]; //参数

    fabricInvokeModel.findData(invokeoptions, fcn, args, callback);
}

module.exports = seedsoakdrugModel;