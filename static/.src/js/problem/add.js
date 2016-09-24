$(document).ready(function(e){
	var opts = {
		container: 'epiceditor',
		textarea: 'epiceditor',
		basePath: '/static/epiceditor',
		clientSideStorage: true,
		localStorageName: 'epiceditor',
		useNativeFullscreen: true,
		parser: function (str) {
			var rst = $.ajax({
				async: false,
				type: "post",
				url: "/api/markdown/preview",
				data: {"content": str},
				dataType:"json",

			}).responseText;
			var json = eval('('+rst+')')
			console.log(json);
			return json.preview;
		},
		file: {
			name: 'epiceditor',
			defaultContent: '',
			autoSave: 100
		},
		theme: {
			base: '/themes/base/epiceditor.css',
			preview: '/themes/preview/github.css',
			editor: '/themes/editor/epic-light.css'
		},
		button: {
			preview: true,
			fullscreen: true,
			bar: "auto"
		},
		focusOnLoad: false,
		shortcut: {
			modifier: 18,
			fullscreen: 70,
			preview: 80
		},
		string: {
			togglePreview: 'Toggle Preview Mode',
			toggleEdit: 'Toggle Edit Mode',
			toggleFullscreen: 'Enter Fullscreen'
		},
		autogrow: false,
		highlight: true,
	}
	mdeditor = new EpicEditor(opts).load();

	function checkbox_val(selector){
		var s=''; 
		$(selector+":checked").each(function(){ 
			s+=$(this).val()+','; 
		}); 

		if (s.length > 0) { 
			s = s.substring(0,s.length - 1); 
		}
		
		return s;
	}

	$(".item>form").submit(function (e) {
		var title = $("input[name='title']").val();
		var type = checkbox_val("input[name='type']");
		var description = mdeditor.exportFile();
		var precode = editor.getValue();
		var input = $(".io-data>textarea[name='input']").val();
		var output = $(".io-data>textarea[name='output']").val();
		var email = $(".author-info>input[name='author-email']").val();

		var form_data = {
				"title":title,
				"type":type,
				"description":description,
				"precode":precode,
				"input":input,
				"output":output,
				"email":email,
			}

		$.ajax({
			type: $(this).attr("method"),
			url: $(this).attr("action"),
			data: form_data,
			dataType:"json",
			success: function(json){
				if (json.result) {
					alert("add success");
					window.location.href="/problems";
				} else{
					alert("add failed");
				};
			}
		});

		return false;
	});
	
});