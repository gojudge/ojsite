<div class="left-menu">
	<ul>
		<li>账户设置</li>
		<li {{{if eq .nav "profile"}}}class="current"{{{end}}}><a href="/user/setting/profile">个人信息</a></li>
		<li {{{if eq .nav "pwd"}}}class="current"{{{end}}}><a href="/user/setting/pwd">修改密码</a></li>
		<li {{{if eq .nav "bind"}}}class="current"{{{end}}}><a href="/user/setting/bind">社交账号绑定</a></li>
	</ul>
</div>
