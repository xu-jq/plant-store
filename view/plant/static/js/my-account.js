$(function(){
    $(".order-state li").click(function(){
        $(this).addClass("on").siblings().removeClass("on");
        var txt1=$(this).find("a").text();
        $(".dkuang").find("p.one").each(function(){
            var txt2=$(this).text();
            if(txt1==txt2){
                $(this).parent(".dkuang").show().siblings(".dkuang").hide();
            }
            $(".order-state li").eq(0).click(function(){
                $(".dkuang").show();
            })
        });
    });

    $(".forgetpassword").click(function () {
        $(".ymm").hide();
        $(".yxmm").show();
    })

});