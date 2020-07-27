var materialModel = require('../models/materialModel');
var productModel = require('../models/productModel');
var material2productModel = require('../models/material2productModel');
var DeepProcessController = {};

//遍历material表
DeepProcessController.material = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	materialModel.findAllMaterial(function(data) {
		resData.data = data.data;

		res.render('DeepProcess/material', resData);
	});

}

//遍历product表
DeepProcessController.product = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	productModel.findAllProduct(function(data) {
		resData.data = data.data;

		res.render('DeepProcess/product', resData);
	});

}

//遍历material2product表
DeepProcessController.material2product = function(req, res, next) {
	var resData = {
		code: 0,
		data: {}
	}
	material2productModel.findAllMaterial2product(function(data) {
		resData.data = data.data;

		res.render('DeepProcess/material2product', resData);
	});

}
//显示添加material页面
DeepProcessController.addmaterial = function(req, res, next) {
	res.render('DeepProcess/addmaterial');
}
//添加material
DeepProcessController.doaddmaterial = function(req, res, next) {
	var materialid = req.body.materialid;
	var kind = req.body.kind;
	var weight = req.body.weight;
	var source = req.body.source;
	var date = req.body.date;
	var personid = req.body.personid;

	materialModel.insertMaterial(materialid, kind, weight, source, date, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/material');
		}else {
			res.redirect('/addmaterial');
		}
	});
}

//显示添加产品页面
DeepProcessController.addproduct = function(req, res, next) {
	res.render('DeepProcess/addproduct');
}
//添加产品
DeepProcessController.doaddproduct = function(req, res, next) {
	var productid = req.body.productid;
	var name = req.body.name;
	var specification = req.body.specification;
	var flavor = req.body.flavor;
	var date = req.body.date;
	var personid = req.body.personid;

	productModel.insertProduct(productid, name, specification, flavor, date, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/product');
		}else {
			res.redirect('/addproduct');
		}
	});
}
//显示添加原料产品转换页面
DeepProcessController.addmaterial2product = function(req, res, next) {
	res.render('DeepProcess/addmaterial2product');
}
//添加原料产品转换
DeepProcessController.doaddmaterial2product = function(req, res, next) {
	var materialid = req.body.materialid;
	var productid = req.body.productid;
	var personid = req.body.personid;

	material2productModel.insertMaterial2product(materialid, productid, personid, function(resData) {

		if (resData == 1) {
			res.redirect('/material2product');
		}else {
			res.redirect('/addmaterial2product');
		}
	});
}

module.exports = DeepProcessController;