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

	// parms: TA, ER, AC
	function update_status(status){
		$(".solu-submit>span").show();
		if (status == "TA") { // task add
			$(".solu-submit>span>a").html("Waiting");
			// status
			$(".solu-submit>span").removeClass().addClass("status sta-task-add");
			// icon
			$(".solu-submit>span>i").removeClass().addClass("icon-spin4 animate-spin");
		}else if (status == "ER") { // error
			$(".solu-submit>span>a").html("Error");
			$(".solu-submit>span").removeClass().addClass("status sta-error");
			$(".solu-submit>span>i").removeClass().addClass("icon-cancel-1");
		} else if (status == "AC"){ // accept
			$(".solu-submit>span>a").html("Accept");
			$(".solu-submit>span").removeClass().addClass("status sta-accept");
			$(".solu-submit>span>i").removeClass().addClass("icon-ok");
		};
	}

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
				update_status("TA");
				if (json.result) {
					// get status after 1s
					var tc = window.setInterval(function(){
						$.ajax({
							url: $("form.submition").attr("get-status"),
							method: "get",
							data: {"sbid":json.id},
							dataType: "json",
							success: function(json){
								console.log(json);
								if (json.result) {
									// if get result clearInterval
									if (json.status!="TA"){
										window.clearInterval(tc);
										if (json.status == "AC") {
											update_status("AC");
										} else{
											update_status("ER");
										};
									}else{
										update_status("TA");
									}
								};
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


