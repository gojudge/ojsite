<div class="left-menu">
	<ul>
		<li>账户设置</li>
		<li><a href="">个人信息</a></li>
		<li><a href="">修改密码</a></li>
		<li><a href="">社交账号绑定</a></li>
	</ul>
</div>
<script>
$(document).ready(function (e) {
	var host = window.location.host;
	var href = window.location.href;

	var uri = href.replace("http://","").replace(host,"");
	console.log("uri",uri);

	if(uri == "/user/setpassword"){
		$(".left-menu>ul>li:eq(2)").addClass("current");
	}else if (uri == "/user/setoauth") {
		$(".left-menu>ul>li:eq(3)").addClass("current");
	}else{
		$(".left-menu>ul>li:eq(1)").addClass("current");
	};
	
});
	
</script>