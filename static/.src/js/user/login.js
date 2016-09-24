$(document).ready(function (e) {
	$(".login>form").submit(function (e) {
		$.ajax({
			url: $(".login>form").attr("action"),
			method: "post",
			data: $(this).serialize(),
			dataType: "json",
			success: function(json){
				console.log(json);
				if (json.result) {
					goj.info_success("success");
					window.location = "/";
				} else{
					goj.info_danger("login failed");
					console.log(json);
				};
			}
		});
		return false;
	});
});