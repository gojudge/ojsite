{{template "inc/header.tpl" .}}
{{asset "sass/problem.scss"}}

{{asset "codemirror/lib/codemirror.css"}}
{{asset "codemirror/lib/codemirror.js"}}
{{asset "codemirror/mode/clike/clike.js"}}
{{asset "codemirror/addon/edit/matchbrackets.js"}}


	<div class="detail">
		<div class="problem-title"><h3>{{.problem_title}}</h3></div>
		<div class="problem-description">
			{{.problem_description}}
		</div>
		<div class="problem-links">discuss</div>
	</div>

	<form action="/problem/submit" method="post" class="submition">
		<div class="subm-header">
			<select name="language" id="">
				<option value="C">C</option>
			</select>
			<button class="btn"><i class="icon-arrows-cw"></i></button>
			<input type="hidden" name="type" value="assert">
		</div>
		<div class="subm-editor">

			<textarea id="code_editor" name="code">{{.problem_pre_code}}</textarea>

		</div>
		<div class="solu-submit">
			<button class="btn"><i class="icon-goj"></i>Submit</button>
		</div>
	</form>

{{asset "js/problem.js"}}
{{template "inc/footer.tpl" .}}
