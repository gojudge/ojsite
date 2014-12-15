<div class="left-menu">
	<ul>
		<li>教师面板</li>
		<li><a href="">题库管理</a></li>
		<li><a href="">试题管理</a></li>
		<li><a href="">人工判题</a></li>
	</ul>
</div>
<script>
$(document).ready(function (e) {
	var host = window.location.host;
	var href = window.location.href;

	var uri = href.replace("http://","").replace(host,"");
	console.log("uri",uri);

	if (uri=="/teacher") {
		$(".left-menu>ul>li:eq(1)").addClass("current");
	} else if(uri == "/teacher/exam"){
		$(".left-menu>ul>li:eq(2)").addClass("current");
	}else if (uri == "/teacher/manjudge") {
		$(".left-menu>ul>li:eq(3)").addClass("current");
	} else{
		console.log("unknow uri",uri,"could not locate current tag");
	};
	
});
	
</script>