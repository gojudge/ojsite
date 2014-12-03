{{template "inc/header.tpl" .}}
	<div class="announcement">
	</div>
	<div class="problem-list">
		<ul>
			{{range $i,$problem := .problems}}
			<li>
				{{$problem.id}}|{{$problem.type}}|{{$problem.description}}|{{$problem.pre_code}}|{{$problem.io_data}}|{{$problem.tags}}|{{$problem.level}}
			</li>
			{{end}}
		</ul>
	</div>
{{template "inc/footer.tpl" .}}