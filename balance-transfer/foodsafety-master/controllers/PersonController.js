var personModel = require("../models/personModel");

var PersonController = {};

//显示页面
PersonController.person = function(req, res, next) {

	var resData = {
		code: 0,
		data: {}
	}
    var peerorg = {
        peer_id: ["peer1"],
        user_id: "Jim",
        org_id: "org1"
    };
	personModel.findAllUser(function(data) {
		resData.data = data.data;
		res.render('Person/person', resData);
	})
}

//显示添加用户页面
PersonController.addperson = function(req, res, next) {
	res.render('Person/addperson');
}

//实现注册
PersonController.doaddperson = function(req, res, next) {
	var personid = req.body.personid;
	var name = req.body.name;
	var sex = req.body.sex;
	var workplace = req.body.workplace;
	var job = req.body.job;
	var password = req.body.password;

	personModel.insertUser(personid, name, sex, workplace, job, password, function(resData) {

		if (resData == 1) {//插入成功
			res.redirect('/person');
		}else {
			res.redirect('/addperson');
		}
	});
}

module.exports = PersonController;