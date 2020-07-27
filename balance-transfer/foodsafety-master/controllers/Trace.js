var seed2seedlingModel = require("../models/seed2seedlingModel");
var seedling2plantModel = require("../models/seedling2plantModel");
var plant2inputModel = require("../models/plant2inputModel");
var input2warehouseModel = require("../models/input2warehouseModel");
var warehouse2feedModel = require("../models/warehouse2feedModel");
var feed2productModel = require("../models/feed2productModel");
var config = require('../config');

let api = {},
    private = {};

let Trace = {
    api: api,
    _private: private
}

var fabricQueryModel = require("../models/fabricQueryModel");
Trace.index = function(req, res, next) {
    var data = {
        seed: {},
        product: {}
    }

    /*connection.query('select seedid from seed', function(error, results, fields) {
    	if(error) {
    		next();
    	}
    	data.seed = results;
    	console.log(results)
    	connection.query('select productid from product', function(error, results, fields) {
    		if(error) {
    			next();
    		}
    		data.product = results;
    		res.render('Trace/trace', data);
    	})
    })*/
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
    var fcn = "readAll"; //invoke操作方法
    var args = ["SEED000", "SEED999"]; //参数
    var args2 = ["PRODUCT000", "PRODUCT999"]; //参数2

    fabricQueryModel.findData(options, fcn, args, function(redata) {
        data.seed = redata.data;
        fabricQueryModel.findData(options2, fcn, args2, function(redata2) {
            data.product = redata2.data;
            res.render('Trace/trace', data);
        });
    });


}
Trace.indexsafe = function(req, res, next) {
    var data = {
        seed: {},
        product: {}
    }
    //res.render('Trace/safetrace', data);
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
    var fcn = "readAll"; //invoke操作方法
    var args = ["SEED000", "SEED999"]; //参数
    var args2 = ["PRODUCT000", "PRODUCT999"]; //参数

    fabricQueryModel.findData(options, fcn, args, function(redata) {
        data.seed = redata.data;
        fabricQueryModel.findData(options2, fcn, args2, function(redata2) {
            data.product = redata2.data;
            res.render('Trace/safetrace', data);
        });
    });
}

Trace.trace = function(req, res, next) {
    let kind = req.body.kind,
        code = req.body.code;
    console.log(kind + " : " + code);

    private.trace(kind, code, function(data) {
        res.send(data);
    });


}

// 溯源中转
private.trace = function(kind, code, cb) {
    if (kind == 'seed') {
        this.forwardTrace(code, cb);
    } else if (kind == 'product') {
        this.backwardTrace(code, cb);
    } else {
        next();
    }
}

/*

 {
	1: {
		seed: id
	},
	seedling: id,
	plant: id,
	input: id,
	warehouse: id,
	feed: id,
	material: id,
	product: id
 }


 */



// 正向溯源，从种子开始
private.forwardTrace = function(seedid, cb) {
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
    var args = ["XSEEDSEEDLING" + seedid]; //seed2seedling参数

    /*fabricQueryModel.findData(chaincode_id, fcn, args, function(redata1) {
            redata1.data.forEach(function(v1) {
                args = ["XSEEDLINGPLANT" + v1.seedlingid]; //参数
                fabricQueryModel.findData(chaincode_id, fcn, args, function(redata2) {
                    redata2.data.forEach(function(v2) {
                        args = ["XPLANTINPUT" + v2.plantid]; //参数
                        fabricQueryModel.findData(chaincode_id, fcn, args, function(redata3) {
                            redata3.data.forEach(function(v3) {
                                args = ["XINPUTWAREHOUSE" + v3.inputid]; //参数
                                fabricQueryModel.findData(chaincode_id, fcn, args, function(redata4) {
                                    redata4.data.forEach(function(v4) {
                                        args = ["XWAREHOUSEFEED" + v4.warehouseid]; //参数
                                        fabricQueryModel.findData(chaincode_id, fcn, args, function(redata5) {
                                            redata5.data.forEach(function(v5) {
                                                args = ["XFEEDPRODUCT" + v5.feedid]; //参数
                                                fabricQueryModel.findData(chaincode_id, fcn, args, function(redata6) {
                                                    redata6.data.forEach(function(v6) {
                                                        var value = new Object(); 
value.seedid = v1.seedid; 
value.seedlingid = v2.seedlingid; 
value.plantid = v3.plantid; 
value.inputid = v4.inputid; 
value.warehouseid = v5.warehouseid; 
value.feedid = v6.feedid; 
value.productid = v6.productid; 
console.log(value);

                                                    });
                                                });

                                            });
                                        });
                                    });
                                });
                            });
                        });
                    });
                });
            });
        });*/
    var resData = {
        code: 0,
        data: null
    }
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
                                    cb(resData);
                                }

                            });
                        });
                    })
                });
            });
        });
    });
    //var oo = JSON.stringify(values);
    //console.log(oo);
}

// 反向溯源，从产品开始
private.backwardTrace = function(productid, cb) {
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
    var fcn2 = "readI2WByWarehouseid"; //invoke查询入库方法
    var args = ["XFEEDPRODUCT" + productid]; //参数
    var resData = {
        code: 0,
        data: null
    }
    var srcArr = [];
    fabricQueryModel.findData(options2, fcn, args, function(redata1) {
        args = ["XMATERIALPRODUCT" + redata1.data.productid]; //参数
        fabricQueryModel.findData(options2, fcn, args, function(redata2) {
            args = ["XWAREHOUSEFEED" + redata1.data.feedid]; //参数
            fabricQueryModel.findData(options2, fcn, args, function(redata3) {
                args = [redata3.data.warehouseid]; //参数                
                fabricQueryModel.findData(options2, fcn2, args, function(redata4) {
                    //可能仓库相同，入库不同，产生多组数据
                    redata4.data.forEach(function(v4, i) {
                        args = ["XPLANTINPUT" + v4.inputid]; 
                        fabricQueryModel.findData(options2, fcn, args, function(redata5) {
                            args = ["XSEEDLINGPLANT" + redata5.data.plantid]; //参数
                            fabricQueryModel.findData(options, fcn, args, function(redata6) {
                                args = ["XSEEDSEEDLING" + redata6.data.seedlingid]; //seed2seedling参数
                                fabricQueryModel.findData(options, fcn, args, function(redata7) {

                                    var value = new Object();
                                    value.productid = redata1.data.productid;
                                    value.materialid = redata2.data.materialid;
                                    value.feedid = redata3.data.feedid;
                                    value.warehouseid = v4.warehouseid;
                                    value.inputid = redata5.data.inputid;
                                    value.plantid = redata6.data.plantid;
                                    value.seedlingid = redata7.data.seedlingid;
                                    value.seedid = redata7.data.seedid;
                                    //console.log(value);
                                    srcArr.push(value);
                                    if (srcArr.length == redata4.data.length) {
                                        resData.code = 1;
                                        resData.data = srcArr;
                                        console.log(srcArr);
                                        cb(resData);
                                    }
                                });
                            });
                        });
                    })
                });
            });
        });
    });

    /*var sql = '\
				select  feed2product.productid, material2product.materialid, warehouse2feed.feedid, \
				        input2warehouse.warehouseid, plant2input.inputid, seedling2plant.plantid, \
				        seed2seedling.seedlingid, seed2seedling.seedid  \
				from 	seed2seedling, seedling2plant, plant2input, input2warehouse,  \
					 	warehouse2feed, feed2product, material2product \
				where 	feed2product.productid = "' + productid + '" and \
				        material2product.productid = feed2product.productid and  \
				        feed2product.feedid = warehouse2feed.feedid and \
				      	warehouse2feed.warehouseid = input2warehouse.warehouseid and \
				      	input2warehouse.inputid = plant2input.inputid and \
				      	plant2input.plantid = seedling2plant.plantid and \
				      	seedling2plant.seedlingid = seed2seedling.seedlingid';

    connection.query(sql, function(error, results, fields) {
        var resData = {
            code: 0,
            data: null
        }
        if (!error) {
            resData.code = 1;
            resData.data = results;
        } else {
            console.log(error)
        }
        console.log(resData.data)
        cb(resData);
    })*/
}



private.deepCopyPath = function(obj) {
    let newobj = {};
    for (let key in obj) {
        newobj[key] = obj[key];
    }
    return newobj;
}



module.exports = Trace;