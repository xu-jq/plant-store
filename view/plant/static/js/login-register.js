function register(){
    let name1 = $('#name1').val();
    let name2 = $('#name2').val();
    let email = $('#email').val();
    let phone = $('#phone').val();
    let pwd = $('#pwd').val();
    let pwd2 = $('#pwd2').val();
    $.ajax({
        dataType: "json",
        // 提交方式 get 或 post
        type: "post",
        // 需要访问的 Servlet 的映射路径 urlPatterns
        url:  "http://localhost:8080/user/register",
        data: {
            name1: name1,
            name2: name2,
            email: email,
            phone: phone,
            pwd: pwd
        },
        success: function(data) {
            if (data.Flag){
                alert(data.Msg);
                setTimeout(flash());
            }else {
                alert("注册失败")
            }
        }

    });
}

function login(){
    let username = $('#login_name').val();
    let password = $('#password').val();
    $.ajax({
        dataType: "json",
        // 提交方式 get 或 post
        type: "post",
        // 需要访问的 Servlet 的映射路径 urlPatterns
        url:  "http://localhost:8080/user/login",
        data:{
            username:username,
            password:password
        },
        success: function(data){
            if (data.Flag){
                alert(data.Msg);
                setTimeout(flash());
            }else {
                alert("登录失败")
            }
        }
    });
}