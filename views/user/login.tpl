{{template "inc/header.tpl" .}}
	<form action="/login">
		<input type="text" name="username" id="">
		<input type="password" name="password" id="">
		<input type="button" value="{{i18n "login"}}">
	</form>
	<div>
		<a href="https://github.com/login/oauth/authorize?client_id={{.github_client_id}}&scope=user,public_repo" target="_blank">Github登录</a>
	</div>
{{template "inc/footer.tpl" .}}