<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <title>物业管理系统</title>
    <link rel="stylesheet" type="text/css" href="/static/easyui/jquery-easyui/themes/default/easyui.css"/>
    <link rel="stylesheet" type="text/css" href="/static/easyui/jquery-easyui/themes/icon.css"/>
    <link rel="stylesheet" type="text/css" href="/static/easyui/css/default.css"/>
    <link rel="stylesheet" type="text/css" href="/static/easyui/css/k.css"/>
    <link rel="stylesheet" type="text/css" href="/static/easyui/css/home_base.css"/>
    <link rel="stylesheet" type="text/css" href="/static/easyui/css/home_index.css"/>
    <script type="text/javascript" src="/static/easyui/jquery-easyui/jquery.min.js"></script>
    <script type="text/javascript" src="/static/easyui/jquery-easyui/jquery.easyui.min.js"></script>
    <script type="text/javascript" src="/static/easyui/jquery-easyui/common.js"></script>
    <script type="text/javascript" src="/static/easyui/jquery-easyui/easyui_expand.js"></script>
    <script type="text/javascript" src="/static/easyui/jquery-easyui/phpjs-min.js"></script>
    <script type="text/javascript" src="/static/easyui/js/jquery.fixedheadertable.js"></script>
</head>

<body>
<form name="form1" method="post" action="home.aspx" id="form1">
    <div>
        <input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE"
               value="/wEPDwULLTEyOTIxNDQ3ODUPZBYEAgEPZBYCAgEPFgIeB2NvbnRlbnQFWOeJqeS4mueuoeeQhui9r+S7tizniankuJrmlLbotLnova/ku7Ys54mp5Lia6L2v5Lu2LOeJqeS4mueuoeeQhuezu+e7nyzniankuJrmlLbotLnns7vnu59kAgMPZBYCAgEPZBYMAgEPFgIeCWlubmVyaHRtbAV4PGxpPjxzcGFuPuS7iuWkqXwyMDE2LTA1LTI5PC9zcGFuPjxhIG9uY2xpY2s9J2phdmFzY3JpcHQ6dm9pZCgpOycgdGl0bGU9J2RzZHNkc2RzZHMnIHRhcmdldD0nX2JsYW5rJz5kc2RzZHNkc2RzPC9hPjwvbGk+ZAIDDxYCHwEFyQQ8dGFibGUgc3R5bGU9J21hcmdpbjowcHggYXV0bycgd2lkdGg9JzE4MHB4Jz48dHIgc3R5bGU9J2hlaWdodDoyMHB4Oyc+PHRkPueuoeeQhuWkhOaVsDo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjc8L3RkPjwvdHI+PHRyIHN0eWxlPSdoZWlnaHQ6MjBweDsnPjx0ZD7ljaDlnLDpnaLnp686PC90ZD48dGQgYWxpZ249J3JpZ2h0Jz4wLjAwPC90ZD48L3RyPjx0ciBzdHlsZT0naGVpZ2h0OjIwcHg7Jz48dGQ+6L2m5L2N5pWwOjwvdGQ+PHRkIGFsaWduPSdyaWdodCc+NjwvdGQ+PC90cj48dHIgc3R5bGU9J2hlaWdodDoyMHB4Oyc+PHRkPuaIv+mXtOaVsDo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjY0MTwvdGQ+PC90cj48dHIgc3R5bGU9J2hlaWdodDoyMHB4Oyc+PHRkPuS4muS4u+aVsDo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjU3NTwvdGQ+PC90cj48dHIgc3R5bGU9J2hlaWdodDoyMHB4Oyc+PHRkPuenn+aIt+aVsDo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjE2PC90ZD48L3RyPjx0ciBzdHlsZT0naGVpZ2h0OjIwcHg7Jz48dGQ+5YWl5L2P546HKCUpOjwvdGQ+PHRkIGFsaWduPSdyaWdodCc+OTIuMjA8L3RkPjwvdHI+PC90YWJsZT5kAgUPFgIfAQWFBDxsaT48c3Bhbj4yMDE2LTAyLTI1PC9zcGFuPjxhIGhyZWY9J3Nob3duZXdzLmFzcHg/SWQ9MjAxNjAyMjUwMDAwMDE4OScgdGl0bGU9J+aLk+aJkeeJqeS4mueuoeeQhuezu+e7n+eugOS7iycgdGFyZ2V0PSdfYmxhbmsnPuaLk+aJkeeJqeS4mueuoeeQhuezu+e7n+eugOS7izwvYT48L2xpPjxsaT48c3Bhbj4yMDE2LTAyLTI1PC9zcGFuPjxhIGhyZWY9J3Nob3duZXdzLmFzcHg/SWQ9MjAxNjAyMjUwMDAwMDE5MCcgdGl0bGU9J+aLk+aJkeeJqeS4mueuoeeQhuezu+e7n+S9v+eUqOmhu+efpScgdGFyZ2V0PSdfYmxhbmsnPuaLk+aJkeeJqeS4mueuoeeQhuezu+e7n+S9v+eUqOmhu+efpTwvYT48L2xpPjxsaT48c3Bhbj4yMDE2LTAyLTI1PC9zcGFuPjxhIGhyZWY9J3Nob3duZXdzLmFzcHg/SWQ9MjAxNjAyMjUwMDAwMDE5MScgdGl0bGU9J+WFs+S6jklFNi4w5rWP6KeI5Zmo55qE5Y2H57qn6K+05piOJyB0YXJnZXQ9J19ibGFuayc+5YWz5LqOSUU2LjDmtY/op4jlmajnmoTljYfnuqfor7TmmI48L2E+PC9saT5kAgcPFgIfAQX5BTx0YWJsZSBzdHlsZT0nbWFyZ2luOjBweCBhdXRvJyB3aWR0aD0nMTgwcHgnPjx0ciBzdHlsZT0naGVpZ2h0OjIxcHg7Jz48dGQ+5pS26LS55Lq65qyhOjwvdGQ+PHRkIGFsaWduPSdyaWdodCc+MTwvdGQ+PC90cj48dHIgc3R5bGU9J2hlaWdodDoyMXB4Oyc+PHRkPuaUtui0ueeslOaVsDo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjg8L3RkPjwvdHI+PHRyIHN0eWxlPSdoZWlnaHQ6MjFweDsnPjx0ZD7lupTmlLbph5Hpop06PC90ZD48dGQgYWxpZ249J3JpZ2h0Jz4yNjYuNjA8L3RkPjwvdHI+PHRyIHN0eWxlPSdoZWlnaHQ6MjFweDsnPjx0ZD7lh4/lhY3ph5Hpop06PC90ZD48dGQgYWxpZ249J3JpZ2h0Jz4wLjAwPC90ZD48L3RyPjx0ciBzdHlsZT0naGVpZ2h0OjIxcHg7Jz48dGQ+5LyY5oOg6YeR6aKdOjwvdGQ+PHRkIGFsaWduPSdyaWdodCc+MC4wMDwvdGQ+PC90cj48dHIgc3R5bGU9J2hlaWdodDoyMXB4Oyc+PHRkPuWunuaUtumHkeminTo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjxmb250IGNvbG9yPSdyZWQnPjI2Ni42MDwvZm9udD48L3RkPjwvdHI+PHRyIHN0eWxlPSdoZWlnaHQ6MjFweDsnPjx0ZD7mirzph5Hph5Hpop06PC90ZD48dGQgYWxpZ249J3JpZ2h0Jz48Zm9udCBjb2xvcj0ncmVkJz4wLjAwPC9mb250PjwvdGQ+PC90cj48dHIgc3R5bGU9J2hlaWdodDoyMXB4Oyc+PHRkPumihOaUtuasvumHkeminTo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjxmb250IGNvbG9yPSdyZWQnPjAuMDA8L2ZvbnQ+PC90ZD48L3RyPjwvdGFibGU+ZAIJDxYCHwEFoAM8bGk+PHNwYW4+MjAxNi0wNS0yNzwvc3Bhbj48YSBocmVmPSdzaG93c21zLmFzcHg/SWQ9MjAxNjA1MjcxNDAwMDIwOScgdGl0bGU9J2EnIHRhcmdldD0nX2JsYW5rJz5hPC9hPjwvbGk+PGxpPjxzcGFuPjIwMTYtMDItMjU8L3NwYW4+PGEgaHJlZj0nc2hvd3Ntcy5hc3B4P0lkPTIwMTYwMjI1MDAwMDAyMDEnIHRpdGxlPSfmioTooajmlLbotLnop4bpopHmlZnnqIsnIHRhcmdldD0nX2JsYW5rJz7mioTooajmlLbotLnop4bpopHmlZnnqIs8L2E+PC9saT48bGk+PHNwYW4+MjAxNi0wMi0yNTwvc3Bhbj48YSBocmVmPSdzaG93c21zLmFzcHg/SWQ9MjAxNjAyMjUwMDAwMDIwMicgdGl0bGU9J+aJk+WNsOaOp+S7tueahOWuieijheS9v+eUqCcgdGFyZ2V0PSdfYmxhbmsnPuaJk+WNsOaOp+S7tueahOWuieijheS9v+eUqDwvYT48L2xpPmQCCw8WAh8BBdUSPHRhYmxlIHN0eWxlPSdtYXJnaW46MHB4IGF1dG8nIHdpZHRoPScxODBweCc+PHRyIHN0eWxlPSdoZWlnaHQ6MjNweDsnPjx0ZCBzdHlsZT0nd2lkdGg6ODBweDsnPuW6lOaUtuWIsOacnzo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjx0YWJsZSB3aWR0aD0nMTAwcHgnPjx0cj48dGQgc3R5bGU9J3dpZHRoOjI1cHg7JyBhbGlnbj0ncmlnaHQnPuS7iuWkqTwvdGQ+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J2NlbnRlcic+MDwvdGQ+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J3JpZ2h0Jz7mnKzmnIg8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdjZW50ZXInPjxmb250IGNvbG9yPSdyZWQnPjE2MTwvZm9udD48L3RkPjwvdHI+PC90YWJsZT48L3RkPjwvdHI+PHRyIHN0eWxlPSdoZWlnaHQ6MjNweDsnPjx0ZCBzdHlsZT0nd2lkdGg6ODBweDsnPui0ueeUqOWIsOacnzo8L3RkPjx0ZCBhbGlnbj0ncmlnaHQnPjx0YWJsZSB3aWR0aD0nMTAwcHgnPjx0cj48dGQgc3R5bGU9J3dpZHRoOjI1cHg7JyBhbGlnbj0ncmlnaHQnPuS7iuWkqTwvdGQ+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J2NlbnRlcic+MDwvdGQ+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J3JpZ2h0Jz7mnKzmnIg8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdjZW50ZXInPjxmb250IGNvbG9yPSdyZWQnPjI8L2ZvbnQ+PC90ZD48L3RyPjwvdGFibGU+PC90ZD48L3RyPjx0ciBzdHlsZT0naGVpZ2h0OjIzcHg7Jz48dGQgc3R5bGU9J3dpZHRoOjgwcHg7Jz7lrqLmiLflkIjlkIw6PC90ZD48dGQgYWxpZ249J3JpZ2h0Jz48dGFibGUgd2lkdGg9JzEwMHB4Jz48dHI+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J3JpZ2h0Jz7ku4rlpKk8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdjZW50ZXInPjA8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdyaWdodCc+5pys5pyIPC90ZD48dGQgc3R5bGU9J3dpZHRoOjI1cHg7JyBhbGlnbj0nY2VudGVyJz4wPC90ZD48L3RyPjwvdGFibGU+PC90ZD48L3RyPjx0ciBzdHlsZT0naGVpZ2h0OjIzcHg7Jz48dGQgc3R5bGU9J3dpZHRoOjgwcHg7Jz7ovabkvY3lkIjlkIw6PC90ZD48dGQgYWxpZ249J3JpZ2h0Jz48dGFibGUgd2lkdGg9JzEwMHB4Jz48dHI+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J3JpZ2h0Jz7ku4rlpKk8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdjZW50ZXInPjA8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdyaWdodCc+5pys5pyIPC90ZD48dGQgc3R5bGU9J3dpZHRoOjI1cHg7JyBhbGlnbj0nY2VudGVyJz4wPC90ZD48L3RyPjwvdGFibGU+PC90ZD48L3RyPjx0ciBzdHlsZT0naGVpZ2h0OjIzcHg7Jz48dGQgc3R5bGU9J3dpZHRoOjgwcHg7Jz7np5/otYHlkIjlkIw6PC90ZD48dGQgYWxpZ249J3JpZ2h0Jz48dGFibGUgd2lkdGg9JzEwMHB4Jz48dHI+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J3JpZ2h0Jz7ku4rlpKk8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdjZW50ZXInPjA8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdyaWdodCc+5pys5pyIPC90ZD48dGQgc3R5bGU9J3dpZHRoOjI1cHg7JyBhbGlnbj0nY2VudGVyJz4wPC90ZD48L3RyPjwvdGFibGU+PC90ZD48L3RyPjx0ciBzdHlsZT0naGVpZ2h0OjIzcHg7Jz48dGQgc3R5bGU9J3dpZHRoOjgwcHg7Jz7miL/pl7TlrqLmiLfnlJ/ml6U6PC90ZD48dGQgYWxpZ249J3JpZ2h0Jz48dGFibGUgd2lkdGg9JzEwMHB4Jz48dHI+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J3JpZ2h0Jz7ku4rlpKk8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdjZW50ZXInPjA8L3RkPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdyaWdodCc+5pys5pyIPC90ZD48dGQgc3R5bGU9J3dpZHRoOjI1cHg7JyBhbGlnbj0nY2VudGVyJz48Zm9udCBjb2xvcj0ncmVkJz4xPC9mb250PjwvdGQ+PC90cj48L3RhYmxlPjwvdGQ+PC90cj48dHIgc3R5bGU9J2hlaWdodDoyM3B4Oyc+PHRkIHN0eWxlPSd3aWR0aDo4MHB4Oyc+56ef6LWB5a6i5oi355Sf5pelOjwvdGQ+PHRkIGFsaWduPSdyaWdodCc+PHRhYmxlIHdpZHRoPScxMDBweCc+PHRyPjx0ZCBzdHlsZT0nd2lkdGg6MjVweDsnIGFsaWduPSdyaWdodCc+5LuK5aSpPC90ZD48dGQgc3R5bGU9J3dpZHRoOjI1cHg7JyBhbGlnbj0nY2VudGVyJz4wPC90ZD48dGQgc3R5bGU9J3dpZHRoOjI1cHg7JyBhbGlnbj0ncmlnaHQnPuacrOaciDwvdGQ+PHRkIHN0eWxlPSd3aWR0aDoyNXB4OycgYWxpZ249J2NlbnRlcic+MDwvdGQ+PC90cj48L3RhYmxlPjwvdGQ+PC90cj48L3RhYmxlPmRkuyCKU3QFbJglszLOOe1mYlK0Tnk="/>
    </div>

    <script type="text/javascript" language="javascript">var warn_rent_num = 0;
    var warn_zlkh_num = 0;
    var warn_feedq_num = 0;
    var home_swtxdialogstoptime = 0; </script>

    <div id="d_home" style="padding:5px;">

        <div class="index">
            <div class="warp cf mt10">
                <div class="idx_left fl">
                    <div class="mode">
                        <div class="title"><a href="" target="_blank" class="more">更多 &gt;&gt;</a><span>通告栏</span>
                        </div>
                        <div class="announce">
                            <ul id="d_news" style="height:370px;">
                                {{range .notices}}
                                <li><span>{{.Created}}</span><a href='notice_content?Id={{.Id}}'
                                                              title='{{.Title}}'
                                                              target='_blank'>{{.Title}}</a></li>

                                {{end}}
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="idx_right fl">
                    <div class="mode">
                        <div class="title">&gt; 小区信息</div>
                        <div class="questions">
                            <ul id="d_feeinfo" style="height:370px;">
                                <table style='margin:0px auto' width='180px'>
                                    <tr style='height:21px;'>
                                        <td>小区名称:</td>
                                        <td align='right'>{{.communityInfo.CommunityName}}</td>
                                    </tr>
                                    <tr style='height:21px;'>
                                        <td>主要负责人:</td>
                                        <td align='right'>{{.communityInfo.PrincipalName}}</td>
                                    </tr>
                                    <tr style='height:21px;'>
                                        <td>联系电话:</td>
                                        <td align='right'>{{.communityInfo.PrincipalPhone}}</td>
                                    </tr>
                                    <tr style='height:21px;'>
                                        <td>建造日期:</td>
                                        <td align='right'>{{.communityInfo.BuildDate}}</td>
                                    </tr>
                                    <tr style='height:21px;'>
                                        <td>停车场面积(m<SUP>2</SUP>):</td>
                                        <td align='right'>{{.communityInfo.ParkingLotArea}}</td>
                                    </tr>
                                    <tr style='height:21px;'>
                                        <td>建筑面积(m<SUP>2</SUP>):</td>
                                        <td align='right'>{{.communityInfo.BuildingArea}}</td>
                                    </tr>
                                    <tr style='height:21px;'>
                                        <td>绿化面积(m<SUP>2</SUP>):</td>
                                        <td align='right'>{{.communityInfo.GreeningArea}}</td>
                                    </tr>
                                    <tr style='height:21px;'>
                                        <td>道路面积(m<SUP>2</SUP>):</td>
                                        <td align='right'>{{.communityInfo.RoadArea}}</td>
                                    </tr>
                                </table>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </div>

    </div>

</form>
</body>

</body>
</html>
