{{template "../public/header.tpl"}}
<script type="text/javascript">
    var statuslist = [
        {statusid: '1', name: '禁用'},
        {statusid: '2', name: '启用'}
    ];
    var URL = "/pms/building";
    $(function () {
        //楼宇列表
        $("#datagrid").datagrid({
            title: '楼宇列表',
            url: URL + '/index',
            method: 'POST',
            pagination: true,
            fitColumns: true,
            striped: true,
            rownumbers: true,
            singleSelect: true,
            idField: 'Id',
            pagination: true,
            pageSize: 20,
            pageList: [10, 20, 30, 50, 100],
            columns: [[
                {field: 'Id', title: 'ID', width: 30, sortable: true},
                {field: 'Name', title: '楼宇名称', width: 40, sortable: true},
                {field: 'Floors', title: '层数', width: 30, align: 'right', editor: 'numberbox', sortable: true},
                {field: 'Height', title: '高度', width: 30, align: 'right', editor: 'numberbox', sortable: true},
                {field: 'Area', title: '面积', width: 30, align: 'right', editor: 'numberbox', sortable: true},
                {
                    field: 'BuildDate', title: '建成日期', width: 60, align: 'center',
                    editor: {type: 'datebox', options: {formatter: myformatter, parser: myparser}},
                    formatter: function (value, row, index) {
                        if (value) return phpjs.date("Y-m-d", phpjs.strtotime(value));
                        return value;
                    },
                    sortable: true
                },
                {
                    field: 'Created', title: '添加时间', width: 80, align: 'center',
                    formatter: function (value, row, index) {
                        if (value) return phpjs.date("Y-m-d H:i:s", phpjs.strtotime(value));
                        return value;
                    },
                    sortable: true
                },
                {field: 'Remark', title: '备注', width: 50, align: 'center', editor: 'text'},

            ]],
            onAfterEdit: function (index, data, changes) {
                if (vac.isEmpty(changes)) {
                    return;
                }
                changes.Id = data.Id;
                vac.ajax(URL + '/updateBuilding', changes, 'POST', function (r) {
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
                if (data.rows.length == 0) {
                    var body = $(this).data().datagrid.dc.body2;
                    body.find('table tbody').append('<tr><td width="' + body.width() + '" style="height: 25px; text-align: center;" colspan="6">没有数据</td></tr>');
                }
            }
        });
        //创建添加楼宇窗口
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
                        url: URL + '/addBuilding',
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
        //创建修改密码窗口
        $("#dialog2").dialog({
            modal: true,
            resizable: true,
            top: 150,
            closed: true,
            buttons: [{
                text: '保存',
                iconCls: 'icon-save',
                handler: function () {
                    var selectedRow = $("#datagrid").datagrid('getSelected');
                    var password = $('#password').val();
                    vac.ajax(URL + '/UpdateUser', {Id: selectedRow.Id, Password: password}, 'post', function (r) {
                        if (r.status) {
                            $("#dialog2").dialog("close");
                        } else {
                            vac.alert(r.info);
                        }
                    })
                }
            }, {
                text: '取消',
                iconCls: 'icon-cancel',
                handler: function () {
                    $("#dialog2").dialog("close");
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

    //添加楼宇弹窗
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
                vac.ajax(URL + '/deleteBuilding', {Id: row.Id}, 'POST', function (r) {
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
        if ($('#query_name').val() != '') {
            postData.name = $('#query_name').val();
        }

        if ($('#query_floors').val() != '') {
            postData.floors = $('#query_floors').val();
        }
        if ($('#query_area').val() != '') {
            postData.area = $('#query_area').val();
        }
        if ($('#query_height').val() != '') {
            postData.area = $('#query_height').val();
        }
        $('#datagrid').datagrid('load', postData);
    }
</script>
<body style="padding:2px; margin:0px;" class="panel-noscroll easyui-layout">
<div>
    <div data-options="region:'north'" style="margin: 2px; height: 58px;">
        <div class="panel-header">
            <div class="panel-title">查询条件</div>
            <div class="panel-tool"><a href="javascript:void(0)" class="layout-button-up"></a></div>
        </div>
        <div data-options="region:'north',split:true,title:'查询条件'" style="padding: 2px;"
             title="" class="panel-body layout-body">
            <table cellpadding="3" cellspacing="3">
                <tbody>
                <tr>
                    <td>名称:</td>
                    <td><input type="text" id="query_name" value="" size="20"></td>
                    <td>层数:</td>
                    <td><input type="text" id="query_floors" value="" size="10"></td>
                    <td>高度:</td>
                    <td><input type="text" id="query_height" value="" size="10"></td>
                    <td>面积:</td>
                    <td><input type="text" id="query_area" value="" size="10"
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
<div id="dialog" title="添加楼宇" style="width:400px;height:400px;">
    <div style="padding:20px 20px 40px 80px;">
        <form id="form1" method="post">
            <table>
                <tr>
                    <td>名称：</td>
                    <td><input name="Name" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>层数：</td>
                    <td><input name="Floors" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>高度：</td>
                    <td><input name="Height" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>面积：</td>
                    <td><input name="Area" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>建筑日期：</td>
                    <td><input name="BuildDate" class="easyui-datebox" required="true"
                               data-options="formatter:myformatter,parser:myparser"/></td>
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