function ajax(method, url, data, cb) {
	var xhr;
	if(window.XMLHttpRequest) {
		xhr = new XMLHttpRequest();
	}else {
		xhr = new ActiveXObject("Microsoft.XMLHTTP");
	}

	xhr.onreadystatechange = function() {
		if(xhr.readyState == 4 && xhr.status == 200) {
			cb(xhr.responseText);
		}
	}
	xhr.open(method, url);
	xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
	var str = [];
	for(var key in data) {
		str.push(key + "=" + data[key]);
	}
	xhr.send(str.join("&"));
}



function Circle(data) {
	var elem = null;
	var hasElem = this.pool.some(function(item, index) {
		if(item.name == data.name && item.code == data.code) {
			elem = item;
			return true;
		}
	});

	if(!hasElem) {
		this.x = this.accumulate.x;
		this.y = this.accumulate.y;

		this.dom = null;

		this.name = data.name;
		this.code = data.code;
		this.rank = 0;

		this.create(data);	
	}else {
		if(Circle.prototype.isSafeTrace) {
			console.log("safe")
			Circle.prototype.accumulate.x -= 150;
			new Path(this.storePath.start, elem);
			Circle.prototype.storePath.start = elem;
		}else {
			Circle.prototype.accumulate.x += 150;
			Circle.prototype.storePath.start = elem;
			
		}
		return elem;
	}
	
}
Circle.prototype.init = function() {
	console.log("init")
	this.accumulate = {
		x: 100,
		y: 80
	};
	this.storePath = {
		start: null
	};
	Circle.prototype.isSafeTrace = 0;
	this.pool = [];
};
Circle.prototype.addLine = function() {
	console.log("addLine")
	this.accumulate = {
		x: 100,
		y: this.accumulate.y += 150
	};
};
Circle.prototype.accumulate = {
	x: 100,
	y: 100
};
Circle.prototype.pool = [];
Circle.prototype.isSafeTrace = 0;
Circle.prototype.create = function(data) {

	var circle = document.createElement("div");
	circle.className = "circle";
	if(!Circle.prototype.isSafeTrace) {
		this.accumulate.x = this.accumulate.x + 150;
	}else {
		this.accumulate.x = this.accumulate.x - 150;
	}
	circle.style.left = this.x + "px";
	circle.style.top = this.y + "px";
	circle.title = "id: " + data.code;
	circle.innerHTML = data.name;
	map.appendChild(circle);
	this.dom = circle;

	if(this.storePath.start == null) {
		this.storePath.start = this;
	}else {
		new Path(this.storePath.start, this);
		this.storePath.start = this;
	}
};


Circle.prototype.storePath = {
	start: null
}

Circle.prototype.showInfo = function(e) {
	var target = e.target;
	// 显示具体信息
	if(target.className == 'circle') {
	    if(!Circle.prototype.infoPool[target.title]) {
	    	var infodiv = document.createElement('div');
		    infodiv.className = 'circle-info';
		    infodiv.innerHTML = '<div class="info-arrow"></div><div class="info-content"></div>';

			infodiv.style.left = target.offsetLeft - 0 + 'px';
			infodiv.style.top = target.offsetTop + 70 + 'px';

			var name = target.innerHTML,
				code = target.title.slice(3).trim();
			var url = "";
			switch(name) {
				case "种子" :
					url = '/api/getseedinfo/' + code;
					break;
				case "秧苗" :
					url = '/api/getseedlinginfo/' + code;
					break;
				case "种植" :
					url = '/api/getplantinfo/' + code;
					break;
				case "收割入库" :
					url = '/api/getinputinfo/' + code;
					break;
				case "仓库" :
					url = '/api/getwarehouseinfo/' + code;
					break;
				case "进料" :
					url = '/api/getfeedinfo/' + code;
					break;
				case "原料" :
					url = '/api/getmaterialinfo/' + code;
					break;
				case "产品" :
					url = '/api/getproductinfo/' + code;
					break;
			}

			var str = "";
			
			ajax('get', url, {}, function(data) {
				var bj=JSON.parse(data)
				if(bj instanceof Array){//判断对象是否是Array的实例
					data = bj[0];
				} else{
					data=bj;
				}
				
				console.log(data);
				for(var key in data) {
					str += "<li>" + key + ": " + data[key] + "</li>"
				}
				infodiv.querySelector(".info-content").innerHTML = str;
				infodiv.style.visibility = 'visible';
				map.appendChild(infodiv);
				Circle.prototype.infoPool[target.title] = infodiv;
			});
	    }else {
	    	map.removeChild(Circle.prototype.infoPool[target.title]);
	    	Circle.prototype.infoPool[target.title] = null;
	    }

	}
	// else {
	// 	infodiv.style.visibility = 'hidden';
	// }
}

Circle.prototype.infoPool = {};

function Path(start, end) {
	this.start = start;
	this.end = end;

	this.create(start, end);
}
Path.prototype.create = function(start, end) {
	var path = document.createElement("span");
	path.className = "path";

	var deg = (end.y  - start.y) / (end.x - start.x);

	path.style.width = Math.hypot(end.y - start.y, end.x - start.x) - 80 + 'px';
	path.style.left = start.x + 40 + 'px';
	path.style.top = start.y + 40 + 'px';

	if(!Circle.prototype.isSafeTrace) {
		path.style.transform = "rotate(" + Math.atan(deg) * 180 / Math.PI + "deg)" + 
							"translate(40px, 0px)";	
	}else{
		path.style.transform = "rotate(" + (parseInt(Math.atan(deg) * 180 / Math.PI) + 180) + "deg)" + 
							"translate(40px, 0px)";	
	}

	path.style.transformOrigin = "0 0";
	
	map.appendChild(path)
}

var transformKey = {
	seedid: "种子",
	seedlingid: "秧苗",
	plantid: "种植",
	inputid: "收割入库",
	warehouseid: "仓库",
	feedid: "进料",
	materialid: "原料",
	productid: "产品"
};


// 数字越大，越有问题，颜色越深
var safeRank = {
	"0": "#fff",
	"1": "green",
	"2": "yellow",
	"3": "red",
	"4": "blue"
}