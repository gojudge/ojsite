{{{template "inc/header.tpl" .}}}
{{{asset "sass/login.scss"}}}
	<div class="pannel">
		{{{template "user/setting/inc_menu.tpl" .}}}

		<div class="main-pannel" ng-app="tpApp" >
			<ul>
				<li>
					<span>Bind Github</span>
					<input type="text" name="github" id="">
					<input type="submit" value="绑定">
				</li>
				<li>
					<span>Bind OSC</span>
					<input type="text" name="osc" id="">
					<input type="submit" value="绑定">
				</li>
			</ul>

		</div>

	</div>
{{{asset "js/user.js"}}}
{{{template "inc/footer.tpl" .}}}