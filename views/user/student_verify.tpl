{{template "inc/header.tpl" .}}
{{asset "sass/login.scss"}}
	<div class="info info-success">here is the message...</div>
	<div class="login">
		<form action="/login">
			<ul>
				<li><label for="">学号</label><input type="text" name="username" id=""></li>
				<li><label for="">密码</label><input type="password" name="password" id=""></li>
				<li class="login-btns">
					<button class="btn"><span class="octicon octicon-key"></span>验证</button>
				</li>
			</ul>
		</form>
	</div>
{{asset "js/user.js"}}
{{template "inc/footer.tpl" .}}