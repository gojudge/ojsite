{{template "inc/header.tpl" .}}
<body>
	<form action="/login">
		<input type="text" name="username" id="">
		<input type="password" name="password" id="">
		<input type="button" value="{{i18n "login"}}">
	</form>
</body>
{{template "inc/footer.tpl" .}}