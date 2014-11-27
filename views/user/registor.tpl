{{template "inc/header.tpl" .}}
{{asset "sass/registor.scss"}}
	<div class="info info-success">here is the message...</div>
	<div class="registor">
		<form action="/registor" >
			<ul>
				<li><label for="">用户名</label><input type="text" name="username" id=""><span>用户名，登陆用，必填</span></li>
				<li><label for="">密码</label><input type="password" name="password" id=""><span>密码，登陆用，必填</span></li>
				<li><label for="">确认密码</label><input type="password" name="password" id=""><span>再输一遍密码</span></li>
				<li><label for="">邮箱</label><input type="text" name="email" id=""><span>邮箱，必填</span></li>
				<li><button class="btn">{{i18n "registor"}}</button>
			</ul>
		</form>
	</div>
{{template "inc/footer.tpl" .}}