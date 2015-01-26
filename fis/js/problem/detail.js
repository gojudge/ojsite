$(document).ready(function(e){
	
	editor = ace.edit("editor");
    editor.setTheme("ace/theme/tomorrow");
    editor.getSession().setMode("ace/mode/c_cpp");
    editor.setOption("minLines", 20);
    editor.setOption("maxLines", 20);

	editor.on("change", function() {
		var content = editor.getValue();

		$("#code_editor").html(content);
	});

	// submit task
	$("form.submition").submit(function(e){
		var pid = $("input[name='pid']").val();
		var language = $("select[name='language']").val();
		var ptype = $("input[name='type']").val();
		var code = editor.getValue();
		// console.log(pid, language, ptype, code);

		$.ajax({
			url: $(this).attr("action"),
			method: "post",
			data: {"pid":pid, "language":language, "ptype":ptype, "code":code},
			dataType: "json",
			success: function(json){
				console.log(json);
				if (json.result) {
					// get status after 1s
					window.setInterval(function(){
						$.ajax({
							url: $("form.submition").attr("get-status"),
							method: "get",
							data: {"sbid":json.id},
							dataType: "json",
							success: function(json){
								// debugger;
								console.log(json);
							}
						});
					}, 1000);
				};
			}
		});
		return false;
	});

	$("#reset").click(function(e){
		var raw_content = $("#code_editor_raw").val();
		editor.setValue(raw_content);
	});

	
});


