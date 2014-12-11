{{{template "inc/header.tpl" .}}}
{{{asset "js/angular.min.js"}}}
{{{asset "js/teacher.js"}}}


	<div class="pannel">
		{{{template "teacher/inc_menu.tpl" .}}}

		<div class="main-pannel" ng-app >
			
			<p ng-controller="TeacherController">{{someText}}</p>
		</div>
	</div>
{{{template "inc/footer.tpl" .}}}
