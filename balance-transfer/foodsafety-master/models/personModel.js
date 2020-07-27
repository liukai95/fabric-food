var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var personModel = {};


//查询所有person数据
personModel.findAllUser = function(callback) {
    //链码和证书等参数
    var options = {
        user_id: config.USER_ID_ADMIN_ORG1,
        msp_id: config.MSP_ID_ORG1,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_ALL_ID,
        network_url: config.NETWOEK_URL_ORG1_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG1,
        signedCert: config.SIGNEDCERT_ADMIN_ORG1,
        tls_cacerts: config.TLS_CACERTS_ORG1_PEER0,
        server_hostname: config.SERVER_HOSTNAME_ORG1_PEER0
    };
    var fcn = "readAll"; //invoke操作方法
    var args = ["PERSON000", "PERSON999"]; //参数

    fabricQueryModel.findData(options, fcn, args, callback);


}
//改进查询所有person数据
personModel.findAllUser2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["PERSON000", "PERSON999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_ALL_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}
//登陆时检索是否有该质检员
personModel.findUser = function(inputpersonid, inputpasswd, callback) {
    //链码和证书等参数
    var options = {
        user_id: config.USER_ID_ADMIN_ORG1,
        msp_id: config.MSP_ID_ORG1,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_ALL_ID,
        network_url: config.NETWOEK_URL_ORG1_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG1,
        signedCert: config.SIGNEDCERT_ADMIN_ORG1,
        tls_cacerts: config.TLS_CACERTS_ORG1_PEER0,
        server_hostname: config.SERVER_HOSTNAME_ORG1_PEER0
    };
    var fcn = "readOneById"; //invoke操作方法
    var args = [inputpersonid]; //参数
    var resData = {
        code: 0,
        personid: ""
    }
    fabricQueryModel.findData(options, fcn, args, function(result) {
        if (inputpasswd == result.data.password) {
            resData.code = 1;
            resData.personid = result.data.personid;
        }
        callback(resData);

    });
}

//插入质检员
personModel.insertUser = function(inputpersonid, inputname, inputsex, inputworkplace, inputjob, inputpassword, callback) {
    //链码和证书等参数
    var invokeoptions = {
        user_id: config.USER_ID_ADMIN_ORG1,
        msp_id: config.MSP_ID_ORG1,
        channel_id: config.CHANNEL_ID,
        chaincode_id: config.CHAINCODE_ALL_ID,
        peer_url: config.NETWOEK_URL_ORG1_PEER0, //因为启用了TLS，所以是grpcs,如果没有启用TLS，那么就是grpc 
        event_url: config.ORG1_PEER0_EVENT_URL,
        orderer_url: config.ORDERER_URL,
        privateKeyFolder: config.PRIVATEKEYFOLDER_ADMIN_ORG1,
        signedCert: config.SIGNEDCERT_ADMIN_ORG1,
        peer_tls_cacerts: config.TLS_CACERTS_ORG1_PEER0,
        orderer_tls_cacerts: config.TLS_CACERTS_ORDERER,
        server_hostname: config.SERVER_HOSTNAME_ORG1_PEER0
    };

    var fcn = "initPerson"; //invoke操作方法
    var args = [inputpersonid, inputname, inputsex, inputworkplace, inputjob, inputpassword]; //参数

    fabricInvokeModel.findData(invokeoptions, fcn, args, callback);
}
/*
personModel.insertUser = function(inputid, inputname, inputpasswd, inputaccount, callback) {
    connection.query('insert into person set ?', {
        personid: inputid,
        name: inputname,
        password: inputpasswd,
        account: inputaccount
    }, function(error, results, fields) {
        var resData = {
            code: 0,
            data: {}
        }

        if (error) {
            console.log("insert error");
            throw error;
        } else {
            console.log(results);
            resData.code = 1;
            resData.data = results;
        }

        callback(resData);

    });
}
*/
module.exports = personModel;