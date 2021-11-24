
var msgdsq;
//错误时：提示调用方法
function show_err_msg(msg) {
    $('.msg_bg').html('');
    clearTimeout(msgdsq);
    $('body').append('<div class="sub_err" style="position:absolute;top:60px;left:0;width:500px;z-index:999999;"></div>');
    var errhtml = '<div style="padding:8px 0px;border:1px solid #ff0000;width:100%;margin:0 auto;background-color:#fff;color:#B90802;border:3px #ff0000 solid;text-align:center;font-size:16px;font-family:微软雅黑;"><img style="margin-right:10px;" src="style/images/error.png">';
    var errhtmlfoot = '</div>';
    $('.msg_bg').height($(document).height());
    $('.sub_err').html(errhtml + msg + errhtmlfoot);
    var left = ($(document).width() - 500) / 2;
    $('.sub_err').css({
        'left': left + 'px'
    });
    var scroll_height = $(document).scrollTop();
    $('.sub_err').animate({
        'top': scroll_height + 120
    }, 300);
    msgdsq = setTimeout(function() {
        $('.sub_err').animate({
            'top': scroll_height + 80
        }, 300);
        setTimeout(function() {
            $('.msg_bg').remove();
            $('.sub_err').remove();
        }, 300);
    }, "1000");
}

//正确时：提示调用方法
function show_msg(msg, url) {
    $('.msg_bg').html('');
    clearTimeout(msgdsq);
    $('body').append('<div class="sub_err" style="position:absolute;top:60px;left:0;width:500px;z-index:999999;"></div>');
    var htmltop = '<div style="padding:8px 0px;border:1px solid #090;width:100%;margin:0 auto;background-color:#FFF2F8;color:#090;border:3px #090 solid;;text-align:center;font-size:16px;"><img style="margin-right:10px;" src="style/images/success.png">';
    var htmlfoot = '</div>';
    $('.msg_bg').height($(document).height());
    var left = ($(document).width() - 500) / 2;
    $('.sub_err').css({
        'left': left + 'px'
    });
    $('.sub_err').html(htmltop + msg + htmlfoot);
    var scroll_height = $(document).scrollTop();
    $('.sub_err').animate({
        'top': scroll_height + 120
    }, 500);
    msgdsq = setTimeout(function() {
        $('.sub_err').animate({
            'top': scroll_height + 80
        }, 500);
        setTimeout(function() {
            $('.msg_bg').remove();
            $('.sub_err').remove();
            if (url != '')
            {
                location.href = url;
            }
        }, 800);

    }, "1200");
}

//显示加载动画
function show_loading()
{
    var str = '<div class="msg_bg" style="background:#000;opacity:0.5;filter:alpha(opacity=50);z-index:99998;width:100%;position:absolute;left:0;top:0"></div>';
    str += '<div class="msg_bg" style="z-index:99999;width:100%;position:absolute;left:0;top:0;text-align:center;"><img src="style/images/loading.gif" alt="" class="loading"></div>'
    $('body').append(str);
    var scroll_height = $(document).scrollTop();
    $('.msg_bg').height($(document).height());
    $('.loading').css('margin-top', scroll_height + 240);
}
console.log("\u002f\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u000d\u000a\u0020\u002a\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u002a\u0009\u0009\u000d\u000a\u0020\u002a\u0020\u0009\u0009\u0009\u0009\u0009\u0009\u0020\u0020\u0020\u0020\u0020\u0020\u4ee3\u7801\u5e93\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u002a\u000d\u000a\u0020\u002a\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0077\u0077\u0077\u002e\u0064\u006d\u0061\u006b\u0075\u002e\u0063\u006f\u006d\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u002a\u000d\u000a\u0020\u002a\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0009\u0009\u0020\u0020\u52aa\u529b\u521b\u5efa\u5b8c\u5584\u3001\u6301\u7eed\u66f4\u65b0\u63d2\u4ef6\u4ee5\u53ca\u6a21\u677f\u0009\u0009\u0009\u002a\u000d\u000a\u0020\u002a\u0020\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u0009\u002a\u000d\u000a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002a\u002f");
