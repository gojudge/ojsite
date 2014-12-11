{{{template "inc/header.tpl" .}}}
{{{asset "js/angular.min.js"}}}
{{{asset "js/teacher.js"}}}
{{{asset "sass/pannel.scss"}}}

	<div class="pannel">
		{{{template "teacher/inc_menu.tpl" .}}}

		<div class="main-pannel" ng-app >
			<div ng-controller="TeacherController" class="problem-list">
				<ul>
					<li class="item">
						<ul>
							<li class="item-id">编号</li>
							<li class="item-title">标题</li>
							<li class="item-type">类型</li>
							<li class="item-time">时间</li>
							<li class="item-right item-nav">
								<a ng-click="nextPage()" ng-class="nextPageDisabled()"><i class="icon-right-open"></i></a>
							</li>
							<li class="item-right item-nav">
								<a ng-click="prevPage()" ng-class="prevPageDisabled()"><i class="icon-left-open"></i></a>
							</li>
						</ul>
					</li>
					<li ng-repeat="item in data.list" class="item">
						<ul>
							<li class="item-id">{{item.id}}</li>
							<li class="item-title">{{item.title}}</li>
							<li class="item-type">{{item.type}}</li>
							<li class="item-time">{{item.time}}</li>
							<li class="item-right">
								<a href="/problem/edit/{{item.edit}}"><i class="icon-edit"></i></a>
							</li>
							<li class="item-right">
								<a href="/problem/del/{{item.id}}"><i class="icon-trash"></i></a>
							</li>
						</ul>
					</li>
					
				</ul>
			</div>
		</div>
	</div>
{{{template "inc/footer.tpl" .}}}
