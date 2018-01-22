/**
 * 错误提示
 * @param info
 */
function common_error(info){
    sweetAlert("", info, "error");
}

/**
 * 确认提示
 * @param info
 * @param callback
 */
function common_confirm(info,callback){
    swal({
        title: "",
        text: info,
        type: "warning",
        showCancelButton: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        closeOnConfirm: false }, callback);
}

/**
 * 数据提交提示
 */
function common_submit_ing(){
    swal({
        title: "",
        text:"正在提交数据......",
        type:"info",

        showConfirmButton:false
    });
}
/**
 * 用户管理，删除添加的提示框 !自定义
 */
function common_submit_ing_crop(timer){
    swal({
        title: "",
        text:"正在提交数据......",
        type:"info",
        timer:timer,
        showConfirmButton:false
    });
}

/**
 * 打开弹层
 * @param title
 * @param url
 */
function common_open_dialog(title,url,modal) {

    if(url.indexOf("?")>0){
        url += "&time="+ new Date().getTime();
    }else{
        url += "?time="+ new Date().getTime();
    }
    $('.dialog_preview').html('');
    $('.dialog_preview').dialog2({
        showCloseHandle: true,
        removeOnClose: false,
        autoOpen: true,
        modalClass:modal,
        closeOnEscape: false,
        closeOnOverlayClick: false,
        title: "<h4>"+title+"</h4>",
        content: url,
        loaded:function(){
            common_dialog_init();
        }
    });
}

/**
 * 关闭弹层
 * @param reload
 */
function common_close_dialog(reload){
    $(".dialog_preview").parents(".modal").find(".modal-body").dialog2("close");
    if(reload){
        window.location.reload(true);
    }
}

/**
 * 列表页初始化操作
 */
function common_list_init(){
    $("input:checkbox").uniform();

    $("input:radio").uniform();
    $('.datepicker').datepicker({
        "dateFormat":"yy-mm-dd"
    }).on('changeDate', function (ev) {
        $(this).datepicker('hide');
    });
    $('.select2').each(function(){
        var object = $(this).attr("object");
        var filter = $(this).attr("filter");
        if(object && object.length > 0){
            if(filter && object.length > 0){
                common_init_filter($(this),object,$(this).attr("key"),$(this).attr("val"),$(this).attr("v"),$(this).attr("filter"),$(this).attr("filter_key"));
            }else {
                common_init_select($(this),object,$(this).attr("key"),$(this).attr("val"),$(this).attr("v"));
            }

        }
        else{
            $(this).select2({
                placeholder:"请选择"
            });
        }
    });
    var object_view_list = {};
    $(".object_view").each(function(){
        var object = $(this).attr("object");
        if(!object_view_list[object]){
            object_view_list[object] = {
                'key':$(this).attr("key"),
                'field':$(this).attr("field"),
                'val':[]
            };
        }
        object_view_list[object]['val'].push($(this).attr("val"));
    });
    for(var info in object_view_list){
        common_init_object_view(info,object_view_list[info]);
    }
    var form = $('form.search');


    form.delegate("a[action='new']",'click',function () {
        var object_name = $(this).parents(".search").attr('object');
        common_object_new(object_name,$(this),"modalEditAndAdd");
    });
    form.delegate("a[action='show']",'click',function () {
        var object_name = $(this).parents(".search").attr('object');
        common_object_show(object_name,$(this).attr('object-id'),$(this));
    });
    form.delegate("a[action='edit']",'click',function () {

        var object_name = $(this).parents(".search").attr('object');
        common_object_edit(object_name,$(this).attr('object-id'),$(this),"modalEditAndAdd");
    });
    form.delegate("a[action='delete']",'click',function () {
        var object_name = $(this).parents(".search").attr('object');
        common_object_delete(object_name,$(this).attr('object-id'),$(this));
    });
     form.delegate("a[action='add']",'click',function () {
            var object_name = $(this).parents(".search").attr('object');
            common_object_delete(object_name,$(this).attr('object-id'),$(this));
        });

    form.delegate("a[action='publish']",'click',function () {
        var object_name = $(this).parents(".search").attr('object');
        common_object_publish(object_name,$(this).attr('object-id'),$(this).attr('status'));
    });
    form.delegate("a[action='order_up']",'click',function () {
        var object_name = $(this).parents(".search").attr('object');
        common_object_order_up(object_name,$(this).attr('object-id'),$(this).attr('order'));
    });
    form.delegate("a[action='order_down']",'click',function () {
        var object_name = $(this).parents(".search").attr('object');
        common_object_order_down(object_name,$(this).attr('object-id'),$(this).attr('order'));
    });
    form.delegate("a[action='do']","click",function () {
        common_object_do($(this).attr('url'));
    })
    form.delegate("a[action='view']",'click',function () {
        common_open_dialog($(this).attr('title'), $(this).attr('url'));
    });
    <!--停用,启动，状态 的 操作 绑定-->
    form.delegate("a[action='update']",'click',function () {
            var object_name = $(this).parents(".search").attr('object');
            if (object_name=="user"){

                     common_update_status_avatar(object_name,$(this).attr('object-id'),$(this).attr('name'),$(this).attr('username'));
                    }else{

                common_update_status(object_name,$(this).attr('object-id'),$(this).attr('name'),$(this).attr('value'),$(this).attr('title'))
                    }
        });
      <!--管理团队绑定-->
      form.delegate("a[action='ManageCorp']",'click',function () {
                 var object_name = $(this).parents(".search").attr('object');
                 common_manage_corp($(this),object_name,$(this).attr('object-id'),"modalcorp");
         });

      <!--管理用户团队---移除-->
               $('body').on('click', '#userCorp tbody td a[action="ManageRemoveAdd"]',function(){

                    common_delect_add_user_corp($(this),$(this).attr('user-id'),$(this).attr('object-id'));
               })

      <!-- 管理所有团队--添加用户-->
      $('body').on('click', '#AllCorp tbody td a[action="ManageRemoveAdd"]',function(){

               common_delect_add_user_corp($(this),$("#addUserCorp").attr('user-id'),$(this).attr('object-id'));
         })
        <!-- 管理团队-搜索团队-->
      $('body').on('click', '#corpSearch',function(){

      common_manage_corp($(this),"user",$(this).attr('object-id'),"modalcorp",$("#selectCorp").val());
         })
        <!--管理团队角色-->
                      $('body').on('change', '#userCorp tbody td select[id="userRole"]',function(){

                           common_delect_add_user_corp($(this),$(this).parent("td").next().children().attr("user-id"),$(this).parent("td").next().children().attr("object-id"));
                      })
 <!--管理用户屏-->
      form.delegate("a[action='userScreen']",'click',function () {
                 var object_name = $(this).parents(".search").attr('object');
                 common_manage_screen($(this),object_name,$(this).attr('object-id'),"modaluserscreen");
         });
           <!--管理用户大屏---移除-->
               $('body').on('click', '#userScreen tbody td a[action="ManageScreenRemove"]',function(){
            common_delect_userscreen($(this),$("#addUserScreen").attr('user-id'),$(this).attr('object-id'));
               })

            <!--管理团队成员绑定--通用-->
                form.delegate("a[action='BangDing']",'click',function () {
                                var object_name = $(this).parents(".search").attr('object');
                                common_manage_bangding($(this),object_name,$(this).attr('object-id'),"modalcorp");
                        });
              <!--管理团队成员绑定--通用-移除->
                            $('body').on('click', '#leftTable tbody td a[action="RemoveAdd"]',function(){
                                            var object_name = $("#pad-wrapper").attr('object');
                                            var object_id = $("#pad-wrapper").attr('object-id');
                             common_remove_add_user($(this),object_name,object_id,"modalcorp");
                                    });
              <!--管理团队成员绑定--通用-添加->
                          $('body').on('click', '#rightTable tbody td a[action="RemoveAdd"]',function(){
                                                 var object_name = $("#pad-wrapper").attr('object');
                                                 var object_id = $("#pad-wrapper").attr('object-id');
                           common_remove_add_user($(this),object_name,object_id,"modalcorp");
                          });
                   <!--管理团队角色-->
                         $('body').on('change', '#leftTable tbody td select[id="role"]',function(){
                  var object_name = $("#pad-wrapper").attr('object');
                  var object_id = $(this).attr('object-id');
                  var corp_id = $("#pad-wrapper").attr('object-id');
                  common_change_role($(this),object_name,object_id,corp_id,"modalcorp")
                  });
              <!-- 管理团队-搜索团队-->
               $('body').on('click', '#userSearch',function(){
                  var object_name = $("#pad-wrapper").attr('object');
                  var object_id = $("#pad-wrapper").attr('object-id');
                 common_manage_bangding($(this),object_name,object_id,"modalcorp",$("#selectCorp").val());
               })
               <!-- 管理大屏-主页删除通用->
                              $('body').on('click', '#remove',function(){
                                 var object_name = $("#pad-wrapper").attr('object');
                                 var object_id = $("#pad-wrapper").attr('object-id');
                                common_manage_bangding($(this),object_name,object_id,"modalcorp",$("#selectCorp").val());
                              })
            <!--邀请码--添加->
                              form.delegate("a[action='addcode']",'click',function () {
                             var object_name = $(this).parents(".search").attr('object');
                             var  amount=  $("#codeamont").attr("value")

                                 common_object_add(object_name,amount,$(this));
             });
}


/**
 * 弹层初始化
 */
function common_dialog_init(){
    var form = $('.dialog_preview').find('form');
    var operation = form.attr('operation');
    form.validate({
        submitHandler: function(f) {
            switch(operation) {
                case 'add':
                    common_dialog_add(form);
                    break;
                case 'update':
                    common_dialog_update(form);
                    break;
                case 'do_file_upload':
                    common_dialog_file_upload(form);
                    break;
                case 'do_image_upload':
                    common_dialog_image_upload(form);
                    break;
                case 'do_image_desc':
                    common_dialog_image_desc(form);
                    break;
                default:
                    common_dialog_custom(form);
            }
        }
    });
    form.delegate("input[action='cancel']",'click',function () {
        common_close_dialog(false);
    });
    form.find("input:checkbox").uniform();
    form.find("input:radio").uniform();
    form.find(".editor").each(function(){
        var id = $(this).attr("id");
        var editor = new baidu.editor.ui.Editor();
        editor.render(id);
    });
    form.find('.datepicker').datepicker({
        "dateFormat":"yy-mm-dd"
    }).on('changeDate', function (ev) {
        $(this).datepicker('hide');
    });
    form.find('.img-uploader').each(function(){
        var index = $(this).attr('id');
        var view = index + "_view";
        common_init_uploader(index,view);
    });
    form.find('.file-uploader').each(function(){
        var index = $(this).attr('id');
        common_init_uploader(index);
    });
    form.find('.select2').each(function(){
        var object = $(this).attr("object");
        if(object && object.length > 0){
            common_init_select($(this),object,$(this).attr("key"),$(this).attr("val"),$(this).attr("v"));
        }
        else{
            $(this).select2({
                placeholder:"请选择"
            });
        }
    });
    form.delegate("a[action='product_select']",'click',function () {
        var object_id = $(this).attr("object-id");
        var product_id = $(this).attr("product-id");
        product_select(product_id,object_id);
    });
}

/**
 * 编辑页初始化
 */
function common_window_init(){
    var form = $('.container-fluid').find('form');
    var operation = form.attr('operation');
    form.validate({
        submitHandler: function(f) {
            switch(operation) {
                case 'add':
                    common_window_add(form);
                    break;
                case 'update':
                    common_window_update(form);
                    break;
                default:
                    common_window_custom(form);
            }
        }
    });
    form.delegate("input[action='cancel']",'click',function () {
        var back_url = form.attr("back_url");
        if(back_url){
            window.location.href = back_url;
        }else {
            window.location.href = "/" + form.attr("object");
        }
    });
    form.find("input:checkbox").uniform();
    form.find("input:radio").uniform();
    form.find(".editor").each(function(){
        var id = $(this).attr("id");
        var editor = new baidu.editor.ui.Editor();
        editor.render(id);
    });
    form.find('.datepicker').datepicker({
        "dateFormat":"yy-mm-dd"
    }).on('changeDate', function (ev) {
        $(this).datepicker('hide');
    });
    form.find('.img-uploader').click(function(){
        var index = $(this).attr('id');
        var view = index + "_view";
        common_open_dialog('上传图片','/image_uploader?input=' + index + '&view=' + view);
    });
    form.find('.file-uploader').click(function(){
        var index = $(this).attr('file_id');
        common_open_dialog('上传文件','/file_uploader?input=' + index);
    });
    form.find('.thrumb_preview').click(function(){
        if($(this).attr("src") != '/static/img/no-img-gallery.png'){
            common_open_dialog('编辑图片属性','/image_desc?path=' + $(this).attr("src"));
        }
    });
    form.find('.select2').each(function(){
        var object = $(this).attr("object");
        var filter = $(this).attr("filter");
        if(object && object.length > 0){
            if(filter && filter.length > 0){
                common_init_filter($(this),object,$(this).attr("key"),$(this).attr("val"),$(this).attr("v"),$(this).attr("filter"),$(this).attr("filter_key"));
            }else {
                common_init_select($(this),object,$(this).attr("key"),$(this).attr("val"),$(this).attr("v"));
            }
        }
        else{
            $(this).select2({
                placeholder:"请选择"
            });
        }
    });
}

/**
 * 异步获取操作
 * @param url
 * @param success_callback
 * @param confirm
 * @param confirm_info
 */
function common_ajax_get(url, success_callback,confirm,confirm_info) {
    if(confirm && confirm_info){
        common_confirm(confirm_info, function (is_confirm) {
            if(!is_confirm){
                return false;
            }
            $.ajax({
                type: "get",
                dataType: "json",
                url: url,
                beforeSend: function() {
                    common_submit_ing();
                },
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
        });
    }
    else {
        $.ajax({
            type: "get",
            dataType: "json",
            url: url,
            beforeSend: function () {
                common_submit_ing();
            },
            success: function (result) {
                if (result.code == 200) {
                    success_callback(result);
                } else {
                    common_error(result.msg);
                }
            },
            error: function () {
                common_error('接口异常');
            }
        });
    }
}

/**
 * 异步提交操作
 * @param url
 * @param data
 * @param success_callback
 * @param confirm
 * @param confirm_info
 */
function common_ajax_post(url, data, success_callback,confirm,confirm_info) {
    if(confirm && confirm_info){
        common_confirm(confirm_info, function (is_confirm) {
            if(!is_confirm){
                return false;
            }
            $.ajax({
                type: "post",
                dataType: "json",
                data:data,
                url: url,
                beforeSend: function() {
                    common_submit_ing();
                },
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
        });
    }
    else {
        $.ajax({
            type: "post",
            dataType: "json",
            data: data,
            url: url,
            beforeSend: function () {
                common_submit_ing();
            },
            success: function (result) {
                if (result.code == 200) {
                    success_callback(result);
                } else {
                    common_error(result.msg);
                }
            },
            error: function () {
                common_error('接口异常');
            }
        });
    }
}

/**
 * 异步表单提交操作
 * @param form
 * @param api
 * @param success_callback
 */
function common_ajax_form(form, api, success_callback) {
    form.ajaxSubmit({
        url: api,
        type: "post",
        dataType: "json",
        timeout: 10000,
        beforeSubmit:function(){
            common_submit_ing();
        },
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
 * 弹层新建操作
 * @param form
 */
function common_dialog_add(form){
    var object_name = form.attr('object');
    common_ajax_form(form,"/" + object_name + "/add",function(){
        common_close_dialog(true);
    });
}

/**
 * 弹层更新操作
 * @param form
 */
function common_dialog_update(form){
    var object_name = form.attr('object');
    var object_id = form.find("#id").val();
    common_ajax_form(form,"/" + object_name + "/update?id=" + object_id,function(){
        common_close_dialog(true);
    });
}

/**
 * 弹层表单提交操作
 * @param form
 */
function common_dialog_custom(form){
    var object_name = form.attr('object');
    var operation = form.attr('operation');
    common_ajax_form(form,"/" + object_name + "/" + operation,function(){
        common_close_dialog(true);
    });
}

/**
 * 弹层上传文件
 * @param form
 */
function common_dialog_file_upload(form){
    common_ajax_form(form,"/do_file_upload",function(result){
        var bind_input = $("#bind_input").val();
        $("#"+bind_input).val(result.data);
        swal.close();
        common_close_dialog(false);
    });
}

/**
 * 弹层上传图片
 * @param form
 */
function common_dialog_image_upload(form){
    common_ajax_form(form,"/do_image_upload",function(result){
        var bind_input = $("#bind_input").val();
        var bind_view = $("#bind_view").val();
        $("#"+bind_input).val(result.data);
        $("#"+bind_view).attr('src',result.data);
        swal.close();
        common_close_dialog(false);
    });
}

/**
 * 弹层修改图片属性
 * @param form
 */
function common_dialog_image_desc(form){
    common_ajax_form(form,"/do_image_desc",function(result){
        swal.close();
        common_close_dialog(false);
    });
}

/**
 * 页面新建操作
 * @param form
 * @param object_name
 */
function common_window_add(form){
    var object_name = form.attr('object');
    common_ajax_form(form,"/" + object_name + "/add",function(){
        var back_url = form.attr("back_url");
        if(back_url){
            window.location.href = back_url;
        }else {
            window.location.href = "/" + object_name;
        }
    });
}

/**
 * 页面编辑操作
 * @param form
 * @param object_name
 * @param object_id
 */
function common_window_update(form){
    var object_name = form.attr('object');
    var object_id = $("#id").val();
    common_ajax_form(form,"/" + object_name + "/update?id=" + object_id,function(){
        var back_url = form.attr("back_url");
        if(back_url){
            window.location.href = back_url;
        }else {
            window.location.href = "/" + object_name;
        }
    });
}

/**
 * 页面提交操作
 * @param form
 * @param object_name
 * @param action
 */
function common_window_custom(form){
    var object_name = form.attr('object');
    var operation = form.attr('operation');
    common_ajax_form(form,"/" + object_name + "/" + operation,function(){
        var back_url = form.attr("back_url");
        if(back_url){
            window.location.href = back_url;
        }else {
            window.location.href = "/" + object_name;
        }
    });
}

/**
 * 对象删除操作 通用
 * @param object_name
 * @param object_id
 */
function common_object_delete(object_name,object_id,button){
    var name=button.attr("name")
    common_ajax_get('/' + object_name + '/remove?id=' + object_id,function(){
        window.location.reload(true);
    },true,"确定删除"+name+"?");
}

/**
 * 对象添加操作 通用 无弹层
 * @param object_name
 * @param object_id
 */
function common_object_add(object_name,amount,button){
    var name=button.attr("name")
    common_ajax_get('/' + object_name + '/add?id=' + amount,function(){
        window.location.reload(true);
    },true,"确定添加"+amount+"个?");
}

/**
 * 用户膜版启动，停用，审核 操作
 * @param object_name
 * @param object_id
 */
function common_update_status_avatar(object_name,object_id,name,username){
if (name=='Down'){
    common_ajax_get('/' + object_name + '/updateStatusAva?id=' + object_id+'&status=-1',function(){
        window.location.reload(true);
    },true,"确定停用"+username+"吗？停用后无法登录");
    }else if (name=='Up'){
        common_ajax_get('/' + object_name + '/updateStatusAva?id=' + object_id+'&status=1',function(){
            window.location.reload(true);
        },true,"确定启用"+username+"吗？");
        }
        else{
         common_ajax_get('/' + object_name + '/updateStatusAva?id=' + object_id+'&status=1',function(){
                    window.location.reload(true);
                },true,"确定审核"+username+"通过？");
        }
}

/**
 * 通用 启动，停用，审核 操作
 * @param object_name
 * @param object_id
 */
function common_update_status(object_name,object_id,name,value,title){
common_ajax_get('/' + object_name + '/update?id=' + object_id+'&status='+value,function(){
                    window.location.reload(true);
                },true,"确定"+title+name+"?");

}




/**
 * 管理团队
 * @param object_name
 * @param object_id

 */

function common_manage_corp(button,object_name,object_id,modal,corpname){
    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
    var url = '/' + object_name + '/getCorp?id=' + object_id;
    if (corpname!="undefined"){
    url=url+"&corpName="+corpname
    }

    if(query){

        url += "?"+query;
    }
    if(methods == 'window'){

        window.location.href = url;
    }
    else {

        common_open_dialog(title, url,modal);

    }
}
/**
 * 管理团队移除添加,角色，修改
 * @param object_name
 * @param object_id
 */
function common_delect_add_user_corp(button,user_id,object_id){
    var title = button.attr('title');
    var corpid = button.attr('corpid');
    var value=button.attr("value")
    var url=""
    if (title=='移除用户'){
     url = '/user/delectAndAddCorp?id=' + object_id+"&user_id="+user_id+"&title="+1+"&corp_id="+corpid;
    }
    if (title=="添加用户"){
     url = '/user/delectAndAddCorp?id=' + object_id+"&user_id="+user_id+"&title="+2;
        }
      if (title=="改变用户角色"){
             url = '/user/delectAndAddCorp?id=' + object_id+"&user_id="+user_id+"&role="+value;
       }
        var redirecturl='/user/getCorp?id=' + user_id;
    common_ajax_get_corp_screen(url,function(){
    common_open_dialog(title, redirecturl)
},true,"确定"+title+"吗？",user_id,redirecturl);
}


/**
 * 管理团队成员移除添加 修改  通用
 * @param object_name
 * @param object_id
 */
function common_remove_add_user(button,object_name,object_id,modal){
    var title = button.attr('title');
    var headtitle=button.attr('headtitle')
    var parametervalue=button.attr("parametervalue")
    var method=button.attr("method")
    var RedirectMethod=button.attr("redirectMethod")
    var parametername=button.attr("parametername")
    var removed=button.attr("removed")
    var url='/'+object_name+'/'+method +'?id='+object_id+'&'+parametername+'='+parametervalue+'&removed='+removed
    var redirecturl='/'+object_name+'/'+RedirectMethod+'?id=' + object_id;
    common_ajax_get_corp_screen(url,function(){
    common_open_dialog(headtitle, redirecturl)
},true,"确定"+title+"吗？",object_id,redirecturl);
}

/**
 * 管理团队成员角色，修改  通用
 * @param object_name
 * @param object_id
 */
function common_change_role(button,object_name,object_id,corp_id,modal){
    var title = button.attr('title');
    var parametername=button.attr("parametername")
    var parametervalue=button.attr("value")
    var method=button.attr("method")
    var RedirectMethod=button.attr("redirectMethod")
    var url='/'+object_name+'/'+method +'?id='+object_id+'&'+parametername+'='+parametervalue
    var redirecturl='/'+object_name+'/'+RedirectMethod+'?id=' + corp_id;
    common_ajax_get_corp_screen(url,function(){
    common_open_dialog(title, redirecturl)
},true,"确定"+title+"吗？",object_id,redirecturl);
}

/**
 * 管理大屏，移除
 * @param object_name
 * @param object_id
 */
function common_delect_userscreen(button,user_id,object_id){
    var title = button.attr('title');
 var url = '/user/delectUserScreen?id=' + object_id+"&user_id="+user_id;
       var redirecturl='/user/getUserScreen?id=' + user_id;
    common_ajax_get_corp_screen(url,function(){
    common_open_dialog(title, redirecturl)
},true,"确定"+title+"吗？",user_id,redirecturl);
}
/**
 * 管理团队，通用
 * @param object_name
 * @param object_id
 */
function common_manage_bangding(button,object_name,object_id,modal,corpname){
    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
     var method = button.attr('Rmethod');
    var url = '/' + object_name + '/'+method+ '?id=' + object_id;
    if (corpname!="undefined"){
    url=url+"&corpName="+corpname
    }

    if(query){

        url += "?"+query;
    }
    if(methods == 'window'){

        window.location.href = url;
    }
    else {

        common_open_dialog(title, url,modal);

    }
}

/**
 * 管理团队移除添加的自己定义ajax
 * @param object_name
 * @param object_id
 */
function common_ajax_get_corp_screen(url, success_callback,confirm,confirm_info,user_id,redirecturl) {
    if(confirm && confirm_info){
        common_confirm(confirm_info, function (is_confirm) {
            if(!is_confirm){
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
    }
    else {
        $.ajax({
            type: "get",
            dataType: "json",
            url: url,
            beforeSend: function () {
                common_submit_ing();
            },
            success: function (result) {
                if (result.code == 200) {

                } else {
                    common_error(result.msg);
                }
            },
            error: function () {
                common_error('接口异常');
            }
        });
    }
}

/**
 * 管理大屏
 * @param object_name
 * @param object_id

 */

function common_manage_screen(button,object_name,object_id,modal,corpname){
    var title = button.attr('title');
    var url = '/' + object_name + '/getUserScreen?id=' + object_id;
   var query=button.attr("query");
   var methods=button.attr("methods")
    if(query){

        url += "?"+query;
    }
    if(methods == 'window'){

        window.location.href = url;
    }
    else {

        common_open_dialog(title, url,modal);

    }
}




/**
 * 对象上线操作
 * @param object_name
 * @param object_id
 * @param status
 * @returns {boolean}
 */
function common_object_publish(object_name,object_id,status){
    var confirm_info =  status == 1 ? "确定设置上线?":"确定设置下线?";
    common_ajax_get('/' + object_name + '/publish?id=' + object_id + "&status=" + status,function(){
        window.location.reload(true);
    },true,confirm_info);
}

/**
 * 对象上移操作
 * @param object_name
 * @param object_id
 * @param order
 * @returns {boolean}
 */
function common_object_order_up(object_name,object_id,order){
    var order_total = $('.order-table').attr('total');
    var order_list = $('.order-table').attr('order');
    if(order == order_total){
        common_error('已经移动至顶端，不可再移');
        return false;
    }
    var data = {
        'id':object_id,
        'order':order,
        'order_list':order_list,
        'type':0
    };
    common_ajax_post("/"+object_name+"/order",data,function(){
        window.location.reload(true);
    },true,"确认上移么？");
}

/**
 * 对象下移操作
 * @param object_name
 * @param object_id
 * @param order
 * @returns {boolean}
 */
function common_object_order_down(object_name,object_id,order){
    var order_total = $('.order-table').attr('total');
    var order_list = $('.order-table').attr('order');
    if(order == 1){
        common_error('已经移动至底端，不可再移');
        return false;
    }
    var data = {
        'id':object_id,
        'order':order,
        'order_list':order_list,
        'type':1
    };
    common_ajax_post("/"+object_name+"/order",data,function(){
        window.location.reload(true);
    },true,"确认上移么？");
}

/**
 * 对象新建操作
 * @param button
 */
function common_object_new(object_name,button,modal){
    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
var url = '/' + object_name + '/create';
    if(query){
        url += "?"+query;
    }
    if(methods == 'window'){
        window.location.href = url;
    }
    else {
        common_open_dialog(title, url,modal);
    }
}

/**
 * 对象编辑操作
 * @param button
 */
function common_object_edit(object_name,object_id,button,modal){
    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
    var url = '/' + object_name + '/edit?id=' + object_id;
    if(query){
        url += "?"+query;
    }
    if(methods == 'window'){
        window.location.href = url;
    }
    else {
        common_open_dialog(title, url,modal);
    }
}

/**
 * 对象查看操作
 * @param button
 */
function common_object_show(object_name,object_id,button){

    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
    var url = '/' + object_name + '/show?id=' + object_id;
    if(query){
        url += "?"+query;
    }
    if(methods == 'window'){
        window.location.href = url;
    }
    else {
        common_open_dialog(title, url);
    }
}

/**
 * 对象Ajax操作
 * @param button
 */
function common_object_do(url) {
    common_ajax_get(url,function(){
        window.location.reload(true);
    },true,"确定操作么？");
}

/**
 * 获取关联属性
 * @param object_name
 * @param object_view_info
 */
function common_init_object_view(object_name,object_view_info){
    $.ajax({
        url: '/object_view?object_name=' + object_name + "&key=" + object_view_info.val.join(",") + "&field=" + object_view_info.key + "," +object_view_info.field,
        dataType: "json",
        success: function (result) {
            if (result.code == 200) {
                $(".object_view").each(function(){
                    if($(this).attr("object") == object_name){
                        $(this).html(result.data[$(this).attr("val")]);
                    }
                });
            } else {
                common_error(result.msg);
            }
        },
        error:function(){
            common_error('接口异常');
        }
    });
}

/**
 * 初始化关联对象选择列表
 * @param select
 * @param object
 * @param key
 * @param val
 * @param value
 */
function common_init_select(select,object,key,val,value){
    $.ajax({
        url: '/select_init?object_name=' + object + "&fields=" + key + "," + val,
        dataType: "json",
        success: function (result) {
            if (result.code == 200) {
                for(var i = 0; i < result.data.length; i++){
                    if(value && value.length > 0 && result.data[i][key] == value){
                        select.append("<option value='" + result.data[i][key] + "' selected>" + result.data[i][val] + "</option>");
                    }else{
                        select.append("<option value='" + result.data[i][key] + "'>" + result.data[i][val] + "</option>");
                    }
                }
                select.select2({
                    placeholder:"请选择"
                });
            } else {
                common_error(result.msg);
            }
        },
        error:function(){
            common_error('接口异常');
        }
    });
}

function common_init_filter(select,object,key,val,value,filter,filter_key) {
    $("#" + filter).change(function () {
        var filter_value = $(this).val();
        $.ajax({
            url: '/select_filter?object_name=' + object + "&fields=" + key + "," + val + "&filter_key=" + filter_key + "&filter_value=" + filter_value ,
            dataType: "json",
            success: function (result) {
                if (result.code == 200) {
                    select.html("<option value=\"\">请选择</option>");
                    for(var i = 0; i < result.data.length; i++){
                        if(value && value.length > 0 && result.data[i][key] == value){
                            select.append("<option value='" + result.data[i][key] + "' selected>" + result.data[i][val] + "</option>");
                        }else{
                            select.append("<option value='" + result.data[i][key] + "'>" + result.data[i][val] + "</option>");
                        }
                    }
                    select.select2({
                        placeholder:"请选择"
                    });
                } else {
                    common_error(result.msg);
                }
            },
            error:function(){
                common_error('接口异常');
            }
        });
    })
    if(value){
        var filter_value = $("#" + filter).attr("v");
        $.ajax({
            url: '/select_filter?object_name=' + object + "&fields=" + key + "," + val + "&filter_key=" + filter_key + "&filter_value=" + filter_value ,
            dataType: "json",
            success: function (result) {
                if (result.code == 200) {
                    select.html("<option value=\"\">请选择</option>");
                    for(var i = 0; i < result.data.length; i++){
                        if(value && value.length > 0 && result.data[i][key] == value){
                            select.append("<option value='" + result.data[i][key] + "' selected>" + result.data[i][val] + "</option>");
                        }else{
                            select.append("<option value='" + result.data[i][key] + "'>" + result.data[i][val] + "</option>");
                        }
                    }
                    select.select2({
                        placeholder:"请选择"
                    });
                } else {
                    common_error(result.msg);
                }
            },
            error:function(){
                common_error('接口异常');
            }
        });
    }else {
        select.select2({
            placeholder: "请选择"
        });
    }
}

/**
 * 初始化Duploader上传控件
 * @param input
 * @param view
 */
function common_init_uploader(input,view){
    var uploader = new Duploader({
        btn_open: input,
        multiple: false,
        upload_url: "/index/upload",
        upload_type: "post"
    });
    uploader.on('result',function(result){
        if(view) {
            $('#' + view).attr('src', result.file_path);
        }
        $('#' + input).val(result.file_path);
    });
    if(view) {
        uploader.on('file_add', function (file_info) {
            $(".file-list").after("<div class=\"title\">图片alt属性:&nbsp;&nbsp;&nbsp;&nbsp;<input class=\"image_desc\" id=\"image_alt\" name=\"image_alt\" type=\"text\"></div><div class=\"title\">图片title属性:&nbsp;&nbsp;<input class=\"image_desc\" id=\"image_title\" name=\"image_title\" type=\"text\"></div>");
            var file_id = file_info.id;
            var reader = new window.FileReader();
            reader.readAsDataURL(file_info);
            reader.onloadend = function () {
                $('#file_section_' + file_id).append("<img class='uploader_image_preview' src='" + reader.result + "' />");
            };
        });
        uploader.on('file_change', function (file_info) {
            var file_id = file_info.id;
            var reader = new window.FileReader();
            reader.readAsDataURL(file_info);
            reader.onloadend = function () {
                $('.uploader_image_preview').attr('src', reader.result);
            };
        });
        uploader.on('close',function(){
            $(".image_desc").parent().remove();
        });
    }
}

$(document).ready(function(){
    if($(".container-fluid").length > 0){
        common_window_init();
    }
    else if($(".table-hover").length > 0){
        common_list_init();
    }
});
<!-- 全部选中-->

$("#child_delect").change(function() {

            var flag = $("#child_delect").prop("checked");
            if (flag==true) {
                    $(".list_delect").parent("span").attr("class", "checked")
                    $("#listdelect").attr("class", "btn btn-lg btn-primary").attr("disabled",false)
                 $("#listchangetype").attr("class", "btn btn-lg btn-primary").attr("disabled",false)
                }else{
            $(".list_delect").parent("span").attr("class", "")
            $("#listdelect").attr("class", "btn btn-lg ").attr("disabled",true);
            $("#listchangetype").attr("class", "btn btn-lg ").attr("disabled",true);

}
});
<!-- 部分选中-->
$(".list_delect").each(function(){
 var i=0

 $(".list_delect").change(function() {

   var flag = $(this).prop("checked");
   if (flag==true){
    i=i+1
   }else{
   i=i-1
   }
   if (i!=0) {

            $("#child_delect").parent("span").attr("class","checked");
            $("#listdelect").attr("class", "btn btn-lg btn-primary").attr("disabled",false)
           $("#listchangetype").attr("class", "btn btn-lg btn-primary").attr("disabled",false)

           }else{
             $("#child_delect").parent("span").attr("class", "")
   $("#listdelect").attr("class", "btn btn-lg ").attr("disabled",true);
   $("#listchangetype").attr("class", "btn btn-lg ").attr("disabled",true);
                                }});

                   });

 <!-- 批量删除-->
      $('#listdelect').on('click',function(){
      var object_name = $(this).parents(".search").attr('object');
     var title= $(this).parents(".search").attr('name');
      var datas="datas=["
      var data=""
     $(".list_delect").each(function(){
        if (flag = $(this).prop("checked")==true)
        {
            data="{\"object_id\":\""+$(this).parents("td").next().text()+"\""
         +"}";
         if (datas=="datas=["){
         datas=datas+data;
         }else{
         datas=datas+","+data;
         }

        }
         });
         datas=datas+"]"

  common_ajax_post("/"+object_name+"/listremove",datas,function(){
                           window.location.reload(true);
                       },true,"确认(删除/更新)全部"+title+"吗？");


  })
 <!--批量修改绑定-->

      $('#listchangetype').on('click',function() {
                 var object_name = $(this).parents(".search").attr('object');
                 common_manage_changetype($(this),object_name,$(this).attr('object-id'),"changeType");
         });

/**
 * 批量修改
 * @param object_name
 * @param object_id

 */

function common_manage_changetype(button,object_name,object_id,modal){
    var title = button.attr('title');
    var query = button.attr('query');
    var methods = button.attr('methods');
    var url = '/' + object_name + '/changeType';
    if(query){

        url += "?"+query;
    }
    if(methods == 'window'){

        window.location.href = url;
    }
    else {

        common_open_dialog(title, url,modal);

    }
}

/**
 * 批量修改提交
 * @param object_name
 * @param object_id
 */


 $('body').on('click', '#changeType',function(){
       var datas="datas=["
       var data=""
      $(".list_delect").each(function(){
         if (flag = $(this).prop("checked")==true)
         {
             data="{\"object_id\":\""+$(this).parents("td").next().text()+"\""
          +"}";
          if (datas=="datas=["){
          datas=datas+data;
          }else{
          datas=datas+","+data;
          }

         }
          });
          datas=datas+"]"
          if ($('#changType-I').prop("checked")==true){
            datas=datas+"&changType-I=1"
          }else{
          datas=datas+"&changType-I=0"
          }
            if ($('#changType-X').prop("checked")==true){
                      datas=datas+"&changType-X=1"
                 }else{
                    datas=datas+"&changType-X=0"
                    }

   common_ajax_post("/user/listChangeType",datas,function(){
                            window.location.reload(true);
                        },true,"确认修改全部用户吗？");


   })




 var pos = 0;
    var LIST_ITEM_SIZE = 100;
    //滚动条距底部的距离
    var BOTTOM_OFFSET = 0;
    createListItems();
    $(document).ready(function () {
        $("body #dialogchange").scroll(function () {
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
        common_ajax_get_data("/"+"user"+"/getuserdata",function(){


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
            success: function (result) {
                if (result.code == 200) {
                    success_callback(result);
                } else {
                    common_error(result.msg);
                }
            },
            error: function () {
                common_error('接口异常');
            }
        });
    }







