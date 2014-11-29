{{template "inc/header.tpl" .}}
{{asset "sass/register.scss"}}
	<div class="info info-success">here is the message...</div>
	<div class="register">
		<form action="/register" method="post" >
			<ul>
				<li><label for="">{{i18n "username"}}</label><input type="text" name="username" id=""><span>用户名，登陆用，必填</span></li>
				<li><label for="">{{i18n "password"}}</label><input type="password" name="password" id=""><span>密码，登陆用，必填</span></li>
				<li><label for="">{{i18n "password_again"}}</label><input type="password" id=""><span>再输一遍密码</span></li>
				<li><label for="">{{i18n "email"}}</label><input type="text" name="email" id=""><span>邮箱，必填</span></li>
				<li><button class="btn">{{i18n "register"}}</button>
			</ul>
		</form>
	</div>
{{asset "js/user.js"}}
{{template "inc/footer.tpl" .}}