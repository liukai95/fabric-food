var SeedModel = require('../models/seedModel');
var PersonModel = require('../models/personModel');
var SeedlingModel = require('../models/seedlingModel');
var PlantModel = require('../models/plantModel');
var InputModel = require('../models/inputModel');
var Plant2inputModel = require('../models/plant2inputModel');
var WarehouseModel = require('../models/warehouseModel');
var FeedModel = require('../models/feedModel');
var MaterialModel = require('../models/materialModel');
var ProductModel = require('../models/productModel');
var Warehouse2feedModel = require('../models/warehouse2feedModel');
var Input2warehouseModel = require('../models/input2warehouseModel');
var DrugModel = require('../models/drugModel');
var SeedsoakdrugModel = require('../models/seedsoakdrugModel');
var Seed2seedlingModel = require('../models/seed2seedlingModel');
var SeedlingspraydrugModel = require('../models/seedlingspraydrugModel');
var Feed2productModel = require('../models/feed2productModel');
var PlantusedrugModel = require('../models/plantusedrugModel');
var Seedling2plantModel = require('../models/seedling2plantModel');
var Material2productModel = require('../models/material2productModel');

var GetInfo = {};

//种子
GetInfo.getSeedInfo = function(req, res, next) {

    var seedid = req.params.seedid;

    SeedModel.findSeedInfoById(seedid, function(resdata) {
        if (resdata.code == 1) {
            res.send(resdata.data);
        } else {
            res.send("error");
        }
    })

}

//秧苗
GetInfo.getSeedlingInfo = function(req, res, next) {

    var seedlingid = req.params.seedlingid;

    SeedlingModel.findSeedlingInfoById(seedlingid, function(resdata) {
        if (resdata.code == 1) {
            res.send(resdata.data);
        } else {
            res.send("error");
        }
    })

}


//种植
GetInfo.getPlantInfo = function(req, res, next) {

    var plantid = req.params.plantid;

    PlantModel.findPlantInfoById(plantid, function(resdata) {
        if (resdata.code == 1) {
            res.send(resdata.data);
        } else {
            res.send("error");
        }
    })

}

//入库
GetInfo.getInputInfo = function(req, res, next) {

    var inputid = req.params.inputid;

    InputModel.findInputInfoById(inputid, function(resdata) {
        if (resdata.code == 1) {
            res.send(resdata.data);
        } else {
            res.send("error");
        }
    })

}

//仓库
GetInfo.getWarehouseInfo = function(req, res, next) {

    var warehouseid = req.params.warehouseid;

    WarehouseModel.findWarehouseInfoById(warehouseid, function(resdata) {
        if (resdata.code == 1) {
            res.send(resdata.data);
        } else {
            res.send("error");
        }
    })

}

//进料
GetInfo.getFeedInfo = function(req, res, next) {
    var feedid = req.params.feedid;
    FeedModel.findFeedInfoById(feedid, function(resdata) {
        if (resdata.code == 1) {
            res.send(resdata.data);
        } else {
            res.send("error");
        }
    })
}

GetInfo.getMaterialInfo = function(req, res, next) {
    var materialid = req.params.materialid;
    MaterialModel.findMaterialInfoById(materialid, function(resdata) {
        if (resdata.code == 1) {
            res.send(resdata.data);
        } else {
            res.send("error");
        }
    })

}

GetInfo.getProductInfo = function(req, res, next) {
    var productid = req.params.productid;
    ProductModel.findProductInfoById(productid, function(resdata) {
        if (resdata.code == 1) {
            res.send(resdata.data);
        } else {
            res.send("error");
        }
    })

}

GetInfo.getProductBySeed = function(req, res, next) {
    var seedid = req.params.seedid;

    SeedModel.findProductBySeed(seedid, function(resData) {
        if (resData.code == 1) {
            res.send(resData.data);
        } else {
            res.send("error");
        }
    })

}

//分解参数
function getpeerorg(peerorgid) {
    var peerorg = {
        peer_id: ["peer1"],
        user_id: "Jim",
        org_id: "org1"
    };
    if (peerorgid == 2) {
        peerorg.peer_id = ["peer2"];
    } else if (peerorgid == 3) {
        peerorg.user_id = "Barry";
        peerorg.org_id = "org2";
    } else if (peerorgid == 4) {
        peerorg.peer_id = ["peer2"];
        peerorg.user_id = "Barry";
        peerorg.org_id = "org2";
    }
    return peerorg;
}
//质检员
GetInfo.getPersonPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    PersonModel.findAllUser2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Seed
GetInfo.getSeedPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    SeedModel.findAllSeed2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Seedling
GetInfo.getSeedlingPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    SeedlingModel.findAllSeedling2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Plant
GetInfo.getPlantPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    PlantModel.findAllPlant2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Input
GetInfo.getInputPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    InputModel.findAllInput2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//质检员
GetInfo.getPersonPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    PersonModel.findAllUser2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Warehouse
GetInfo.getWarehousePeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    WarehouseModel.findAllWarehouse2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Feed
GetInfo.getFeedPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    FeedModel.findAllFeed2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Material
GetInfo.getMaterialPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    MaterialModel.findAllMaterial2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Product
GetInfo.getProductPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    ProductModel.findAllProduct2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Warehouse2feed
GetInfo.getWarehouse2feedPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    Warehouse2feedModel.findAllWarehouse2feed2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Input2warehouse
GetInfo.getInput2warehousePeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    Input2warehouseModel.findAllInput2warehouse2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Drug
GetInfo.getDrugPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    DrugModel.findAllDrug2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Seedsoakdrug
GetInfo.getSeedsoakdrugPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    SeedsoakdrugModel.findAllSeedsoakdrug2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Seed2seedling
GetInfo.getSeed2seedlingPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    Seed2seedlingModel.findAllSeed2seedling2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Seedlingspraydrug
GetInfo.getSeedlingspraydrugPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    SeedlingspraydrugModel.findAllSeedlingspraydrug2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Feed2product
GetInfo.getFeed2productPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    Feed2productModel.findAllFeed2product2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Plantusedrug
GetInfo.getPlantusedrugPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    PlantusedrugModel.findAllPlantusedrug2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Seedling2plant
GetInfo.getSeedling2plantPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    Seedling2plantModel.findAllSeedling2plant2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}
//Material2product
GetInfo.getMaterial2productPeerInfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    Material2productModel.findAllMaterial2product2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}

//Plant2input
GetInfo.getPlant2inputpeerinfo = function(req, res, next) {

    var peerorgid = req.params.peerid;
    console.log("peerid" + peerorgid);
    var peerorg = getpeerorg(peerorgid);

    Plant2inputModel.findAllPlant2input2(peerorg, function(data) {
        if (data.code == 1) {
            res.send(data.data);
        } else {
            res.send("error");
        }
    })
}

module.exports = GetInfo;
