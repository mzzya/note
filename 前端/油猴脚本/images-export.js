// ==UserScript==
// @name         图片导出
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  支持tmall图片导出
// @author       wangyang
// @match        http*://detail.tmall.com/*
// @match        http*://chaoshi.detail.tmall.com/*
// @match        http*://detail.tmall.hk/*
// @grant        GM_cookie
// @grant        GM_setClipboard
// @grant        GM_setValue
// @grant        GM_getValue
// @grant        GM_xmlhttpRequest
// @grant        GM_registerMenuCommand
// @grant        GM_openInTab
// @grant        GM_download
// @run-at       document-end
// @updateURL    http://localhost:5500/images-export.js
// @downloadURL  http://localhost:5500/images-export.js
// @require      https://lf26-cdn-tos.bytecdntp.com/cdn/expire-1-M/jquery/3.6.0/jquery.min.js
// @require      http://localhost:5500/helper.js
// @require      https://lf9-cdn-tos.bytecdntp.com/cdn/expire-1-M/jszip/3.7.1/jszip.min.js
// @require      https://lf9-cdn-tos.bytecdntp.com/cdn/expire-1-M/FileSaver.js/2014-08-29/FileSaver.js
// @connect      itemcdn.tmall.com
// ==/UserScript==


//js压缩下载文件 https://zhuanlan.zhihu.com/p/352198022
//jszip github https://github.com/Stuk/jszip#readme
//jscdn https://cdn.bytedance.com/

// import JSZip from "jszip";
// import { saveAs } from 'file-saver';

(function () {
    'use strict';

    let clpStartDate = new Date().valueOf();

    function getQueryString(name) {
        var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
        var r = window.location.search.substr(1).match(reg);
        if (r != null) {
            return decodeURIComponent(r[2]);
        }
        return null;
    }


    let container = document.createElement('div');
    container.id = 'clpContainer'
    container.style = 'position: fixed;z-index: 99999999;right: 6px; top:0px;'
    container.innerHTML = `
    <style>
        .clpShade{
            width:360px;
            height:100%;
            background-color: black;
            pointer-events: none;
            opacity: 0.2;
            position:fixed;
            top:0px;
            right:0px;
        }
        .clpHeader{
            position:absolute;
            right:0px;
            display:flex;
            width: 345px;
            justify-content: space-between;
            height:32px;
            top:50px;
        }
        .clpBody{
            height:580px;
        }
        .clpBtn {
            padding: 6px 8px;
            border: 1px solid transparent;
            border-radius: 8px;
            background-color: #0096ad;
            color: #ffffff;
            cursor: pointer;
        }

        #clpTip {
            position: fixed;
            right: 6px;
            top: 90px;
            width: 300px;
            height: 580px;
            overflow: scroll;
            font-size: 12px;
            border: 1px solid #a7dce5;
            background-color: #daf9fe;
            padding: 4px;
            border-radius: 4px;
        }

        #clpBatchDiv {
            position: fixed;
            right:10px;
            top:50px;
            display: none;
            flex-direction: column;
            width: 580px;
            align-items: center;
            margin-top: 40px;
            margin-right: 310px;
            background-color: #daf9fe;
            padding: 10px;
            border-radius: 4px;
        }

        #clpBatchContent {
            border-radius: 4px;
            width: 100%;
            height: 527px;
        }

        #clpBatchStart {
            margin-top: 10px;
        }
    </style>
    <div class="clpShade"></div>
    <div class="clpHeader">
    <button class="clpBtn" id="clpBatchDownload">批量下载</button>
    <input id="clpCode" placeholder="请输入产品编码" />
    <button class="clpBtn" id="clpDownload">下载当前页图片</button>
</div>
<div class="clpBody">
    <div id="clpTip">
    </div>
    <div id="clpBatchDiv">
        <textarea id="clpBatchContent" style="" placeholder="请按照如下格式录入:
产品编码1 抓取连接1
产品编码2 抓取连接2"></textarea>
        <button class="clpBtn" id="clpBatchStart">开始抓取</button>
    </div>
</div>`
    $("body").on("click", "#clpBatchDownload", function () {
        if (!checkLogin()) {
            return;
        }
        var display = $("#clpBatchDiv").css("display");
        $("#clpBatchDiv").css("display", display == "flex" ? "none" : "flex")
    })
    $("body").on("click", "#clpBatchStart", function () {
        clpStartDate = new Date().valueOf();
        var content = $("#clpBatchContent").val();
        if (!content) {
            appendTip(`请输入要抓取的商品编码和链接`, "red");
            return;
        }
        var list = content.split("\n")
        if (!list) {
            appendTip(`输入的内容不正确`, "red");
            return
        }
        var productList = []
        for (let i = 0; i < list.length; i++) {
            const line = list[i].trim();
            if (!line) {
                appendTip(`输入的第${i + 1}行为空，将跳过`, "red");
                continue;
            }
            var lineInfo = line.split(' ')
            if (lineInfo.length != 2) {
                appendTip(`输入的第${i + 1}行未能同时解析到编码和下载链接`, "red");
                continue;
            }
            var code = lineInfo[0];
            var url = lineInfo[1];
            if (!url) {
                appendTip(`输入的第${i + 1}行下载链接不正确`, "red");
                continue;
            }
            url = url.replace(/\#.*/, "")
            productList.push({ code: code, url: url })
        }
        appendTip(`输入的内容初步检测有效数为【${productList.length}/${list.length}】`);
        appendTip(`抓取时将自动打开新窗口，请勿手动关闭右侧窗口，抓取完毕后会自动关闭`, "red")
        for (let i = 0; i < productList.length; i++) {
            const product = productList[i];
            setTimeout(() => {
                appendTip(`开始抓取第${i + 1}/${productList.length}个 ${product.code}`);
                var openUrl = `${product.url}${product.url.indexOf("?") < 0 ? "?" : "&"}clpProductCode=${encodeURI(product.code)}`
                console.log("tools", "openUrl", openUrl)
                var newTap = GM_openInTab(openUrl, { active: true, setParent: true });

            }, i * 10000);
        }
    })
    $("body").on("click", "#clpDownload", function () {
        if (!checkLogin()) {
            return;
        }
        var clpCode = $("#clpCode").val();
        if (!clpCode) {
            productCode = "抓取商品图片信息"
            appendTip(`没有输入产品编码，将使用【${productCode}】作为压缩包名称`, "red");
        } else {
            productCode = clpCode;
        }
        clpStartDate = new Date().valueOf();
        appendTip("开始下载")
        scrollTM(tmall);
    })
    document.body.appendChild(container);

    function checkLogin() {
        var userName = $("#login-info1 .j_Username").text();
        if (!userName) {
            appendTip("请先登录再操作", "red")
        }
        return !!userName;
    }

    function appendTip(html, color) {
        console.log("tools", html)
        $("#clpTip").append(`<div style="color:${color ? color : "#404040"}">${((new Date().valueOf() - clpStartDate) / 1000).toFixed(2)}秒-${html}</div>`)
        $("#clpTip").scrollTop($("#clpTip").prop("scrollHeight"))
    }

    appendTip("抓取提示", "red")


    function saveBlob(blob, name) {
        var a = document.createElement("a");
        a.download = name;
        a.href = URL.createObjectURL(blob);
        a.hidden = true;
        a.click();
        a.remove();
    }

    /**
     * 转一张图片编码
     * @param {string} imgUrl 图片url
     * @return {Promise<string>}
     */
    function getBase64(imgUrl) {
        return new Promise((resolve) => {
            const image = new Image();
            image.crossOrigin = ""; // 解决跨域问题
            image.src = imgUrl;
            image.onload = function () {
                let canvas = document.createElement("canvas");
                canvas.width = image.width;
                canvas.height = image.height;
                let context = canvas.getContext("2d");
                context.drawImage(image, 0, 0, image.width, image.height);
                // 得到图片的base64编码数据
                resolve(canvas.toDataURL("image/png", 1).split(",")[1]);
            };
        });
    }

    async function save(product) {
        console.log("tools", product)
        appendTip("开始打包图片")
        var zip = new JSZip();
        // 创建一个名为images的文件夹
        var zipFolder = zip.folder(product.code);

        for (let i = 0; i < product.mainList.length; i++) {
            var img = product.mainList[i];
            if (!img.url) {
                container;
            }
            var suffix = img.url.substring(img.url.lastIndexOf(".") + 1);
            var imgData = await getBase64(img.url)
            // 3个参数分别是文件名、图片的base64编码、和base64验证
            if (i == 0) {
                zipFolder.file(`main.${suffix}`, imgData, { base64: true });
            }
            zipFolder.file(`${i}_main.${suffix}`, imgData, { base64: true });
        }
        appendTip(`${product.mainList.length}张主图打包完成`)

        for (let i = 0; i < product.descList.length; i++) {
            var img = product.descList[i];
            if (!img.url) {
                container;
            }
            var suffix = img.url.substring(img.url.lastIndexOf(".") + 1);
            var imgData = await getBase64(img.url)
            // 3个参数分别是文件名、图片的base64编码、和base64验证
            zipFolder.file(`description_${i < 9 ? `0${i + 1}` : i + 1}.${suffix}`, imgData, { base64: true });
        }
        appendTip(`${product.descList.length}张详情图打包完成`)
        appendTip(`开始生成压缩包`)
        // 已知问题 当tab不是前置激活窗口时，压缩会暂停 issues见 https://github.com/Stuk/jszip/issues/741
        zip.generateAsync({ type: "blob" }).then(function (content) {
            // 使用FileSaver.js下载到本地
            // saveAs(content, `${product.code}.zip`);
            appendTip(`压缩包生成成功，开始下载`)
            saveBlob(content, `${product.code}.zip`)
            appendTip(`下载完成，文件名：${product.code}.zip`, "green")
            if (productCode != defaultProductCode) {
                appendTip(`批量下载打开的窗口，即将关闭`, "#ffa906")
                setTimeout(() => {
                    window.close();
                }, 1000);
            }
        });
    }

    var delay = 1; //in milliseconds
    var scroll_amount = 1; // in pixels
    //当前页面高度
    var scrollHeight = $("body").outerHeight();
    //获取windows高度
    var windowHeight = $(this).height();
    //已经滚动了的高度
    var scrollTop = $(document).scrollTop();

    // switch (location.host) {
    //     case "detail.tmall.com":
    //         scrollTM(tmall);
    //         break;

    //     case "item.jd.com":
    //         scrollJD(jd);
    //     default:
    //         break;
    // }

    var tmDescRegex = /itemcdn.tmall.com\/desc\/icoss[^\?]+\?var=desc/igm
    var tmDescImgRegex = /src="(https:\/\/img.alicdn.com\/imgextra\/[^"]*)"/igm
    async function scrollTM(callback) {
        let descUrl = null;
        var scripts = document.getElementsByTagName("script")
        for (let i = 0; i < scripts.length; i++) {
            const item = scripts[i];
            console.log("tools", i, item.src, execRes)
            var execRes = tmDescRegex.exec(item.src)
            if (!execRes) {
                execRes = tmDescRegex.exec(item.innerText)
            }
            if (execRes && execRes.length >= 0) {
                descUrl = execRes[0]
                break
            }
        }
        if (descUrl) {
            appendTip("匹配到详情连接，尝试直接通过接口获取")
            const resp = await new Promise((resolve, reject) => {
                GM_xmlhttpRequest({
                    method: 'GET',
                    url: `https://${descUrl}`,
                    onload: function (resp) {
                        resolve(resp)
                    },
                    onerror: function (resp) {
                        reject(resp)
                    }
                });
            });
            console.log("tools", "res", descUrl, resp)
            if (resp.status = 200) {
                var tmDescImg;
                var descList = []
                var i = 0
                while ((tmDescImg = tmDescImgRegex.exec(resp.response)) != null) {
                    var url = tmDescImg[1]
                    descList.push({ url: url, index: i })
                    appendTip(`找到详情图-${i}<br/><img src="${url}" width="280"/>`)
                    console.log("tools", "tmDescImg", url);
                    i++;
                }
                if (descList.length > 1) {
                    tmall(descList)
                    return
                }
            }
            appendTip("通过接口获取详情失败")
        }
        console.log("tools", "尝试使用模拟滚动加载匹配")
        var scrollerEvent = setInterval(() => {
            //当前的页面高度
            scrollHeight = $("body").outerHeight();
            var contentObj = $("#description .content");
            var contentLength = contentObj.html().trim().length;
            console.log("tools", "scrollHeight:" + scrollHeight + "======scrollTop:" + scrollTop + "======windowHeight:" + windowHeight + "======contentLength:" + contentLength);
            if (
                (scrollHeight > scrollTop + windowHeight + 100 &&
                    scrollTop < 2000 &&
                    contentLength < 100) ||
                !contentObj ||
                contentLength < 10
            ) {
                //叠加滚动高度
                scrollTop += scroll_amount;
                $(document).scrollTop(scrollTop);
                appendTip(`没有找到商品详情图片，尝试滚动页面`)
            } else {
                appendTip(`商品详情加载完成，延迟1秒再处理`)
                setTimeout(() => {
                    if (callback) {
                        callback();
                    }
                }, 1000);
                clearInterval(scrollerEvent);
            }
        }, delay);
    }

    function scrollJD(callback) {
        var scrollerEvent = setInterval(() => {
            //当前的页面高度
            scrollHeight = $("body").outerHeight();
            var contentObj = $("#J-detail-content .ssd-module-wrap");
            var contentLength = contentObj.find(".ssd-module").length;
            console.log("tools", "scrollHeight:" + scrollHeight + "======scrollTop:" + scrollTop + "======windowHeight:" + windowHeight + "======contentLength:" + contentLength);
            if (
                (scrollHeight > scrollTop + windowHeight + 100 &&
                    scrollTop < 2000 &&
                    contentLength < 2) ||
                !contentObj ||
                contentLength < 1
            ) {
                //叠加滚动高度
                scrollTop += scroll_amount;
                $(document).scrollTop(scrollTop);
                appendTip(`没有找到商品详情图片，正在模拟滚动页面`)
                // console.log("tools",$("#description .content").html().length);
            } else {
                //  console.log("tools",$("#description .content").html());
                appendTip(`商品详情加载完成`)
                setTimeout(() => {
                    if (callback) {
                        callback();
                    }
                }, 1000);
                clearInterval(scrollerEvent);
            }
        }, delay);
    }

    function tmall(descList) {

        var product = { code: productCode, mainList: [], descList: descList ? descList : [] }
        var mainCount = $("#J_UlThumb img").length;
        $("#J_UlThumb img").each((i, item) => {
            var url = $(item).attr("src")
            url = url.replace('_60x60q90.jpg', '')
            appendTip(`找到主图-${i}<br/><img src="${url}" width="280"/>`)
            product.mainList.push({ url: url, index: i })
        })
        //如果没有通过接口获取到详情内容，则尝试解析页面
        if (!descList?.length) {
            var descCount = $("#description img").length;
            $("#description img").each((i, item) => {
                var url = $(item).attr("data-ks-lazyload")
                url = url ? url : $(item).attr("src");
                appendTip(`找到详情图-${i}<br/><img src="${url}" width="280"/>`)
                product.descList.push({ url: url, index: i })
            })
        }
        appendTip(`识别到主图【${product.mainList.length}】张`)
        appendTip(`识别到详情图【${product.descList.length}】张`)
        save(product)
    }

    function jd() {
        $(".product-intro .lh img").each((i, item) => {
            console.log("tools", "主图", $(item).attr("data-url"))
        })

        var styleHtml = $("#J-detail-content style").html();
        $("#J-detail-content .ssd-module-wrap .ssd-module").each((i, item) => {
            var dataId = $(item).attr("data-id");
            var regex = new RegExp(`.${dataId}\{.*background-image:url\\((.*)\\);.*[^\}]\}`, "igm")
            var execRes = regex.exec(styleHtml)
            if (execRes == null || execRes.length < 2) {
                console.log("tools", "详情图", $(item).attr("data-id"), "抓取失败")
            }
            console.log("tools", "详情图", $(item).attr("data-id"), execRes[1])
        })
    }

    var defaultProductCode = "抓取商品图片信息"
    var productCode = getQueryString('clpProductCode');
    if (productCode) {
        appendTip("批量下载打开的页面，开始自动下载")
        appendTip(`商品编码${productCode}`)
        // setInterval(() => {
        //     appendTip(`批量抓取自动打开的窗口，请勿关闭`, "red")
        // }, 1000);
        scrollTM(tmall);
    } else {
        productCode = defaultProductCode;
    }
})();
