var fabricQueryModel = require("./fabricQueryModel");
var fabricInvokeModel = require("./fabricInvokeModel");
var config = require('../config');
var querySDK = require('../nodesdk/query');
var seedModel = {};

//查询所有seed数据
seedModel.findAllSeed = function(callback) {
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
    var args = ["SEED000", "SEED999"]; //参数

    fabricQueryModel.findData(options, fcn, args, callback);
}
//改进查询所有seed数据
seedModel.findAllSeed2 = function(peerorg, callback) {

    var fcn = "readAll"; //invoke操作方法
    var args = ["SEED000", "SEED999"]; //参数
    querySDK.queryChaincode(peerorg.peer_id, config.CHANNEL_ID, config.CHAINCODE_ONE_ID, args, fcn, peerorg.user_id, peerorg.org_id, callback);

}
//种子查询
seedModel.findSeedInfoById = function(id, callback) {
    var resData = {
        code: 1,
        data: null
    }
    // var sql = 'select * from seed where seedid = "' + id + '"';
    /* var sql = 'select seed.seedid, seed.variety, seed.type, \
                 drug.name, drug.dosage, drug.standard, drug.effect, \
                 seedsoakdrug.concentration, seedsoakdrug.startdate, seedsoakdrug.enddata  \
                 from seed, seedsoakdrug, drug \
                 where seed.seedid = "' + id + '" and seedsoakdrug.seedid = "' + id + '" and drug.drugid = seedsoakdrug.drugid';
console.log(sql);     
     connection.query(sql, function(error, results) {
         if (error) {
             resData.code = 0;
             console.log(error)
         } else {
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

    //查询seed
    fabricQueryModel.findData(options, fcn, args, function(redata1) {
        args = ["XSEEDSOAKDRUG" + id]; //seedsoakdrug参数
        fabricQueryModel.findData(options, fcn, args, function(redata2) {
            args = [redata2.data.drugid]; //drug参数
            fabricQueryModel.findData(options, fcn, args, function(redata3) {
                var value = new Object();
                value.seedid = redata1.data.seedid;
                value.variety = redata1.data.variety;
                value.type = redata1.data.type;
                value.name = redata3.data.name;
                value.dosage = redata3.data.dosage;
                value.standard = redata3.data.standard;
                value.effect = redata3.data.effect;
                value.concentration = redata2.data.concentration;
                value.startdate = redata2.data.startdate;
                value.enddata = redata2.data.enddata;
                resData.data = value;
                callback(resData);

            });
        });
    });
}


seedModel.findProductBySeed = function(id, callback) {
    var resData = {
        code: 1,
        data: null
    }
   /*  var sql = 'select seed2seedling.seedid, seed2seedling.seedlingid, \
                        seedling2plant.plantid, \
                        plant2input.inputid, \
                        input2warehouse.warehouseid, \
                        warehouse2feed.feedid, \
                        feed2product.productid \
                 from    seed2seedling, seedling2plant, plant2input, \
                         input2warehouse, warehouse2feed, feed2product \
                 where   seed2seedling.seedid = "' + id + '" and \
                         seed2seedling.seedlingid = seedling2plant.seedlingid and \
                         seedling2plant.plantid = plant2input.plantid and \
                         plant2input.inputid = input2warehouse.inputid and \
                         input2warehouse.warehouseid = warehouse2feed.warehouseid and \
                         warehouse2feed.feedid = feed2product.feedid';
console.log(sql);
     connection.query(sql, function(error, results) {
         if (error) {
             resData.code = 0;
             console.log(error)
         } else {
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
    var options2 = {
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
    var fcn = "readOneById"; //invoke操作方法
    var fcn2 = "readI2WByInputid"; //invoke查询入库方法
    var args = ["XSEEDSEEDLING" + id]; //seed2seedling参数

    var srcArr = [];
    fabricQueryModel.findData(options, fcn, args, function(redata1) {
        args = ["XSEEDLINGPLANT" + redata1.data.seedlingid]; //参数
        fabricQueryModel.findData(options, fcn, args, function(redata2) {
            args = ["XPLANTINPUT" + redata2.data.plantid]; //参数
            fabricQueryModel.findData(options2, fcn, args, function(redata3) {
                args = [redata3.data.inputid]; //参数                
                fabricQueryModel.findData(options2, fcn2, args, function(redata4) {
                    //可能入库相同，仓库不同，产生多组数据
                    redata4.data.forEach(function(v4, i) {
                        args = ["XWAREHOUSEFEED" + v4.warehouseid]; //参数
                        fabricQueryModel.findData(options2, fcn, args, function(redata5) {
                            args = ["XFEEDPRODUCT" + redata5.data.feedid]; //参数
                            fabricQueryModel.findData(options2, fcn, args, function(redata6) {
                                var value = new Object();
                                value.seedid = redata1.data.seedid;
                                value.seedlingid = redata2.data.seedlingid;
                                value.plantid = redata3.data.plantid;
                                value.inputid = v4.inputid;
                                value.warehouseid = redata5.data.warehouseid;
                                value.feedid = redata6.data.feedid;
                                value.productid = redata6.data.productid;
                                //console.log(value);
                                srcArr.push(value);
                                if (srcArr.length == redata4.data.length) {
                                    resData.code = 1;
                                    resData.data = srcArr;
                                    console.log(srcArr);
                                    callback(resData);
                                }

                            });
                        });
                    })
                });
            });
        });
    });
}

//插入种子信息
seedModel.insertSeed = function(seedid, variety, type, personid, callback) {
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
    var fcn = "initSeed"; //invoke操作方法
    var args = [seedid, variety, type, personid]; //参数

    fabricInvokeModel.findData(invokeoptions, fcn, args, callback);
}

module.exports = seedModel;