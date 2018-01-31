/**
 * 获取当前宽度和屏幕总宽度，弹层居中
 * @param info
 */
function popup(popupName) {
    var _scrollHeight = $(document).scrollTop(),
    //获取当前窗口距离页面顶部高度
    _windowHeight = $(window).height(),
    //获取当前窗口高度
    _windowWidth = $(window).width(),
    //获取当前窗口宽度
    _popupHeight = popupName.height(),
    //获取弹出层高度
    _popupWeight = popupName.width(); //获取弹出层宽度
    _posiTop = (_windowHeight - _popupHeight) / 2 + _scrollHeight;
    _posiLeft = (_windowWidth - _popupWeight) / 2;
     popupName.css({
        "left": _posiLeft + "px",
        "display": "block"
    }); //设置position
}

var formmysf = $('form.search');
//停用,启动，状态 的 操作 绑定-->
formmysf.delegate("a[action='update']", 'click',function() {
    var object_name = $(this).parents(".search").attr('object');
    if (object_name == "user") {

        common_update_status_avatar(object_name, $(this).attr('object-id'), $(this).attr('name'), $(this).attr('username'));
    } else {

        common_update_status(object_name, $(this).attr('object-id'), $(this).attr('name'), $(this).attr('value'), $(this).attr('title'))
    }
});
//管理团队绑定-->
formmysf.delegate("a[action='ManageCorp']", 'click',function() {
    var object_name = $(this).parents(".search").attr('object');
    common_manage_corp($(this), object_name, $(this).attr('object-id'), "modalcorp");
 center("1000px")
});
/**
  * 用户管理，删除添加的提示框 !自定义
  */
function common_submit_ing_crop(timer) {
    swal({
        title: "",
        text: "正在提交数据......",
        type: "info",
        timer: timer,
        showConfirmButton: false
    });
}

//管理用户团队---移除-->
$('body').on('click', '#userCorp tbody td a[action="ManageRemoveAdd"]',function() {

    common_delect_add_user_corp($(this), $(this).attr('user-id'), $(this).attr('object-id'));
})
// 管理所有团队--添加用户-->
$('body').on('click', '#AllCorp tbody td a[action="ManageRemoveAdd"]',function() {

    common_delect_add_user_corp($(this), $("#addUserCorp").attr('user-id'), $(this).attr('object-id'));
})
//管理团队-搜索团队-->
$('body').on('click', '#corpSearch',function() {

    common_manage_corp($(this), "user", $(this).attr('object-id'), "modalcorp", $("#selectCorp").val());
})
//管理团队角色-->
$('body').on('change', '#userCorp tbody td select[id="userRole"]',function() {

    common_delect_add_user_corp($(this), $(this).parent("td").next().children().attr("user-id"), $(this).parent("td").next().children().attr("object-id"));
})
//管理用户屏-->
formmysf.delegate("a[action='userScreen']", 'click',function() {
    var object_name = $(this).parents(".search").attr('object');
    common_manage_screen($(this), object_name, $(this).attr('object-id'), "modaluserscreen");
     center("520px")
});

//管理用户大屏---移除-->
$('body').on('click', '#userScreen tbody td a[action="ManageScreenRemove"]',function() {
    common_delect_userscreen($(this), $("#addUserScreen").attr('user-id'), $(this).attr('object-id'));
})

//管理团队成员绑定--通用-->
formmysf.delegate("a[action='BangDing']", 'click',function() {
    var object_name = $(this).parents(".search").attr('object');
    common_manage_bangding($(this), object_name, $(this).attr('object-id'), "modalcorp");
    center("1000px")
});
//管理团队成员绑定--通用-移除->
$('body').on('click', '#leftTable tbody td a[action="RemoveAdd"]',function() {
    var object_name = $("#pad-wrapper").attr('object');
    var object_id = $("#pad-wrapper").attr('object-id');
    common_remove_add_user($(this), object_name, object_id, "modalcorp");
});
//管理团队成员绑定--通用-添加->
$('body').on('click', '#rightTable tbody td a[action="RemoveAdd"]',function() {
    var object_name = $("#pad-wrapper").attr('object');
    var object_id = $("#pad-wrapper").attr('object-id');
    common_remove_add_user($(this), object_name, object_id, "modalcorp");
});

//管理团队角色-->
$('body').on('change', '#leftTable tbody td select[id="role"]',function() {
    var object_name = $("#pad-wrapper").attr('object');
    var object_id = $(this).attr('object-id');
    var corp_id = $("#pad-wrapper").attr('object-id');
    common_change_role($(this), object_name, object_id, corp_id, "modalcorp")
});
// 管理团队-搜索团队-->
$('body').on('click', '#userSearch',function() {
    var object_name = $("#pad-wrapper").attr('object');
    var object_id = $("#pad-wrapper").attr('object-id');
    common_manage_bangding($(this), object_name, object_id, "modalcorp", $("#selectCorp").val());
})

// 管理大屏-主页删除通用->
$('body').on('click', '#remove',function() {
    var object_name = $("#pad-wrapper").attr('object');
    var object_id = $("#pad-wrapper").attr('object-id');
    common_manage_bangding($(this), object_name, object_id, "modalcorp", $("#selectCorp").val());
});
//邀请码--添加->
formmysf.delegate("a[action='addcode']", 'click',function() {
    var object_name = $(this).parents(".search").attr('object');
    var amount = $("#codeamont").attr("value")

    common_object_add(object_name, amount, $(this));
});

// 获取表格焦点-->
$('body').on('mouseover', ".showlastone",function() {
    $(this).children("td:last ").find("a").css("display", "")
    console.log($(this).children("td:last").attr("test"))
});
// 失去表格焦点-->
$('body').on('mouseout', ".showlastone",function() {
    $(this).children("td:last ").find("a").css("display", "none")
});
//数据显示--详情显示-->
$('body').on('click', "a[action='BangDingData']",function() {
    var object_name = $(this).parents(".search").attr('object');
    common_manage_bangding($(this), object_name, $(this).attr('object-id'), "modalcorp");

});

/**
 * 用户膜版启动，停用，审核 操作
 * @param object_name
 * @param object_id
 */
function common_update_status_avatar(object_name, object_id, name, username) {
    if (name == 'Down') {
        common_ajax_get('/' + object_name + '/updateStatusAva?id=' + object_id + '&status=-1',
        function() {
            window.location.reload(true);
        },
        true, "确定停用" + username + "吗？停用后无法登录");
    } else if (name == 'Up') {
        common_ajax_get('/' + object_name + '/updateStatusAva?id=' + object_id + '&status=1',
        function() {
            window.location.reload(true);
        },
        true, "确定启用" + username + "吗？");
    } else {
        common_ajax_get('/' + object_name + '/updateStatusAva?id=' + object_id + '&status=1',
        function() {
            window.location.reload(true);
        },
        true, "确定审核" + username + "通过？");
    }
}

/**
 * 通用 启动，停用，审核 操作
 * @param object_name
 * @param object_id
 */
function common_update_status(object_name, object_id, name, value, title) {
    common_ajax_get('/' + object_name + '/update?id=' + object_id + '&status=' + value,
    function() {
        window.location.reload(true);
    },
    true, "确定" + title + name + "?");

}

/**
 * 管理团队
 * @param object_name
 * @param object_id

 */

function common_manage_corp(button, object_name, object_id, modal, corpname) {
    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
    var url = '/' + object_name + '/getCorp?id=' + object_id;
    if (corpname != "undefined") {
        url = url + "&corpName=" + corpname
    }

    if (query) {

        url += "?" + query;
    }
    if (methods == 'window') {

        window.location.href = url;
    } else {

        common_open_dialog(title, url, modal);

    }
}

/**
 * 管理团队移除添加,角色，修改
 * @param object_name
 * @param object_id
 */
function common_delect_add_user_corp(button, user_id, object_id) {
    var title = button.attr('title');
    var corpid = button.attr('corpid');
    var value = button.attr("value");
     var url = ""
    if (title == '移除用户') {
        url = '/user/delectAndAddCorp?id=' + object_id + "&user_id=" + user_id + "&title=" + 1 + "&corp_id=" + corpid;
    }
    if (title == "添加用户") {
        url = '/user/delectAndAddCorp?id=' + object_id + "&user_id=" + user_id + "&title=" + 2;
    }
    if (title == "改变用户角色") {
        url = '/user/delectAndAddCorp?id=' + object_id + "&user_id=" + user_id + "&role=" + value;
    }
    var redirecturl = '/user/getCorp?id=' + user_id;
    common_ajax_get_corp_screen(url,
    function() {
        common_open_dialog(title, redirecturl)
    },
    true, "确定" + title + "吗？", user_id, redirecturl);
}

/**
 * 管理团队成员移除添加 修改  通用
 * @param object_name
 * @param object_id
 */
function common_remove_add_user(button, object_name, object_id, modal) {
    var title = button.attr('title');
    var headtitle = button.attr('headtitle');
     var parametervalue = button.attr("parametervalue");
     var method = button.attr("method") ;
     var RedirectMethod = button.attr("redirectMethod") ;
     var parametername = button.attr("parametername") ;
     var removed = button.attr("removed");
     var url = '/' + object_name + '/' + method + '?id=' + object_id + '&' + parametername + '=' + parametervalue + '&removed=' + removed
    var redirecturl = '/' + object_name + '/' + RedirectMethod + '?id=' + object_id;
    common_ajax_get_corp_screen(url,function() {
        common_open_dialog(headtitle, redirecturl)
    },
    true, "确定" + title + "吗？", object_id, redirecturl);
}

/**
 * 管理团队成员角色，修改  通用
 * @param object_name
 * @param object_id
 */
function common_change_role(button, object_name, object_id, corp_id, modal) {
    var title = button.attr('title');
    var parametername = button.attr("parametername") ;
    var parametervalue = button.attr("value") ;
    var method = button.attr("method") ;
    var RedirectMethod = button.attr("redirectMethod");
     var url = '/' + object_name + '/' + method + '?id=' + object_id + '&' + parametername + '=' + parametervalue
    var redirecturl = '/' + object_name + '/' + RedirectMethod + '?id=' + corp_id;
    common_ajax_get_corp_screen(url,
    function() {
        common_open_dialog(title, redirecturl)
    },
    true, "确定" + title + "吗？", object_id, redirecturl);
}

/**
 * 管理大屏，移除
 * @param object_name
 * @param object_id
 */
function common_delect_userscreen(button, user_id, object_id) {
    var title = button.attr('title');
    var url = '/user/delectUserScreen?id=' + object_id + "&user_id=" + user_id;
    var redirecturl = '/user/getUserScreen?id=' + user_id;
    common_ajax_get_corp_screen(url,
    function() {
        common_open_dialog(title, redirecturl)
    },
    true, "确定" + title + "吗？", user_id, redirecturl);
}

/**
 * 管理团队，通用
 * @param object_name
 * @param object_id
 */
function common_manage_bangding(button, object_name, object_id, modal, corpname) {
    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
    var method = button.attr('Rmethod');
    var url = '/' + object_name + '/' + method + '?id=' + object_id;
    if (corpname != "undefined") {
        url = url + "&corpName=" + corpname
    }

    if (query) {

        url += "?" + query;
    }
    if (methods == 'window') {

        window.location.href = url;
    } else {

        common_open_dialog(title, url, modal);

    }
}

/**
 * 管理团队移除添加的自己定义ajax
 * @param object_name
 * @param object_id
 */
function common_ajax_get_corp_screen(url, success_callback, confirm, confirm_info, user_id, redirecturl) {
    if (confirm && confirm_info) {
        common_confirm(confirm_info,
        function(is_confirm) {
            if (!is_confirm) {
                return false;
            }
            $.ajax({
                type: "get",
                dataType: "json",
                url: url,
                beforeSend: function() {
                    common_submit_ing_crop(500);

                },
                success: function(result) {
                    if (result.code == 200) {
                        if (result.msg != "ok") {
                            common_error(result.msg);
                        }
                        success_callback(result);
                    } else {
                        common_error(result.msg);
                    }
                },
                error: function() {
                    common_error('接口异常');
                }
            });
        });
    } else {
        $.ajax({
            type: "get",
            dataType: "json",
            url: url,
            beforeSend: function() {
                common_submit_ing();
            },
            success: function(result) {
                if (result.code == 200) {

} else {
                    common_error(result.msg);
                }
            },
            error: function() {
                common_error('接口异常');
            }
        });
    }
}

//全部选中
$("#child_delect").change(function() {

    var flag = $("#child_delect").prop("checked");
    if (flag == true) {
        $(".list_delect").parent("span").attr("class", "checked");
        $("#listdelect").attr("class", "btn-flat default").attr("disabled", false);
        $("#listchangetype").attr("class", "btn-flat default").attr("disabled", false);
    } else {
        $(".list_delect").parent("span").attr("class", "");
        $("#listdelect").attr("class", "btn btn-lg ").attr("disabled", true);
        $("#listchangetype").attr("class", "btn btn-lg ").attr("disabled", true);

    }
});
// 部分选中
$(".list_delect").each(function() {
    var i = 0

    $(".list_delect").change(function() {

        var flag = $(this).prop("checked");
        if (flag == true) {
            i = i + 1
        } else {
            i = i - 1
        }
        if (i != 0) {

            $("#child_delect").parent("span").attr("class", "checked");
            $("#listdelect").attr("class", "btn-flat default").attr("disabled", false);
             $("#listchangetype").attr("class", "btn-flat default").attr("disabled", false);

        } else {
            $("#child_delect").parent("span").attr("class", "");
             $("#listdelect").attr("class", "btn btn-lg ").attr("disabled", true);
            $("#listchangetype").attr("class", "btn btn-lg ").attr("disabled", true);
        }
    });

});
//批量删除
$('#listdelect').on('click',
function() {
    var object_name = $(this).parents(".search").attr('object');
    var title = $(this).parents(".search").attr('name');
    var datas = "datas=["
    var data = "";
    $(".list_delect").each(function() {
        if (flag = $(this).prop("checked") == true) {
            data = "{\"object_id\":\"" + $(this).parents("td").next().text() + "\"" + "}";
            if (datas == "datas=[") {
                datas = datas + data;
            } else {
                datas = datas + "," + data;
            }

        }
    });
    datas = datas + "]"

    common_ajax_post("/" + object_name + "/listremove", datas,
    function() {
        window.location.reload(true);
    },
    true, "确认(删除/更新)全部" + title + "吗？");

})
//批量修改绑定
$('#listchangetype').on('click',
function() {
    var object_name = $(this).parents(".search").attr('object');
    common_manage_changetype($(this), object_name, $(this).attr('object-id'), "changeType");
});

/**
 * 批量修改
 * @param object_name
 * @param object_id

 */

function common_manage_changetype(button, object_name, object_id, modal) {
    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
    var url = '/' + object_name + '/changeType';
    if (query) {

        url += "?" + query;
    }
    if (methods == 'window') {

        window.location.href = url;
    } else {

        common_open_dialog(title, url, modal);

    }
}

/**
 * 批量修改提交
 * @param object_name
 * @param object_id
 */

$('body').on('click', '#changeType',
function() {
    var datas = "datas=["
    var data = "";
    $(".list_delect").each(function() {
        if (flag = $(this).prop("checked") == true) {
            data = "{\"object_id\":\"" + $(this).parents("td").next().text() + "\"" + "}";
            if (datas == "datas=[") {
                datas = datas + data;
            } else {
                datas = datas + "," + data;
            }

        }
    });
    datas = datas + "]"
    if ($('#changType-I').prop("checked") == true) {
        datas = datas + "&changType-I=1"
    } else {
        datas = datas + "&changType-I=0"
    }
    if ($('#changType-X').prop("checked") == true) {
        datas = datas + "&changType-X=1"
    } else {
        datas = datas + "&changType-X=0"
    }

    common_ajax_post("/user/listChangeType", datas,
    function() {
        window.location.reload(true);
    },
    true, "确认修改全部用户吗？");

})

var pos = 0;
var LIST_ITEM_SIZE = 100;
//滚动条距底部的距离
var BOTTOM_OFFSET = 0;
createListItems();
$(document).ready(function() {
    $("body #dialogchange").scroll(function() {
        var $currentWindow = $(window);
        //当前窗口的高度
        var windowHeight = $currentWindow.height();
        console.log("current widow height is " + windowHeight);
        //当前滚动条从上往下滚动的距离
        var scrollTop = $currentWindow.scrollTop();
        console.log("current scrollOffset is " + scrollTop);
        //当前文档的高度
        var docHeight = $(document).height();
        console.log("current docHeight is " + docHeight);

        //当 滚动条距底部的距离 + 滚动条滚动的距离 >= 文档的高度 - 窗口的高度
        //换句话说：（滚动条滚动的距离 + 窗口的高度 = 文档的高度）  这个是基本的公式
        if ((BOTTOM_OFFSET + scrollTop) >= docHeight - windowHeight) {
            createListItems();
        }
    });
});

function createListItems() {
    var mydocument = document;
    var mylist = mydocument.getElementById("AllCorp");
    var docFragments = mydocument.createDocumentFragment();
    common_ajax_get_data("/" + "user" + "/getuserdata",
    function() {

        for (var i = pos; i < pos + LIST_ITEM_SIZE; ++i) {
            var liItem = mydocument.createElement("li");
            liItem.innerHTML = "This is item " + i;
            docFragments.appendChild(liItem);
        }

        pos += LIST_ITEM_SIZE;

        /* mylist.appendChild(docFragments);*/
    });
}

function common_ajax_get_data(url, success_callback) {

    $.ajax({
        type: "get",
        dataType: "json",
        url: url,
        success: function(result) {
            if (result.code == 200) {
                success_callback(result);
            } else {
                common_error(result.msg);
            }
        },
        error: function() {
            common_error('接口异常');
        }
    });
}

/**
 * 管理大屏
 * @param object_name
 * @param object_id

 */

function common_manage_screen(button, object_name, object_id, modal, corpname) {
    var title = button.attr('title');
    var url = '/' + object_name + '/getUserScreen?id=' + object_id;
    var query = button.attr("query");
    var methods = button.attr("methods");
     if (query) {

        url += "?" + query;
    }
    if (methods == 'window') {

        window.location.href = url;
    } else {

        common_open_dialog(title, url, modal);

    }
}

function center(width){
$("#dialogchange").css("width", width);
$("#dialogchange").css("margin", "0px");
         popup($("#dialogchange"));

}