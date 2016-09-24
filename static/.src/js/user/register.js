/**
 * @author Rex Lee
 * @license http://opensource.org/licenses/MIT MIT License
 * @description js for register page
 */


$(document).ready(function (e) {
	// form check
	var register = $(".register>form").validate({
			errorPlacement: function(error, element) {
				$(element).closest("form>ul>li")
					.find("label[for='" + element.attr( "id" ) + "']")
					.append(error);
			},
			rules: {
				username: {
					required: true,
					minlength: 2
				},
				password: {
					required: true,
					minlength: 5
				},
				confirm: {
					required: true,
					minlength: 5,
					equalTo: "#password"
				},
				email: {
					required: true,
					email: true
				}
			},
			messages: {
				confirm: {
					equalTo: $("input[name='confirm']").attr("if-not-equal")
				},
				
			},
			errorElement: "span",
			
			
		});
		
		//submit
		$(".register>form").submit(function (e) {
			if (!register.valid()){
				console.log("验证失败")
				return false
			}else{
				$.ajax({
					url: $(".register>form").attr("action"),
					method: "post",
					data: $(this).serialize(),
					dataType: "json",
					success: function(json){
						console.log(json);
						if (json.result) {
							goj.info_success("success");
							window.location = "/";
						} else{
							goj.info_danger("register failed");
							console.log(json);
						};
					}
				});
			}

			return false;
		})
});