{{template "../public/header.tpl"}}

<style type="text/css">
    * {
        margin: 0 auto;
    }

    html {
        color: #000;
        background: #FFF;
        font-size: 12px;
    }

    body, div, dl, dt, dd, ul, ol, li, h1, h2, h3, h4, h5, h6, pre, code, form, fieldset, legend, input, textarea, p, blockquote, th, td {
        margin: 0;
        padding: 0
    }

    .header {
        width: 990px;
        height: 85px;
        margin: 0px auto;
        overflow: hidden;
    }


</style>
<body>
<form name="form1" method="post" action="" id="form1">
    <div class="header">

        <table width="100%">
            <tr style="height:40px; vertical-align:bottom">

                <td style="width:700px;">
                </td>

                <td align="left" valign="middle">

                </td>
            </tr>
        </table>

    </div>

    <div style="height:10px; margin:0px auto; width:800px;"></div>
    <div id="d_info" class="Box_News"
         style="padding:1px; margin:0px auto; width:990px; min-height:200px; height:auto !important; height:200px; overflow:visible; text-align:left;">
        <h2 style='padding:5px; font-size:14px;text-align: center;'>{{.complaint.Title}}</h2>
        <div width='100%' style='padding:5px; height:20px;'>
            <div style='text-align:center; color:blue; font-size:14px;'>投诉人:{{.complaint.Owner.Name}}&nbsp;日期:{{.complaint.Created}}</div>
        </div>
        <div class='home_news'>
            <p>{{.complaint.Content}}</p>

        </div>
    </div>
    <div style="height:10px; margin:0px auto; width:800px;"></div>
    <div class="Box_News"
         style="padding:1px; margin:0px auto; width:990px; min-height:200px; height:auto !important; height:200px; overflow:visible; text-align:left;">
        {{if eq "" .complaint.Description}}
        <div width='100%' style='padding:5px; height:20px;'>
            <div style='text-align:center; color:red; font-size:14px;'>处理中</div>
        </div>
        {{else}}
        <div width='100%' style='padding:5px; height:20px;'>
            <div style='text-align:center; color:blue; font-size:14px;'>处理人:{{.complaint.Manager}}&nbsp;日期:{{.complaint.Modified}}</div>
        </div>
        <div class='home_news'>
            <p>{{.complaint.Description}}</p>
        </div>
        {{end}}
    </div>
    <div style="height:10px; margin:0px auto; width:800px;"></div>


</form>
</body>
</html>

