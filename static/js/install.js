var area = {
    submit: function(){
        $.post("/install/do_submit",{
            dburl: $("[name=db-url]").val(),
            dbport: $("[name=db-port]").val(),
            dbuser: $("[name=db-user]").val(),
            dbpwd: $("[name=db-pwd]").val(),
            dbname: $("[name=db-name]").val(),
            adminuser: $("[name=admin-user]").val(),
            adminpwd: $("[name=admin-pwd]").val()
        },function(data){
            if(data.success){
                location.href="/"
            }else{

            }

        })
    }
}
