<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body style="margin:0;padding:0;">
<div style="width:450px;height:400px;margin-left:calc(50% - 225px);margin-top:calc(25% - 200px);">
    <div style="width:100%;height:65px;border-bottom:1px solid #dddddd;margin-top:20px;">
        <img src="https://dev.datahunter.cn/static/images/logo-gray.png" alt="logo" style="height:45px;width:450px;margin-top:3px;">
    </div>
    <div style="font-size:16px;">
        <p style="font-family:'Microsoft YaHei';margin-top:10px;"><span>{{.name}}</span>,您好！</p>
        <p style="font-family:'Microsoft YaHei';margin-top:17px;">点击按钮，接受来自 <span>{{.corp}}</span> 的邀请。</p>
        <div style="height:45px;border-radius:5px;background:#39b86e;color:#ffffff;text-align: center;line-height: 45px;margin-top:10px;cursor: pointer;">
            <a href="{{.url}}" style="display:block;width:100%;height:100%;text-decoration:none;color:#ffffff;font-weight:bold;">点击确认</a></div>
        <p style="color:#727272;font-size:12px;margin-top:10px;">如果按钮无效，请复制以下内容到浏览器中打开。</p>
        <p style="color:#727272;font-size:12px;font-family: Arial;margin-bottom:10px;">{{.url}}</p>
    </div>
    <div style="text-align: center;color:#727272;border-top:1px solid #dddddd;font-size:12px;">
        <p style="margin-top:10px;">如有疑问请联系我们：</p>
        <p style="font-family: Arial;">400-1024-262</p>
        <p style="font-family: Arial;">support@datahunter.cn</p>
    </div>
</div>
</body>
</html>