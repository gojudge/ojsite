{{template "inc/header.tpl" .}}
{{asset "sass/login.scss"}}
	<div class="info info-success">here is the message...</div>
	<div class="login">
		<form action="/login">
			<ul>
				<li><label for="">用户名</label><input type="text" name="username" id=""></li>
				<li><label for="">密码</label><input type="password" name="password" id=""></li>
				<li class="login-btns">
					<button class="btn">{{i18n "login"}}</button>
					<a class="btn oauth-login" href="https://github.com/login/oauth/authorize?client_id={{.github_client_id}}&scope=user,public_repo" target="_blank"><span class="octicon octicon-logo-github"></span>登录</a>
				</li>
			</ul>
		</form>
	</div>
{{asset "js/user.js"}}
{{template "inc/footer.tpl" .}}