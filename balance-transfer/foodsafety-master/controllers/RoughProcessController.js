var feedModel = require('../models/feedModel');
var feed2productModel = require('../models/feed2productModel');

var RoughProcessController = {};

//遍历feed表
RoughProcessController.feed = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	feedModel.findAllfeed(function(data) {
		resData.data = data.data;

		res.render('RoughProcess/feed', resData);
	});

}

//遍历feed2product表
RoughProcessController.feed2product = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	feed2productModel.findAllfeed2product(function(data) {
		resData.data = data.data;

		res.render('RoughProcess/feed2product', resData);
	});

}

//显示添加进料页面
RoughProcessController.addfeed = function(req, res, next) {
	res.render('RoughProcess/addfeed');
}
//添加进料
RoughProcessController.doaddfeed = function(req, res, next) {
	var feedid = req.body.feedid;
	var weight = req.body.weight;
	var watercontent = req.body.watercontent;
	var brokenrice = req.body.brokenrice;
	var qingmilv = req.body.qingmilv;
	var date = req.body.date;
	var personid = req.body.personid;

	feedModel.insertFeed(feedid, weight, watercontent, brokenrice, qingmilv, date, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/feed');
		}else {
			res.redirect('/addfeed');
		}
	});
}
//显示添加进料产品批次转换页面
RoughProcessController.addfeed2product = function(req, res, next) {
	res.render('RoughProcess/addfeed2product');
}
//添加进料产品批次转换
RoughProcessController.doaddfeed2product = function(req, res, next) {
	var feedid = req.body.feedid;
	var productid = req.body.productid;
	var personid = req.body.personid;

	feed2productModel.insertFeed2product(feedid, productid, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/feed2product');
		}else {
			res.redirect('/addfeed2product');
		}
	});
}
module.exports = RoughProcessController;