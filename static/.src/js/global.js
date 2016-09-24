var goj = {};

goj.close_page = function(){
	window.opener = null;
	window.open('','_self');
	window.close();
};

$(document).ready(function (e) {
	// set cookie
	function set_cookie(name,value){
	    var Days = 365; // 1 year
	    var exp  = new Date();
	    exp.setTime(exp.getTime() + Days*24*60*60*1000);
	    document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
	}

	// get cookie
	function get_cookie(objName){
		var arrStr = document.cookie.split("; ");
		for(var i = 0;i < arrStr.length;i ++){
			var temp = arrStr[i].split("=");
			if(temp[0] == objName) return unescape(temp[1]);
		}
	}

	// language choose options
	$(".lang").mouseover(function (e) {
		$(".lang>ul").show();
	}).mouseout(function (e) {
		$(".lang>ul").hide();
	});

	$(".lang>ul").delegate("a","click",function (e) {
		var lang = $(this).attr("lang");
		set_cookie("lang",lang);
		window.location.reload();
	});

	goj.lang = get_cookie("lang");
	$(".lang>ul>li>a[lang='" + goj.lang + "']").addClass("strong");

	// info box show & remove
	goj.info_danger = function (info) {
		$(".main>.info").html(info).removeClass().addClass("info info-danger");
	}

	goj.info_success = function (info) {
		$(".main>.info").html(info).removeClass().addClass("info info-success");
	}

	goj.info_remove = function () {
		$(".main>.info").html("&nbsp").removeClass().addClass("info");
	}
});

console.blue = function (text) {
	console.log("%c"+text, "color:blue;font-weight:bolder")
}

console.green = function (text) {
	console.log("%c"+text, "color:green;font-weight:bolder")
}

console.red = function (text) {
	console.log("%c"+text, "color:red;font-weight:bolder")
}

console.log("%c","padding:16px 24px;line-height:40px;background:url('http://oj.duguying.net/static/img/favicon.ico') no-repeat")
console.log("%cwelcome to Goj front-end debug console.","color:blue;font-weight:bolder;font-size:20px")
console.green("you can join us. QQ group: 235149571")