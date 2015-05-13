{{{template "inc/header.tpl" .}}}
{{{asset "sass/login.scss"}}}
	<div class="pannel">
		{{{template "user/setting/inc_menu.tpl" .}}}

		<div class="main-pannel" ng-app="tpApp" >
			<ul>
				<li>Public profile</li>
				<li>
					<div>
						<img src="http://gravatar.duoshuo.com/avatar/{{{.email_md5}}}" alt="">
					</div>
					<div>
						<input type="file" name="avatar" id="">
						<span>upload you avatar</span>
					</div>
				</li>
			</ul>

			<ul>
				<li>
					<span>username</span>
					<input type="text" name="username" id="">
				</li>
				<li>
					<span>nickname</span>
					<input type="text" name="nickname" id="">
				</li>
				<li>
					<span>email</span>
					<input type="email" name="email" id="">
				</li>
			</ul>

			
		</div>

	</div>
{{{asset "js/user.js"}}}
{{{template "inc/footer.tpl" .}}}