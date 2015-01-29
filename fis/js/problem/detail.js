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

	function update_status(color,cls,icon,msg){
		$(".solu-submit>span").show();
		$(".solu-submit>span>a").html(msg);
		$(".solu-submit>span").removeClass().addClass("status "+cls).css("background-color",color);
		$(".solu-submit>span>i").removeClass().addClass(icon);
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
				update_status("black","","icon-spin4 animate-spin","Waiting...");
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
											update_status("green","","icon-ok","Accept");
										} else if(json.status == "RE"){
											update_status("red","","icon-cancel","Runtime Error");
										}else if(json.status == "MLE"){
											update_status("red","","icon-cancel","Out Of Memory");
										}else if(json.status == "TLE"){
											update_status("red","","icon-cancel","Out Of Time");
										}else if(json.status == "OLE"){
											update_status("red","","icon-cancel","Output Limit");
										}else if(json.status == "PSF"){
											update_status("red","","icon-cancel","Syscal Forbidden");
										}else if (json.status == "CE") {
											update_status("red","","icon-cancel","Compile Error");
										}else{
											update_status("gray","","icon-cancel","Unknown");
										};
									}else{
										debugger
										update_status("black","","icon-spin4 animate-spin","Waiting...");
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


