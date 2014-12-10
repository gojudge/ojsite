$(".CodeMirror").ready(function() {
	// 初始化字体
	var fontsize = 12;
	var fontfamily = "consolas";

	if (typeof Setting.fontsize != 'undefined') {
		fontsize = Setting.fontsize;
	} else {
		Setting.fontsize = fontsize;
	}

	if (typeof Setting.fontfamily != 'undefined') {
		fontfamily = Setting.fontfamily;
	} else {
		Setting.fontfamily = fontfamily;
	}

	$(".CodeMirror").css({
		"font-family" : fontfamily,
		"font-size" : fontsize + "px"
	});

	instance.refreshEditors();
});