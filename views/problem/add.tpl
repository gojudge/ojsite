{{{template "inc/header.tpl" .}}}
{{{asset "sass/problem.scss"}}}

<script type="text/javascript" src="http://ueditor.baidu.com/ueditor/ueditor.config.js"></script>
<script type="text/javascript" src="http://ueditor.baidu.com/ueditor/ueditor.all.js"></script>

{{{asset "codemirror/lib/codemirror.css"}}}
{{{asset "codemirror/lib/codemirror.js"}}}
{{{asset "codemirror/mode/clike/clike.js"}}}
{{{asset "codemirror/addon/edit/matchbrackets.js"}}}

	<div class="problem-add">
		<ul class="box-add">
			<li class="item">
				<ul>
					<li><h3>添加题目</h3></li>
				</ul>
			</li>
			<li class="item">
				<form action="/problem/add" method="post">
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
					<label for="" class="item-title">题目描述</label>
					<!-- 加载编辑器的容器 -->
					<script id="container" name="content" type="text/plain">
					</script>
	<script type="text/javascript" charset="utf-8" src="http://ueditor.baidu.com/ueditor/ueditor.config.js"></script>
	<script type="text/javascript" charset="utf-8" src="http://ueditor.baidu.com/ueditor/ueditor.all.js"></script>
	<script type="text/javascript" charset="utf-8" src="http://ueditor.baidu.com/ueditor/lang/zh-cn/zh-cn.js"></script>
	<script type="text/javascript" charset="utf-8" src="http://ueditor.baidu.com/ueditor/lang/en/en.js"></script>

					<script type="text/javascript">
					    var ue = UE.getEditor('container');
					</script>

					<div class="precode">
						<label for="" class="item-title">预代码</label>
						<textarea id="code_editor_raw" style="display:none;"></textarea>
						<textarea id="code_editor" name="code"></textarea>
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