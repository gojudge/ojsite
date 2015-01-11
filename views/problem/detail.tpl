{{{template "inc/header.tpl" .}}}
{{{asset "sass/problem.scss"}}}

{{{asset "ace/src-min-noconflict/ace.js"}}}

	<div class="detail">
		<div class="problem-title"><h3>{{{.problem_title}}}</h3></div>
		<div class="problem-description">
			{{{str2html .problem_description}}}
		</div>
		<div class="problem-links">discuss</div>
	</div>

	<form action="/problem/submit" method="post" class="submition">
		<input type="hidden" name="pid" value="{{{.problem_id}}}">
		<div class="subm-header">
			<select name="language" id="">
				<option value="C">C</option>
			</select>
			<a class="btn" id="reset" title="reset"><i class="icon-arrows-cw"></i></a>
			<input type="hidden" name="type" value="assert">
		</div>
		<div class="subm-editor">
			
			<textarea id="code_editor_raw" style="display:none;">{{{.problem_pre_code}}}</textarea>
			<div id="editor" name="code">{{{.problem_pre_code}}}</div>

		</div>
		<div class="solu-submit">
			<button class="btn"><i class="icon-goj"></i>Submit</button>
		</div>
	</form>

{{{asset "js/problem/detail.js"}}}
{{{template "inc/footer.tpl" .}}}
