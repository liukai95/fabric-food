var plantModel = require('../models/plantModel');
var plantusedrugModel = require('../models/plantusedrugModel');
var seedling2plantModel = require('../models/seedling2plantModel');
var PlantController = {};

//遍历plant表
PlantController.plant = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	plantModel.findAllPlant(function(data) {
		resData.data = data.data;

		res.render('Plant/plant', resData);
	});

}

//遍历plantusedrug表
PlantController.plantusedrug = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	plantusedrugModel.findAllPlantusedrug(function(data) {
		resData.data = data.data;

		res.render('Plant/plantusedrug', resData);
	});

}

//遍历seedling2plan表
PlantController.seedling2plant = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	seedling2plantModel.findAllSeedling2plant(function(data) {
		resData.data = data.data;

		res.render('Plant/seedling2plant', resData);
	});

}

//显示添加种植页面
PlantController.addplant = function(req, res, next) {
	res.render('Plant/addplant');
}
//添加种植
PlantController.doaddplant = function(req, res, next) {
	var plantid = req.body.plantid;
	var place = req.body.place;
	var startdate = req.body.startdate;
	var personid = req.body.personid;
console.log("ss"+startdate);
	plantModel.insertPlant(plantid, place, startdate, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/plant');
		}else {
			res.redirect('/addplant');
		}
	});
}

//显示添加种植用药页面
PlantController.addplantusedrug = function(req, res, next) {
	res.render('Plant/addplantusedrug');
}
//添加种植用药
PlantController.doaddplantusedrug = function(req, res, next) {
	var plantid = req.body.plantid;
	var drugid = req.body.drugid;
	var dosage = req.body.dosage;
	var effect = req.body.effect;
	var date = req.body.date;
	var personid = req.body.personid;

	plantusedrugModel.insertPlantusedrug(plantid, drugid, dosage, effect, date, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/plantusedrug');
		}else {
			res.redirect('/addplantusedrug');
		}
	});
}

//显示添加秧苗种植转换页面
PlantController.addseedling2plant = function(req, res, next) {
	res.render('Plant/addseedling2plant');
}
//添加秧苗种植
PlantController.doaddseedling2plant = function(req, res, next) {
	var seedlingid = req.body.seedlingid;
	var plantid = req.body.plantid;
	var personid = req.body.personid;

	seedling2plantModel.insertSeedling2plant(seedlingid, plantid, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/seedling2plant');
		}else {
			res.redirect('/addseedling2plant');
		}
	});
}
module.exports = PlantController;