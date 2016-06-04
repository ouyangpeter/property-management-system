{{template "../public/header.tpl"}}
<script type="text/javascript">
    var URL = "/public"

    function modifypassword() {
        $("#dialog").dialog({
            modal: true,
            title: "修改密码",
            width: 300,
            height: 200,
            buttons: [{
                text: '保存',
                iconCls: 'icon-save',
                handler: function () {
                    $("#form1").form('submit', {
                        url: URL + '/changepwd',
                        onSubmit: function () {
                            return $("#form1").form('validate');
                        },
                        success: function (r) {
                            var r = $.parseJSON(r);
                            if (r.status) {
                                $.messager.alert("提示", r.info, 'info', function () {
                                    location.href = URL + "/logout";
                                });
                            } else {
                                vac.alert(r.info);
                            }
                        }
                    });
                }
            }, {
                text: '取消',
                iconCls: 'icon-cancel',
                handler: function () {
                    $("#dialog").dialog("close");
                }
            }]
        });
        $("#oldPassword").focus();
        $("#form1").form('clear');
    }

    $(document).ready(function () {
        addTabIcon("首页", "/public/home", "icon-home");
    });


    $(function () {
        closeTab();
        closePwd();
    })

    $(function () {
        $('#loginOut').click(function () {
            $.messager.confirm('系统提示', '您确定要退出本次登录吗?', function (r) {

                if (r) {
                    location.href = '/public/logout';
                }
            });
        })


    });

    function addTab(subtitle, url) {
        if (!$('#tabs').tabs('exists', subtitle)) {
            $('#tabs').tabs('add', {
                title: subtitle,
                content: createFrame(url),
                closable: true
            });
        } else {
            $('#tabs').tabs('select', subtitle);

            var currTab = $('#tabs').tabs('getSelected');
            var url = $(currTab.panel('options').content).attr('src');
            $('#tabs').tabs('update', {
                tab: currTab,
                options: {
                    content: createFrame(url)
                }
            })
        }
        closeTab();
    }


    function addTabIcon(subtitle, url, icon) {

        if (!$('#tabs').tabs('exists', subtitle)) {
            $('#tabs').tabs('add', {
                title: subtitle,
                content: createFrame(url),
                closable: false,
                icon: icon
            });
        } else {
            $('#tabs').tabs('select', subtitle);

            var currTab = $('#tabs').tabs('getSelected');
            var url = $(currTab.panel('options').content).attr('src');
            $('#tabs').tabs('update', {
                tab: currTab,
                options: {
                    content: createFrame(url)
                }
            })
        }
        closeTab();

    }

    function createFrame(url) {
        var s = '<iframe scrolling="auto" frameborder="0"  src="' + url + '" style="width:100%;height:100%;"></iframe>';
        return s;
    }

    function closeTab() {
        $(".tabs-inner").dblclick(function () {
            var subtitle = $(this).children(".tabs-closable").text();
            $('#tabs').tabs('close', subtitle);
        })
    }


    //关闭登录窗口
    function closePwd() {
        $('#dialog').window('close');
    }

    //修改密码
    function serverLogin() {

        var $oldpass = $('#txtOldPass');

        var $newpass = $('#txtNewPass');
        var $rePass = $('#txtRePass');

        if ($oldpass.val() == '') {
            $.messager.alert('系统提示', '请输入原密码！', 'warning');
            return false;
        }

        if ($newpass.val() == '') {
            $.messager.alert('系统提示', '请输入密码！', 'warning');
            return false;
        }
        if ($rePass.val() == '') {
            $.messager.alert('系统提示', '请在一次输入密码！', 'warning');
            return false;
        }

        if ($newpass.val() != $rePass.val()) {
            $.messager.alert('系统提示', '两次密码不一至！请重新输入', 'warning');
            return false;
        }

        $('#w').window('close');


        var sPost = "";

        sPost += "&oldpass=" + $("#txtOldPass").val();

        sPost += "&password=" + $("#txtNewPass").val();

        var RandTime = new Date();

        $.ajax({

            type: "post",

            async: false,

            cache: false,

            url: "system/ajax_pages/ajax_public.aspx?Action=4&RandTime=" + RandTime.toString() + sPost,

            error: function () {

                $("#txtOldPass").val("");

                $("#txtNewPass").val("");

                $("#txtRePass").val("");

                $.messager.alert('错误', '获取账号信息失败...请联系管理员!', 'error');
            },

            success: function (xml) {
                $(xml).find("item").each(function () {
                    $("#txtOldPass").val("");

                    $("#txtNewPass").val("");

                    $("#txtRePass").val("");


                    var respcode = "";

                    respcode = $(this).find("respcode").text();

                    if (respcode == "1") {
                        $.messager.alert('提示', "密码修改成功,请牢记!", 'info');
                    }
                    else {
                        var respmsg = $(this).find("respmsg").text();

                        $.messager.alert('错误', respmsg, 'error');
                    }
                });
            }

        });
    }
</script>
<body class="easyui-layout" style="overflow-y: hidden;" scroll="no">

<div region="north" border="false" style="overflow: hidden; height: 40px;">

    <div class="heder">

        <div class="top_logo">

            <table cellpadding="0" cellspacing="0" border="0">
                <tr>
                    <td></td>
                    <td valign="top">
                        <div style="padding-top:10px; padding-left:20px;"><span id="Lab_sysname"
                                                                                style="font-size:14px; font-weight:bold; color:White; font-family:宋体; ">物业管理系统</span>
                        </div>
                    </td>
                </tr>
            </table>

        </div>

        <div style="float:right; padding-right:20px; margin-top:10px;" class="head">欢迎 <span id="Lab_user"><font
                color='#ff9933'><strong>{{.userinfo.UserName}}</strong></font></span>&nbsp;&nbsp;|&nbsp;&nbsp;<a
                href="/public/index">返回首页</a>&nbsp;&nbsp;|&nbsp;&nbsp;<a href="#" id="editpass"
                                                                         onclick="modifypassword()">修改密码</a>&nbsp;&nbsp;|&nbsp;&nbsp;<a
                href="#" id="loginOut">安全退出</a>
        </div>

        <div class="clear"></div>
    </div>

</div>

<div region="west" split="true" iconCls="icon-news" title="主菜单" style="width:150px;padding:0px;">
    <div id="d_accordionmenu" class="easyui-accordion">
        <div title='小区管理'>
            <ul>
                <li><a href="#" onclick="addTab('系统公告', '/pms/notice/index');">系统公告</a></li>
            </ul>
        </div>


        <div title='楼盘管理'>
            <ul>
                <li><a href="#" onclick="addTab('楼宇管理', '/pms/building/index')">楼宇管理</a></li>
            </ul>
            <ul>
                <li><a href="#" onclick="addTab('房屋管理', '/pms/house/index')">房屋管理</a></li>
            </ul>
        </div>
        <div title='车位管理'>
            <ul>
                <li><a href="#" onclick="addTab('停车场管理', '/pms/parkingLot/index')">停车场管理</a></li>
            </ul>
            <ul>
                <li><a href="#" onclick="addTab('车位管理', '/pms/parkingSpot/index')">车位管理</a></li>
            </ul>
        </div>
        <div title='收费管理'>
            <ul>
                <li><a href="#" onclick="addTab('收费类型管理', '/pms/chargeType/index')">收费类型管理</a></li>
            </ul>
            <ul>
                <li><a href="#" onclick="addTab('户主收费管理', '/pms/charge/index')">户主收费管理</a></li>
            </ul>
        </div>
        <div title='用户管理'>
            <ul>
                <li><a href="#" onclick="addTab('户主管理', '/pms/owner/index');">户主管理</a></li>
                <li><a href="#" onclick="addTab('系统用户管理', '/pms/user/index');">系统用户管理</a></li>
            </ul>
        </div>
    </div>
</div>
<div region="center" iconCls="icon-index">
    <div class="easyui-tabs" fit="true" border="false" id="tabs">

    </div>
</div>
<div region="south" style="height:40px; background-image: url(/static/easyui/img/footerback.png);">
    <div id="footer" style="padding:10px 5px 5px 5px; color:White;">
        <table width="100%" cellpadding="2" cellspacing="2">
            <tr>
                <td align="right">
                    <div id="d_welinfo" class="index_welinfo"><span style='font-family:Arial;'>&copy;</span>&nbsp;2016&nbsp;毕业设计.
                    </div>
                </td>
            </tr>
        </table>
    </div>
</div>


<!--修改密码窗口-->
<div id="dialog" class="easyui-window" title="修改密码" collapsible="false" minimizable="false" maximizable="false"
     icon="icon-save" style="width: 300px; height: 180px; padding: 5px; background: #fafafa;">
    <div class="easyui-layout" fit="true">
        <div region="center" border="false" style="padding: 10px; background: #fff; border: 0px solid #ccc;">
            <form id="form1" method="post">
                <table width="100%" class="publictable">
                    <tr>
                        <td style="width:90px;">原密码：</td>
                        <td><input id="oldPassword" name="oldPassword" type="Password" class="easyui-validatebox"
                                   required="true"
                                   validType="password[5,20]" missingMessage="请填写当前使用的密码"/></td>
                    </tr>

                    <tr>
                        <td style="width:90px;">新密码：</td>
                        <td><input name="newPassword" type="Password" class="easyui-validatebox" required="true"
                                   validType="password[5,20]" missingMessage="请填写需要修改的密码"/></td>
                    </tr>
                    <tr>
                        <td style="width:90px;">确认新密码：</td>
                        <td><input name="repeatPassword" type="Password" class="easyui-validatebox" required="true"
                                   validType="password[5,20]" missingMessage="请重复填写需要修改的密码"/></td>
                    </tr>
                </table>
            </form>
        </div>
    </div>
</div>

<div id="mm" class="easyui-menu" style="width:150px;">
    <div id="mm-tabupdate">刷新</div>
    <div class="menu-sep"></div>
    <div id="mm-tabclose">关闭</div>

    <div class="menu-sep"></div>
    <div id="mm-exit">退出</div>
</div>


</body>
</html>