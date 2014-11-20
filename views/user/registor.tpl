{{template "inc/header.tpl" .}}
<body>
	<form action="/registor">
		<input type="text" name="username" id="">
		<input type="password" name="password" id="">
		<input type="button" value="{{i18n "registor"}}">
	</form>
</body>
{{template "inc/footer.tpl" .}}