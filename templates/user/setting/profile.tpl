{{{template "inc/header.tpl" .}}}
{{{asset "sass/user.scss"}}}
	<div class="pannel">
		{{{template "user/setting/inc_menu.tpl" .}}}

		<div class="main-pannel" ng-app="tpApp" >
			<ul class="profile">
				<li>Public profile</li>
				<li>
					<div class="fileinput-button">
						<img src="http://gravatar.duoshuo.com/avatar/{{{.email_md5}}}" alt="">
						<input type="file" name="avatar" id="">
					</div>
					<div>
						<span>upload you avatar</span>
					</div>
				</li>
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