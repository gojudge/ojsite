{{{template "inc/header.tpl" .}}}
{{{asset "sass/problem.scss"}}}

{{{asset "ace/src-min-noconflict/ace.js"}}}
{{{asset "epiceditor/js/epiceditor.min.js"}}}

	<div class="problem-add">
		<ul class="box-add">
			<li class="item">
				<ul>
					<li><h3>添加题目</h3></li>
				</ul>
			</li>
			<li class="item">
				<form action="/api/problem/add" method="post">
					<span class="item-add-header">
						<label for="">标题</label><input type="text" name="title" id="" class="title">
						<div>
							<label class="option-right checkbox">
								<input name="type" type="checkbox" value="io" />io
							</label>
							<label class="option-right checkbox">
								<input name="type" type="checkbox" value="assert" />assert
							</label>
							{{{if eq .userIs "teacher"}}}
							<label class="option-right radio">
								<input name="level" type="radio" value="private" />private
							</label>
							{{{end}}}
						</div>
					</span>
					<label for="" class="item-title">题目描述<small>(支持Markdown语法，不支持HTML)</small></label>
					
					<div class="goj-editor">
						<div id="epiceditor"></div>
					</div>

					<div class="precode">
						<label for="" class="item-title">预代码</label>

						<textarea id="code_editor_raw" style="display:none;">{{{.problem_pre_code}}}</textarea>
					
						<pre id="editor"></pre>
					
					</div>

					<div class="io-data">
						<span class="item-title">IO数据</span>
						<label for="">INPUT</label>
						<textarea name="input" id="" cols="30" rows="10"></textarea>
						<label for="">OUTPUT</label>
						<textarea name="output" id="" cols="30" rows="10"></textarea>
					</div>

					{{{if eq .userIs "guest"}}}
					<div class="author-info">
						<span class="item-title">作者信息</span>
						<p>感谢您做出的贡献，请留下联系方式</p>
						<label for="">Email</label>
						<input type="text" name="author-email" id="">
					</div>
					{{{end}}}

					<input class="btn" type="submit" value="添加">
				</form>
			</li>
		</ul>
	</div>

{{{asset "js/problem.js"}}}
{{{template "inc/footer.tpl" .}}}