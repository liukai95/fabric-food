var userModel = require('../models/userModel');

var LoginController = {};

//显示登陆页面
LoginController.login = function (req, res, next) {
	// var resData = {
	// 	code: 0,
	// 	data: {}
	// }
	// userModel.findAllUser(function(data){
	// 	resData.data = data.data;
	// 	res.render('Login/login', resData);
	// })
	userModel.findAllUser(function(data) {
		res.render('Login/login');
	});
 }

//实现登陆
LoginController.dologin = function(req, res, next) {
	var userid = req.body.userid;
	var password = req.body.password;

	userModel.findUser(userid, password, function(resData) {
		
		if(resData.code != 0) {
			req.session.user = resData.userid;
			console.log(req.session.user);
			res.redirect('/index');
		}else {
			res.redirect('/login');
		}
	});
}

//显示注册页面
LoginController.register = function(req, res, next) {
	res.render('Login/register');
}

//实现注册
LoginController.doregister = function(req, res, next) {
	var id = req.body.id;
	var name = req.body.name;
	var passwd = req.body.passwd;

	userModel.insertUser(id, name, passwd, function(resData) {
		if (resData.code != 0) {
			res.redirect('/login');
		}else {
			res.redirect('/register');
		}
	});
}


module.exports = LoginController;
