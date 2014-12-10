$(document).ready(function(e){
	editor = CodeMirror.fromTextArea(document.getElementById("code_editor"), {
		lineNumbers: true,
		mode: "text/x-csrc",
		matchBrackets: true
	});

	$(".CodeMirror").ready(function(e){
		$(".CodeMirror").css({
			"font-family" : "consolas",
			"font-size" : "14px"
		});

		editor.refresh();

		// console.log(content);
	});

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


