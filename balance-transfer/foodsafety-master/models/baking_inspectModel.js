var baking_inspectModel = {};

//遍历baking_quality_inspect表
baking_inspectModel.findAllBaking = function(callback) {
	connection.query('select * from baking_quality_inspect', function (error, results, fields) {
		var resData = {
			code: 0,
			data: {}
		}
		if (error) {
			throw error;
		}else {
			resData.code = 1;
			resData.data = results;
			console.log('The first is: ', results);
			
		}
		callback(resData);
	});
}

module.exports = baking_inspectModel;