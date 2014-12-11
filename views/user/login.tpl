{{{template "inc/header.tpl" .}}}
{{{asset "sass/login.scss"}}}
	<div class="info">&nbsp</div>
	<div class="login">
		<form action="/login">
			<ul>
				<li><label for="">{{{i18n "username"}}}</label><input type="text" name="username" id=""></li>
				<li><label for="">{{{i18n "password"}}}</label><input type="password" name="password" id=""></li>
				<li class="login-btns">
					<button class="btn">{{{i18n "login"}}}</button>
					<a class="btn oauth-login" href="https://github.com/login/oauth/authorize?client_id={{{.github_client_id}}}&scope=user,public_repo" target="_blank"><span class="octicon octicon-logo-github"></span>{{{i18n "login"}}}</a>
				</li>
			</ul>
		</form>
	</div>
{{{asset "js/user.js"}}}
{{{template "inc/footer.tpl" .}}}