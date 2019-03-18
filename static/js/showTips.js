function showTips(content, type) {

    //窗口的宽度
    var windowWidth = $(window).width();
    var tipsDiv = '';
    var backColor = "#43CD80";
    var srcImg = "success.jpg";
    switch (type) {
        case "success":
            backColor = "#EAF8F0";
            srcImg = "success.jpg";
            break;
        case "warning":
            backColor = "#EAF8F0";
            srcImg = "warning.png";
            break;
        case "error":
            backColor = "#FAECE8";
            srcImg = "error.png";
            break;
        case "primary":
            backColor = "#EBF3FD";
            srcImg = "primary.png";
            break;
        default:
            backColor = "#43CD80";
            srcImg = "success.jpg";
    }
    tipsDiv = '<div class="tipsClass"><img src="static/img/'+srcImg+'" width="12px" height="12px"></img>&emsp;' + content + '</div>';
    $('body').append(tipsDiv);
    $('div.tipsClass').css({
        'top': '50px',
        'left': (windowWidth/ 2) - 200 / 2 + 'px',
        'position': 'absolute',
        'padding': '3px 5px',
        'background': backColor,
        'font-size': '12px',
        'margin': '10px auto',
        'text-align': 'center',
        'width': '200px',
        'height': 'auto',
        'color': '#000',
        'border-radius': '4px',
        'opacity': '0.8'
    }).show();
    setTimeout(function () {
        $('div.tipsClass').fadeOut();
    }, (1500));
}