{{{template "inc/header.tpl" .}}}
{{{asset "sass/problem.scss"}}}
	<div class="announcement" style="display:none;"></div>
	<div class="problem">
		<div class="problem-list">
			<ul>
				<li class="list-header">
					<ul>
						<li class="item-id">编号</li>
						<li class="item-title">题目</li>
						<li class="item-type">类型</li>
						<li class="item-rate">通过率</li>
						<li class="item-time">添加时间</li>
					</ul>
				</li>
				{{{range $i,$problem := .problems}}}
				<li class="list-item">
					<ul>
						<li class="item-id">{{{$problem.id}}}</li>
						<li class="item-title">
							<a href="/problem/{{{$problem.title}}}">{{{$problem.title}}}</a>
						</li>
						<li class="item-type">{{{$problem.type}}}</li>
						<li class="item-rate">{{{$problem.pass_rate}}}</li>
						<li class="item-time">{{{date $problem.time}}}</li>
					</ul>
				</li>
				{{{end}}}
			</ul>
		</div>
		<div class="cata">
			<ul class="hot">
				<li><i class="icon-fire"></i>Top 10</li>
				{{{range $i,$pro := .top10}}}
				<li class="item"><a href="">{{{$pro.title}}}</a></li>
				{{{end}}}
			</ul>
			<ul class="tags">
				<li><i class="icon-tags"></i>分类</li>
				{{{range $i,$t := .tag_list}}}
				<li class="item"><a href="">{{{$t.tag}}}<span class="badge">{{{$t.problem_num}}}</span></a></li>
				{{{end}}}
			</ul>
			<ul class="tags">
				<li><i class="icon-link"></i>链接</li>
				
				<li class="item"><a href="/problem/add" target="_blank">我要添加题目</a></li>
				
			</ul>
		</div>
	</div>
	
{{{template "inc/footer.tpl" .}}}