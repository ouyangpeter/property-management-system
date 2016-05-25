{{template "../public/header.tpl"}}
<script type="text/javascript">
    var URL = "/pms/owner";
    $(function () {
        //户主列表
        $("#datagrid").datagrid({
            title: '户主列表',
            url: URL + '/index',
            method: 'POST',
            pagination: true,
            fitColumns: true,
            striped: true,
            rownumbers: true,
            singleSelect: true,
            idField: 'Id',
            pagination: true,
            fit: true,
            pageSize: 20,
            pageList: [10, 20, 30, 50, 100],
            columns: [[
                {field: 'Id', title: 'ID', width: 30, sortable: true},
                {field: 'Name', title: '姓名', width: 30, align: 'center', sortable: true},
                {field: 'PhoneNumber', title: '手机号', width: 30, align: 'center', editor: 'text', sortable: true},
                {field: 'IdCard', title: '身份证', width: 30, align: 'center', editor: 'text', sortable: true},
                {field: 'Company', title: '工作单位', width: 30, align: 'center', editor: 'text', sortable: true},
                {
                    field: 'UserName', title: '用户名', width: 30, align: 'center', sortable: false,
                    formatter: function (value, rec) {
                        if (rec != null && rec.User != null) {
                            return rec.User.UserName;
                        } else {
                            return "";
                        }
                    }
                },
                {
                    field: 'Email', title: '邮箱', width: 50, align: 'center', editor: 'text', sortable: false,
                    formatter: function (value, rec) {
                        if (rec != null && rec.User != null) {
                            return rec.User.Email;
                        } else {
                            return "";
                        }
                    }
                },
                {field: 'Remark', title: '备注', width: 50, align: 'center', editor: 'text'}
            ]],
            onAfterEdit: function (index, data, changes) {
                if (vac.isEmpty(changes)) {
                    return;
                }
                changes.Id = data.Id;
                vac.ajax(URL + '/updateHouse', changes, 'POST', function (r) {
                    if (!r.status) {
                        vac.alert(r.info);
                    } else {
                        $("#datagrid").datagrid("reload");
                    }
                })
            },
            onDblClickRow: function (index, row) {
                editrow();
            },
            onRowContextMenu: function (e, index, row) {
                e.preventDefault();
                $(this).datagrid("selectRow", index);
                $('#mm').menu('show', {
                    left: e.clientX,
                    top: e.clientY
                });
            },
            onHeaderContextMenu: function (e, field) {
                e.preventDefault();
                $('#mm1').menu('show', {
                    left: e.clientX,
                    top: e.clientY
                });
            },
            onLoadSuccess: function (data) {
//                console.log(data.total)
                if (data.total == 0) {
                    //添加一个新数据行，第一列的值为你需要的提示信息，然后将其他列合并到第一列来，注意修改colspan参数为你columns配置的总列数
                    $(this).datagrid('appendRow', {Id: '<div style="text-align:center;color:red">没有相关记录！</div>'}).datagrid('mergeCells', {
                        index: 0,
                        field: 'Id',
                        colspan: 8
                    });
                    //隐藏分页导航条，这个需要熟悉datagrid的html结构，直接用jquery操作DOM对象，easyui datagrid没有提供相关方法隐藏导航条
                    $(this).closest('div.datagrid-wrap').find('div.datagrid-pager').hide();
                }
                //如果通过调用reload方法重新加载数据有数据时显示出分页导航容器
                else $(this).closest('div.datagrid-wrap').find('div.datagrid-pager').show();
            }
        });
        //创建添加户主窗口
        $("#dialog").dialog({
            modal: true,
            resizable: true,
            top: 150,
            closed: true,
            buttons: [{
                text: '保存',
                iconCls: 'icon-save',
                handler: function () {
                    $("#form1").form('submit', {
                        url: URL + '/addOwner',
                        onSubmit: function () {
                            return $("#form1").form('validate');
                        },
                        success: function (r) {
                            var r = $.parseJSON(r);
                            if (r.status) {
                                $("#dialog").dialog("close");
                                $("#datagrid").datagrid('reload');
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
    })

    function editrow() {
        if (!$("#datagrid").datagrid("getSelected")) {
            vac.alert("请选择要编辑的行");
            return;
        }
        $('#datagrid').datagrid('beginEdit', vac.getindex("datagrid"));
    }
    function saverow(index) {
        if (!$("#datagrid").datagrid("getSelected")) {
            vac.alert("请选择要保存的行");
            return;
        }
        $('#datagrid').datagrid('endEdit', vac.getindex("datagrid"));
    }
    //取消
    function cancelrow() {
        if (!$("#datagrid").datagrid("getSelected")) {
            vac.alert("请选择要取消的行");
            return;
        }
        $("#datagrid").datagrid("cancelEdit", vac.getindex("datagrid"));
    }
    //刷新
    function reloadrow() {
        $("#datagrid").datagrid("reload");
    }

    //添加房屋弹窗
    function addrow() {
        $("#dialog").dialog('open');
        $("#form1").form('clear');
    }


    //删除
    function delrow() {
        $.messager.confirm('Confirm', '你确定要删除?', function (r) {
            if (r) {
                var row = $("#datagrid").datagrid("getSelected");
                if (!row) {
                    vac.alert("请选择要删除的行");
                    return;
                }
                vac.ajax(URL + '/deleteOwner', {Id: row.Id}, 'POST', function (r) {
                    if (r.status) {
                        $("#datagrid").datagrid('reload');
                    } else {
                        vac.alert(r.info);
                    }
                })
            }
        });
    }

    function Query() {
        var postData = new Object();

        if ($('#query_building_id').val() != '0') {
            postData.building_id = $('#query_building_id').val();
        }

        if ($('#query_unit_name').val() != '') {
            postData.unit_name = $('#query_unit_name').val();
        }

        if ($('#query_area').val() != '') {
            postData.area = $('#query_area').val();
        }

        if ($('#query_house_no').val() != '') {
            postData.house_no = $('#query_house_no').val();
        }
        if ($('#query_owner_id').val() != '') {
            postData.owner_id = $('#query_owner_id').val();
        }
        $('#datagrid').datagrid('load', postData);
    }
</script>
<body style="padding:2px; margin:0px;" class="panel-noscroll">
<div class="easyui-layout layout" fit="true">
    <div data-options="region:'north'" style="margin: 2px; height: 65px;">
        <div class="panel-header">
            <div class="panel-title">查询条件</div>
            <div class="panel-tool"><a href="javascript:void(0)" class="layout-button-up"></a></div>
        </div>
        <div data-options="region:'north',split:true,title:'查询条件'" style="padding: 2px;"
             title="" class="panel-body layout-body">
            <table cellpadding="3" cellspacing="3">
                <tbody>
                <tr>
                    <td>楼宇名称:</td>
                    <td><input class="easyui-combobox"
                               id="query_building_id"
                               data-options="
    valueField: 'Id',
    textField: 'Name',
    url: '/pms/building/buildingList'
    "></td>
                    <td>单元名称:</td>
                    <td><input class="easyui-combobox" id="query_unit_name"></td>
                    <td>房号:</td>
                    <td><input type="text" id="query_house_no" value="" size="10"></td>
                    <td>面积:</td>
                    <td><input type="text" id="query_area" value="" size="10"></td>
                    <td>户主:</td>
                    <td><input type="text" id="query_owner_id" value="" size="10"
                               onkeydown="if(event.keyCode == 13) {Query();}"></td>

                    <td><input value="查询" type="button" id="BN_Find" style="width:80px; height:30px;" class="button"
                               onclick="Query();"></td>
                </tr>
                </tbody>
            </table>

        </div>
    </div>
    <div data-options="region:'center'" style="margin: 2px;">
        <table id="datagrid" toolbar="#tb"></table>
    </div>
</div>
<div id="tb" style="padding:5px;height:auto">
    <a href="#" icon='icon-add' plain="true" onclick="addrow()" class="easyui-linkbutton">新增</a>
    <a href="#" icon='icon-edit' plain="true" onclick="editrow()" class="easyui-linkbutton">编辑</a>
    <a href="#" icon='icon-save' plain="true" onclick="saverow()" class="easyui-linkbutton">保存</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="delrow()" class="easyui-linkbutton">删除</a>
    <a href="#" icon='icon-reload' plain="true" onclick="reloadrow()" class="easyui-linkbutton">刷新</a>
</div>
<!--表格内的右键菜单-->
<div id="mm" class="easyui-menu" style="width:120px;display: none">
    <div iconCls='icon-add' onclick="addrow()">新增</div>
    <div iconCls="icon-edit" onclick="editrow()">编辑</div>
    <div iconCls='icon-save' onclick="saverow()">保存</div>
    <div iconCls='icon-cancel' onclick="cancelrow()">取消</div>
    <div class="menu-sep"></div>
    <div iconCls='icon-cancel' onclick="delrow()">删除</div>
    <div iconCls='icon-reload' onclick="reloadrow()">刷新</div>
    <div class="menu-sep"></div>
    <div>Exit</div>
</div>
<!--表头的右键菜单-->
<div id="mm1" class="easyui-menu" style="width:120px;display: none">
    <div icon='icon-add' onclick="addrow()">新增</div>
</div>
<div id="dialog" title="添加户主" style="width:400px;height:400px;">
    <div style="padding:20px 20px 40px 80px;">
        <form id="form1" method="post">
            <table>
                <tr>
                    <td>楼宇：</td>
                    <td><input class="easyui-combobox"
                               name="BuildingId"
                               data-options="
    valueField: 'Id',
    textField: 'Name',
    url: '/pms/building/buildingList',
    onSelect: function(rec){
    $('#UnitName').combobox('clear');
    $('#HouseId').combobox('clear');
    var url = '/pms/building/unitList?BuildingId='+rec.Id;
    $('#UnitName').combobox('reload', url);
    }
    " required="true"></td>
                </tr>
                <tr>
                    <td>单元名称：</td>
                    <td><select class="easyui-combobox" id="UnitName" style="width:120px;" required="true"
                                data-options="valueField:'UnitName',textField:'UnitName',onSelect: function(rec){
    $('#HouseId').combobox('clear');
    var url = '/pms/house/houseList?building_id='+rec.Building.Id+'&unit_name='+rec.UnitName;
    $('#HouseId').combobox('reload', url);
    }">
                    </select></td>
                </tr>
                <tr>
                    <td>房号：</td>
                    <td><select class="easyui-combobox" id="HouseId" name="HouseId" style="width:120px;"
                                required="true" data-options="valueField:'Id',textField:'HouseNo', method:'post'">
                    </select></td>
                </tr>
                <tr>
                    <td>姓名：</td>
                    <td><input name="Name" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>手机号：</td>
                    <td><input name="PhoneNumber" class="easyui-validatebox"/></td>
                </tr>
                <tr>
                    <td>身份证：</td>
                    <td><input name="IdCard" class="easyui-validatebox"/></td>
                </tr>
                <tr>
                    <td>工作单位：</td>
                    <td><input name="Company" class="easyui-validatebox"/></td>
                </tr>
                <tr>
                    <td>用户名：</td>
                    <td><input name="UserName" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>备注：</td>
                    <td><textarea name="Remark" class="easyui-validatebox" validType="length[0,200]"></textarea></td>
                </tr>
            </table>
        </form>
    </div>
</div>
</body>
<script type="text/javascript">
    function myformatter(date) {
        var y = date.getFullYear();
        var m = date.getMonth() + 1;
        var d = date.getDate();
        return y + '-' + (m < 10 ? ('0' + m) : m) + '-' + (d < 10 ? ('0' + d) : d) + "T00:00:00Z";
    }
    function myparser(s) {
        if (!s) return new Date();
        var ss = (s.split('-'));
        var y = parseInt(ss[0], 10);
        var m = parseInt(ss[1], 10);
        var tmp = ss[2].split('T')

        var d = parseInt(tmp[0], 10);
        if (!isNaN(y) && !isNaN(m) && !isNaN(d)) {
            return new Date(y, m - 1, d);
        } else {
            return new Date();
        }
    }

</script>
</html>