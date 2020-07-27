var drugModel = require('../models/drugModel');
var SeedlingModel = require('../models/seedlingModel');
var SeedModel = require('../models/seedModel');
var SeedsoakdrugModel = require('../models/seedsoakdrugModel');
var Seed2seedlingModel = require('../models/seed2seedlingModel');
var SeedlingspraydrugModel = require('../models/seedlingspraydrugModel');

var SeedingController = {};

//遍历drug表
SeedingController.drug = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	drugModel.findAllDrug(function(data) {
		resData.data = data.data;

		res.render('Seeding/drug', resData);
	});

}

//遍历seedling表
SeedingController.seedling = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	SeedlingModel.findAllSeedling(function(data) {
		resData.data = data.data;

		res.render('Seeding/seedling', resData);
	});

}

//遍历seed表
SeedingController.seed = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	SeedModel.findAllSeed(function(data) {
		resData.data = data.data;

		res.render('Seeding/seed', resData);
	});

}

//遍历seedsoakdrug表
SeedingController.seedsoakdrug = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	SeedsoakdrugModel.findAllSeedsoakdrug(function(data) {
		resData.data = data.data;

		res.render('Seeding/seedsoakdrug', resData);
	});

}

//遍历seed2seedling表
SeedingController.seed2seedling = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	Seed2seedlingModel.findAllSeed2seedling(function(data) {
		resData.data = data.data;

		res.render('Seeding/seed2seedling', resData);
	});

}

//遍历seedlingspraydrug表
SeedingController.seedlingspraydrug = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	SeedlingspraydrugModel.findAllSeedlingspraydrug(function(data) {
		resData.data = data.data;

		res.render('Seeding/seedlingspraydrug', resData);
	});

}

//显示添加药品页面
SeedingController.adddrug = function(req, res, next) {
	res.render('Seeding/adddrug');
}
//添加药品
SeedingController.doadddrug = function(req, res, next) {
	var drugid = req.body.drugid;
	var name = req.body.name;
	var dosage = req.body.dosage;
	var standard = req.body.standard;
	var effect = req.body.effect;
	var personid = req.body.personid;

	drugModel.insertDrug(drugid, name, dosage, standard, effect, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/drug');
		}else {
			res.redirect('/adddrug');
		}
	});
}


//显示添加种子页面
SeedingController.addseed = function(req, res, next) {
	res.render('Seeding/addseed');
}
//添加种子
SeedingController.doaddseed = function(req, res, next) {
	var seedid = req.body.seedid;
	var variety = req.body.variety;
	var type = req.body.type;
	var personid = req.body.personid;

	SeedModel.insertSeed(seedid, variety, type, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/seed');
		}else {
			res.redirect('/addseed');
		}
	});
}


//显示添加种子秧苗页面
SeedingController.addseed2seedling = function(req, res, next) {
	res.render('Seeding/addseed2seedling');
}
//添加种子秧苗
SeedingController.doaddseed2seedling = function(req, res, next) {
	var seedid = req.body.seedid;
	var seedlingid = req.body.seedlingid;
	var personid = req.body.personid;

	Seed2seedlingModel.insertSeed2seedling(seedid, seedlingid, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/seed2seedling');
		}else {
			res.redirect('/addseed2seedling');
		}
	});
}


//显示添加秧苗页面
SeedingController.addseedling = function(req, res, next) {
	res.render('Seeding/addseedling');
}
//添加秧苗
SeedingController.doaddseedling = function(req, res, next) {
	var seedlingid = req.body.seedlingid;
	var nurseryplace = req.body.nurseryplace;
	var startdate = req.body.startdate;
	var personid = req.body.personid;

	SeedlingModel.insertSeedling(seedlingid, nurseryplace, startdate, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/seedling');
		}else {
			res.redirect('/addseedling');
		}
	});
}


//显示添加秧苗喷药页面
SeedingController.addseedlingspraydrug = function(req, res, next) {
	res.render('Seeding/addseedlingspraydrug');
}
//添加秧苗喷药
SeedingController.doaddseedlingspraydrug = function(req, res, next) {
	var seedlingid = req.body.seedlingid;
	var drugid = req.body.drugid;
	var dosage = req.body.dosage;
	var data = req.body.data;
	var personid = req.body.personid;

	SeedlingspraydrugModel.insertSeedlingspraydrug(seedlingid, drugid, dosage, data, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/seedlingspraydrug');
		}else {
			res.redirect('/addseedlingspraydrug');
		}
	});
}


//显示添加种子用药页面
SeedingController.addseedsoakdrug = function(req, res, next) {
	res.render('Seeding/addseedsoakdrug');
}
//添加药品
SeedingController.doaddseedsoakdrug = function(req, res, next) {
	var seedid = req.body.seedid;
	var drugid = req.body.drugid;
	var concentration = req.body.concentration;
	var startdate = req.body.startdate;
	var enddata = req.body.enddata;
	var personid = req.body.personid;

	SeedsoakdrugModel.insertSeedsoakdrug(seedid, drugid, concentration, startdate, enddata, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/seedsoakdrug');
		}else {
			res.redirect('/addseedsoakdrug');
		}
	});
}
module.exports = SeedingController;