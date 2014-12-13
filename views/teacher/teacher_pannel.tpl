{{{template "inc/header.tpl" .}}}
{{{asset "js/angular.min.js"}}}
{{{asset "js/angular-route.min.js"}}}
{{{asset "js/teacher.js"}}}
{{{asset "sass/pannel.scss"}}}

	<div class="pannel">
		{{{template "teacher/inc_menu.tpl" .}}}

		<div class="main-pannel" ng-app="tpApp" >
			<div ng-controller="TeacherCtrl" class="problem-list" ng-view>

			</div>
		</div>
	</div>
	
{{{template "inc/footer.tpl" .}}}