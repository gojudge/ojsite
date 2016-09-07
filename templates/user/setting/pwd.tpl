{{{template "inc/header.tpl" .}}}
{{{asset "sass/login.scss"}}}
	<div class="pannel">
		{{{template "user/setting/inc_menu.tpl" .}}}

		<div class="main-pannel" ng-app="tpApp" >
			<ul>
				<li>Change Password</li>
				<li>
					<form action="">
						<input type="password" name="" id="">
						<input type="password" name="" id="">
						<input type="submit" value="修改">
					</form>
				</li>
			</ul>

		</div>

	</div>
{{{asset "js/user.js"}}}
{{{template "inc/footer.tpl" .}}}