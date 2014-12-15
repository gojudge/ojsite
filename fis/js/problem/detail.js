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

	// sbmit task
	$("form.submition").submit(function(e){
		$.ajax({
			url: $(this).attr("action"),
			method: "post",
			data: $(this).serialize(),
			dataType: "json",
			success: function(json){
				console.log(json);
			}
		});
		return false;
	});

	$("#reset").click(function(e){
		var raw_content = $("#code_editor_raw").val();
		editor.setValue(raw_content);
	});
});


