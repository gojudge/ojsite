{{template "inc/header.tpl" .}}
{{asset "sass/login.scss"}}
	<div class="info">&nbsp</div>
	<div class="login binding-confirm">
		<form action="/oauth/github" method="post">
			<ul>
				<li>您即将绑定<span class="octicon octicon-logo-github"></span>帐号<span>{{.github_username}}</span>，请确认</li>
				<li>
					<div class="avatars">
						<img class="github-avatar" src="{{.github_avatar}}" alt="">
						<img class="goj-avatar" src="{{.goj_avatar}}" alt="">
					</div>
				</li>
				<li><label for="">GOJ&nbsp{{i18n "password"}}</label><input type="password" name="password" id=""></li>
				<li class="login-btns">
					<input type="hidden" name="token" value="{{.token}}">
					<button class="btn">确定</button>
				</li>
			</ul>
		</form>
	</div>
{{asset "js/user.js"}}
{{template "inc/footer.tpl" .}}