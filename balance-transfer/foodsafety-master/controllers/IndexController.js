var personModel = require('../models/personModel');


var IndexController = {};

IndexController.index = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	personModel.findAllUser(function(data) {
		resData.data = data.data;

		res.render('index', resData);
	});

}




module.exports = IndexController;