var express = require('express');
var Router = express.Router();

var IndexController = require('./controllers/IndexController');
var LoginController = require('./controllers/LoginController');
var PersonController = require('./controllers/PersonController');
var StorageController = require('./controllers/StorageController');
var RoughProcessController = require('./controllers/RoughProcessController');
var DeepProcessController = require('./controllers/DeepProcessController');
var SeedingController = require('./controllers/SeedingController');
var PlantController = require('./controllers/PlantController');

var getInfo = require('./api/getInfo');


var Trace = require('./controllers/Trace');




Router.get('/', IndexController.index);
Router.get('/index', IndexController.index);

Router.get('/login', LoginController.login);//显示登陆页面
Router.post('/dologin', LoginController.dologin);//实现登陆

Router.get('/register', LoginController.register);//显示注册页面
Router.post('/doregister', LoginController.doregister);//实现注册

Router.get('/person', PersonController.person);//显示人员管理页面
Router.get('/addperson', PersonController.addperson);//显示添加人员页面
Router.post('/doaddperson', PersonController.doaddperson);//实现添加用户

Router.get('/drug', SeedingController.drug);//显示drug页面
Router.get('/seedling', SeedingController.seedling);//显示seedling页面
Router.get('/seed', SeedingController.seed);//显示seed页面
Router.get('/seedsoakdrug', SeedingController.seedsoakdrug);//显示seedsoakdrug页面
Router.get('/seedlingspraydrug', SeedingController.seedlingspraydrug);//显示seedspraydrug页面
Router.get('/seed2seedling', SeedingController.seed2seedling);//显示seed2seedling页面

Router.get('/plant', PlantController.plant);//显示plant页面
Router.get('/plantusedrug', PlantController.plantusedrug);//显示plantusedrug页面
Router.get('/seedling2plant', PlantController.seedling2plant);//显示seedling2plant页面


Router.get('/input', StorageController.input);//显示input页面
Router.get('/warehouse', StorageController.warehouse);//显示warehouse页面
Router.get('/warehouse2feed', StorageController.warehouse2feed);//显示warehouse2feed页面
Router.get('/plant2input', StorageController.plant2input);//显示plant2input页面
Router.get('/input2warehouse', StorageController.input2warehouse);//显示input2warehouse页面

Router.get('/addplant', PlantController.addplant);
Router.post('/doaddplant', PlantController.doaddplant);
Router.get('/addplantusedrug', PlantController.addplantusedrug);
Router.post('/doaddplantusedrug', PlantController.doaddplantusedrug);
Router.get('/addseedling2plant', PlantController.addseedling2plant);
Router.post('/doaddseedling2plant', PlantController.doaddseedling2plant);

Router.get('/addproduct', DeepProcessController.addproduct);
Router.post('/doaddproduct', DeepProcessController.doaddproduct);
Router.get('/addmaterial2product', DeepProcessController.addmaterial2product);
Router.post('/doaddmaterial2product', DeepProcessController.doaddmaterial2product);
Router.get('/addmaterial', DeepProcessController.addmaterial);
Router.post('/doaddmaterial', DeepProcessController.doaddmaterial);

Router.get('/addinput', StorageController.addinput);
Router.post('/doaddinput', StorageController.doaddinput);
Router.get('/addwarehouse', StorageController.addwarehouse);
Router.post('/doaddwarehouse', StorageController.doaddwarehouse);
Router.get('/addinput2warehouse', StorageController.addinput2warehouse);
Router.post('/doaddinput2warehouse', StorageController.doaddinput2warehouse);
Router.get('/addwarehouse2feed', StorageController.addwarehouse2feed);
Router.post('/doaddwarehouse2feed', StorageController.doaddwarehouse2feed);
Router.get('/addplant2input', StorageController.addplant2input);
Router.post('/doaddplant2input', StorageController.doaddplant2input);

Router.get('/feed', RoughProcessController.feed);//显示feed页面
Router.get('/feed2product', RoughProcessController.feed2product);//显示feed2product页面

Router.get('/addfeed', RoughProcessController.addfeed);
Router.post('/doaddfeed', RoughProcessController.doaddfeed);
Router.get('/addfeed2product', RoughProcessController.addfeed2product);
Router.post('/doaddfeed2product', RoughProcessController.doaddfeed2product);

Router.get('/addseed', SeedingController.addseed);
Router.post('/doaddseed', SeedingController.doaddseed);
Router.get('/addseed2seedling', SeedingController.addseed2seedling);
Router.post('/doaddseed2seedling', SeedingController.doaddseed2seedling);
Router.get('/addseedling', SeedingController.addseedling);
Router.post('/doaddseedling', SeedingController.doaddseedling);
Router.get('/addseedlingspraydrug', SeedingController.addseedlingspraydrug);
Router.post('/doaddseedlingspraydrug', SeedingController.doaddseedlingspraydrug);
Router.get('/adddrug', SeedingController.adddrug);
Router.post('/doadddrug', SeedingController.doadddrug);
Router.get('/addseedsoakdrug', SeedingController.addseedsoakdrug);
Router.post('/doaddseedsoakdrug', SeedingController.doaddseedsoakdrug);

Router.get('/material', DeepProcessController.material);//显示material页面
Router.get('/product', DeepProcessController.product);//显示product页面
Router.get('/material2product', DeepProcessController.material2product);//显示material2product页面

Router.get('/trace/', Trace.index);
Router.get('/safetrace/', Trace.indexsafe);
Router.post('/trace/trace', Trace.trace);


Router.get('/api/getseedinfo/:seedid', getInfo.getSeedInfo);//种子查询
Router.get('/api/getseedlinginfo/:seedlingid', getInfo.getSeedlingInfo);//秧苗查询
Router.get('/api/getplantinfo/:plantid', getInfo.getPlantInfo);//种植查询
Router.get('/api/getinputinfo/:inputid', getInfo.getInputInfo);//入库查询
Router.get('/api/getwarehouseinfo/:warehouseid', getInfo.getWarehouseInfo);//仓库查询
Router.get('/api/getfeedinfo/:feedid', getInfo.getFeedInfo);//进料查询
Router.get('/api/getmaterialinfo/:materialid', getInfo.getMaterialInfo);//进料查询
Router.get('/api/getproductinfo/:productid', getInfo.getProductInfo);//进料查询

Router.get('/api/getproductbyseed/:seedid', getInfo.getProductBySeed);//


Router.get('/api/getpersonpeerinfo/:peerid', getInfo.getPersonPeerInfo);//按照peerid质检员查询
Router.get('/api/getseedpeerinfo/:peerid', getInfo.getSeedPeerInfo);
Router.get('/api/getseedlingpeerinfo/:peerid', getInfo.getSeedlingPeerInfo);
Router.get('/api/getplantpeerinfo/:peerid', getInfo.getPlantPeerInfo);
Router.get('/api/getinputpeerinfo/:peerid', getInfo.getInputPeerInfo);
Router.get('/api/getwarehousepeerinfo/:peerid', getInfo.getWarehousePeerInfo);
Router.get('/api/getfeedpeerinfo/:peerid', getInfo.getFeedPeerInfo);
Router.get('/api/getmaterialpeerinfo/:peerid', getInfo.getMaterialPeerInfo);
Router.get('/api/getproductpeerinfo/:peerid', getInfo.getProductPeerInfo);
Router.get('/api/getwarehouse2feedpeerinfo/:peerid', getInfo.getWarehouse2feedPeerInfo);
Router.get('/api/getinput2warehousepeerinfo/:peerid', getInfo.getInput2warehousePeerInfo);
Router.get('/api/getdrugpeerinfo/:peerid', getInfo.getDrugPeerInfo);
Router.get('/api/getseedsoakdrugpeerinfo/:peerid', getInfo.getSeedsoakdrugPeerInfo);
Router.get('/api/getseed2seedlingpeerinfo/:peerid', getInfo.getSeed2seedlingPeerInfo);
Router.get('/api/getseedlingspraydrugpeerinfo/:peerid', getInfo.getSeedlingspraydrugPeerInfo);
Router.get('/api/getfeed2productpeerinfo/:peerid', getInfo.getFeed2productPeerInfo);
Router.get('/api/getplantusedrugpeerinfo/:peerid', getInfo.getPlantusedrugPeerInfo);
Router.get('/api/getseedling2plantpeerinfo/:peerid', getInfo.getSeedling2plantPeerInfo);
Router.get('/api/getmaterial2productpeerinfo/:peerid', getInfo.getMaterial2productPeerInfo);
Router.get('/api/getplant2inputpeerinfo/:peerid', getInfo.getPlant2inputpeerinfo);
module.exports = Router;
