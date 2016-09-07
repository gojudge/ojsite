{{{template "inc/header.tpl" .}}}
{{{asset "sass/login.scss"}}}
	<div class="info info-success">here is the message...</div>
	<div class="login">
		<span class="title">请使用长江大学教务处帐号验证</span>
		<form action="/login">
			<ul>
				<li class="school-logo"><i class="icon-yangtzeu"></i></li>
				<li><label for="">学号</label><input type="text" name="username" id=""></li>
				<li><label for="">密码</label><input type="password" name="password" id=""></li>
				<li class="login-btns">
					<button class="btn"><i class="octicon octicon-key"></i>验证</button>
				</li>
			</ul>
		</form>
	</div>
{{{asset "js/user.js"}}}
{{{template "inc/footer.tpl" .}}}