{{{template "inc/header.tpl" .}}}
{{{asset "js/angular.min.js"}}}
{{{asset "js/teacher.js"}}}
{{{asset "sass/pannel.scss"}}}

	<div class="pannel">
		{{{template "teacher/inc_menu.tpl" .}}}

		<div class="main-pannel" ng-app="tpApp" >

			<div class="list-banner">
				<ul>
					<li class="item">
						<ul >
							<li class="item-id">编号</li>
							<li class="item-title">标题</li>
							<li class="item-type">类型</li>
							<li class="item-time">时间</li>
							
							<li class="item-right" title="add">
								<a href="/problem/add" target="_blank">
									<i class="icon-plus"></i>
								</a>
							</li>
						</ul>
					</li>
				</ul>
			</div>

			<div ng-controller="ProblemListCtrl" class="problem-list">
				<ul>
					<li class="item">
						<ul>
							<li>题库列表</li>

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
							<li class="item-right" title="edit">
								<a href="/problem/edit/{{item.id}}" class="edit">
									<i class="icon-edit"></i>
								</a>
							</li>
							<li class="item-right" title="delete">
								<a href="javascript:void(0);" class="del" delete ng-model="item">
									<i class="icon-trash"></i>
								</a>
							</li>
						</ul>
					</li>
				</ul>
			</div>

			<div ng-controller="AuditCtrl" class="audit-list">
				<ul>
					<li class="item">
						<ul>
							<li>待审题目列表</li>

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
							<li class="item-right" title="check" check>
								<a href="javascript:void(0);" class="check" accept ng-model="item">
									<i class="icon-check"></i>
								</a>
							</li>
							<li class="item-right" title="cancel">
								<a href="javascript:void(0);" class="del" deny ng-model="item">
									<i class="icon-cancel"></i>
								</a>
							</li>
						</ul>
					</li>
				</ul>
			</div>

		</div>
	</div>
	
{{{template "inc/footer.tpl" .}}}