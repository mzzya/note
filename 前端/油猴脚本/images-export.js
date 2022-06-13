// ==UserScript==
// @name         图片导出
// @namespace    http://tampermonkey.net/
// @version      0.3
// @description  (本脚本仅供学习开发油猴脚本使用，请勿用于商业目的)支持淘宝、天猫（含超市、国际）、京东（含超市、国际、医药）等平台商品图片导出，多平台导出时，请先登录各平台，可以提高成功率。
// @author       hellojqk
// @match        http*://detail.tmall.com/*
// @match        http*://chaoshi.detail.tmall.com/*
// @match        http*://detail.tmall.hk/*
// @match        http*://item.taobao.com/*
// @match        http*://item.jd.com/*
// @match        http*://npcitem.jd.hk/*
// @match        http*://item.yiyaojd.com/*
// @match        http*://item.jkcsjd.com/*
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
// @connect      cd.jd.com
// ==/UserScript==


//参考文档
//js压缩下载文件 https://zhuanlan.zhihu.com/p/352198022
//jszip github https://github.com/Stuk/jszip#readme
//jscdn https://cdn.bytedance.com/

(function () {
    'use strict';

    /**
     * 处理标题特殊符号
     * @param {string} title
     * @returns
     */
    function trimTitle(title) {
        if (title) {
            title = title.trim().replace(/[<>:"\/\\|?*]/g, "_").replace(/_+/g, '_');
        }
        if (!title) {
            title = "抓取商品信息"
        }
        return title;
    }

    // 默认产品编码
    let defaultProductCode = "";
    // 站点类型，用于处理详情
    let siteType = "";
    // 处理默认产品编码
    switch (location.host.toLowerCase()) {
        case "item.taobao.com":
            siteType = "tb";
            defaultProductCode = $('#J_Title .tb-main-title').attr("data-title");
            defaultProductCode = trimTitle(defaultProductCode)
            break;
        case "detail.tmall.com":
        case "chaoshi.detail.tmall.com":
        case "detail.tmall.hk":
            siteType = "tmall";
            defaultProductCode = $('input[name="title"]').val();
            defaultProductCode = trimTitle(defaultProductCode)
            break;
        case "item.jd.com":
        case "npcitem.jd.hk":
        case "item.yiyaojd.com":
        case "item.jkcsjd.com":
            siteType = "jd"
            defaultProductCode = $(".product-intro .sku-name").text()
            defaultProductCode = trimTitle(defaultProductCode)
            break;
        default:
            appendTip("不支持当前页面")
            break;
    }

    // 下载
    function download() {
        appendTip("开始下载")
        switch (siteType) {
            case "tb":
            case "tmall":
                tmPageExecHandler();
                break;
            case "jd":
                jdPageExecHandler();
                break;
            default:
                appendTip("不支持当前页面")
                break;
        }
    }

    let clpStartDate = new Date().valueOf();

    /**
     * 获取url参数值
     * @param {string} name
     * @returns
     */
    function getQueryString(name) {
        let reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
        let r = window.location.search.substr(1).match(reg);
        if (r != null) {
            return decodeURIComponent(r[2]);
        }
        return null;
    }


    let container = document.createElement('div');
    container.id = 'clpContainer'
    container.style = 'position: fixed;z-index: 9999999999;right: 6px; top:0px;'
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
        <textarea id="clpBatchContent" style="" placeholder="支持淘宝、天猫（含超市、国际）、京东（含超市、国际、医药）等平台商品图片导出，多平台导出时，请先登录各平台，可以提高成功率。

请按照如下格式录入:
产品编码1 抓取连接1
产品编码2 抓取连接2"></textarea>
        <button class="clpBtn" id="clpBatchStart">开始抓取</button>
    </div>
</div>`
    //批量下载
    $("body").on("click", "#clpBatchDownload", function () {
        if (!checkLogin()) {
            return;
        }
        let display = $("#clpBatchDiv").css("display");
        $("#clpBatchDiv").css("display", display == "flex" ? "none" : "flex")
    })
    //批量下载开始
    $("body").on("click", "#clpBatchStart", function () {
        clpStartDate = new Date().valueOf();
        let content = $("#clpBatchContent").val();
        if (!content) {
            appendTip(`请输入要抓取的商品编码和链接`, "red");
            return;
        }
        let list = content.split("\n")
        if (!list) {
            appendTip(`输入的内容不正确`, "red");
            return
        }
        let productList = []
        for (let i = 0; i < list.length; i++) {
            const line = list[i].trim();
            if (!line) {
                appendTip(`输入的第${i + 1}行为空，将跳过`, "red");
                continue;
            }
            let lineInfo = line.split(' ')
            if (lineInfo.length != 2) {
                appendTip(`输入的第${i + 1}行未能同时解析到编码和下载链接`, "red");
                continue;
            }
            let code = lineInfo[0];
            let url = lineInfo[1];
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
                let openUrl = `${product.url}${product.url.indexOf("?") < 0 ? "?" : "&"}clpProductCode=${encodeURI(product.code)}`
                console.log("tools", "openUrl", openUrl)
                let newTap = GM_openInTab(openUrl, { active: true, setParent: true });

            }, i * 10000);
        }
    })
    //下载当前页图片
    $("body").on("click", "#clpDownload", function () {
        if (!checkLogin()) {
            return;
        }
        let clpCode = $("#clpCode").val();
        if (!clpCode) {
            productCode = defaultProductCode;
            appendTip(`没有输入产品编码，将使用【${productCode}】作为压缩包名称`, "red");
        } else {
            productCode = clpCode;
        }
        clpStartDate = new Date().valueOf();
        download();
    })
    document.body.appendChild(container);

    /**
     * 检查登录
     * @returns
     */
    function checkLogin() {
        let userName = "";
        switch (siteType) {
            case "tb":
                userName = $("#J_SiteNavBdL .site-nav-login-info-nick").text();
                break;
            case "tmall":
                userName = $("#login-info1 .j_Username").text();
                break;
            case "jd":
                userName = $("#ttbar-login .nickname").text();
                break;
        }
        if (!userName) {
            appendTip("请先登录再操作", "red")
        }
        return !!userName;
    }

    /**
     * 写日志
     * @param {string} html
     * @param {string} color
     */
    function appendTip(html, color) {
        console.log("tools", html)
        $("#clpTip").append(`<div style="color:${color ? color : "#404040"}">${((new Date().valueOf() - clpStartDate) / 1000).toFixed(2)}秒-${html}</div>`)
        $("#clpTip").scrollTop($("#clpTip").prop("scrollHeight"))
    }

    appendTip("抓取提示", "red")


    /**
     * 下载文件
     * @param {Blob} blob
     * @param {string} name
     */
    function saveBlob(blob, name) {
        let a = document.createElement("a");
        a.download = name;
        a.href = URL.createObjectURL(blob);
        a.hidden = true;
        a.click();
        a.remove();
    }

    /**
     * 图片url转base64
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

    /**
     * 保存商品
     * @param {Object} product
     */
    async function save(product) {
        console.log("tools", product)
        appendTip("开始打包图片")
        let zip = new JSZip();
        let isDefaultName = product.code == defaultProductCode;
        // 创建一个名为images的文件夹
        let zipFolder = isDefaultName ? zip.folder(product.code) : null;
        let mainFolder = (isDefaultName ? zipFolder : zip).folder("main");
        let descriptionFolder = (isDefaultName ? zipFolder : zip).folder("description");
        for (let i = 0; i < product.mainList.length; i++) {
            let img = product.mainList[i];
            if (!img.url) {
                container;
            }
            let suffix = img.url.substring(img.url.lastIndexOf(".") + 1);
            let imgData = await getBase64(img.url)
            // 3个参数分别是文件名、图片的base64编码、和base64验证
            if (i == 0) {
                mainFolder.file(`${isDefaultName ? "main" : product.code}.${suffix}`, imgData, { base64: true });
            }
            mainFolder.file(`${i}_${isDefaultName ? "main" : product.code}.${suffix}`, imgData, { base64: true });
        }
        appendTip(`${product.mainList.length}张主图打包完成`)

        for (let i = 0; i < product.descList.length; i++) {
            let img = product.descList[i];
            if (!img.url) {
                container;
            }
            let suffix = img.url.substring(img.url.lastIndexOf(".") + 1);
            let imgData = await getBase64(img.url)
            // 3个参数分别是文件名、图片的base64编码、和base64验证
            descriptionFolder.file(`${isDefaultName ? "description" : product.code}_${i < 9 ? `0${i + 1}` : i + 1}.${suffix}`, imgData, { base64: true });
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

    let delay = 1; //in milliseconds
    let scroll_amount = 1; // in pixels
    //当前页面高度
    let scrollHeight = $("body").outerHeight();
    //获取windows高度
    let windowHeight = $(this).height();
    //已经滚动了的高度
    let scrollTop = $(document).scrollTop();

    /**
     * 获取详情内容接口的正则
     * @param {Regex} regex
     * @returns
     */
    function getDescApiUrl(regex) {
        let descUrl = null;
        let scripts = document.getElementsByTagName("script")
        for (let i = 0; i < scripts.length; i++) {
            const item = scripts[i];
            console.log("tools", i, item.src, execRes)
            let execRes = regex.exec(item.src)
            if (!execRes) {
                execRes = regex.exec(item.innerText)
            }
            if (execRes && execRes.length >= 0) {
                descUrl = execRes[0]
                break
            }
        }
        return descUrl;
    }

    async function getDescContent(regex) {
        let descUrl = getDescApiUrl(regex)
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
            return resp;
        }
        console.log("tools", "res", descUrl, resp)
        return null;
    }

    /**
     * 匹配天猫详情内容接口的正则
     */
    let tmDescApiUrlRegex = /itemcdn.tmall.com\/desc\/icoss[^\?]+\?let=desc/igm
    /**
     * 匹配天猫详情内容接口内容中商品图片的正则
     */
    let tmDescImageRegex = /src="(https:\/\/img.alicdn.com\/imgextra\/[^"]*)"/igm

    /**
     * 天猫页面处理函数
     * @returns
     */
    async function tmPageExecHandler() {

        const resp = await getDescContent(tmDescApiUrlRegex)

        if (resp && resp.status == 200) {
            let tmDescImg;
            let descList = []
            let i = 0
            while ((tmDescImg = tmDescImageRegex.exec(resp.response)) != null) {
                let url = tmDescImg[1]
                descList.push({ url: url, index: i })
                appendTip(`找到详情图-${i}<br/><a href="${url}" title="${url}" target="_blank"><img src="${url}" width="280"/></a>`)
                console.log("tools", "tmDescImg", url);
                i++;
            }
            if (descList.length > 1) {
                generateTmallProduct(descList)
                return
            }
        }

        appendTip("通过接口获取详情失败")
        console.log("tools", "尝试使用模拟滚动加载匹配")

        let scrollerEvent = setInterval(() => {
            //当前的页面高度
            scrollHeight = $("body").outerHeight();
            let contentObj = $("#description .content");
            let contentLength = contentObj.html().trim().length;
            console.log("tools", "scrollHeight:" + scrollHeight + "======scrollTop:" + scrollTop + "======windowHeight:" + windowHeight + "======contentLength:" + contentLength);
            if (
                (scrollHeight > scrollTop + windowHeight + 100 && scrollTop < 2000 && contentLength < 100) ||
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
                    generateTmallProduct();
                }, 1000);
                clearInterval(scrollerEvent);
            }
        }, delay);
    }


    /**
     * 生成天猫商品
     * @param {Object[]} descList
     */
    function generateTmallProduct(descList) {
        let product = { code: productCode, mainList: [], descList: descList ? descList : [] }

        //获取主图
        $("#J_UlThumb img").each((i, item) => {
            let url = $(item).attr("src")
            if (!url) {
                appendTip(`获取第${i}张主图失败`, "red")
                return
            }
            url = url.replace(/_\d+x\d+.*\.(jpg|webp)/, '')
            appendTip(`找到主图-${i}<br/><a href="${url}" title="${url}" target="_blank"><img src="${url}" width="280"/></a>`)
            product.mainList.push({ url: url, index: i })
        })

        //如果没有通过接口获取到详情内容，则尝试解析页面
        if (!descList?.length) {
            let descCount = $("#description img").length;
            $("#description img").each((i, item) => {
                let url = $(item).attr("data-ks-lazyload")
                url = url ? url : $(item).attr("src");
                appendTip(`找到详情图-${i}<br/><a href="${url}" title="${url}" target="_blank"><img src="${url}" width="280"/></a>`)
                product.descList.push({ url: url, index: i })
            })
        }

        appendTip(`识别到主图【${product.mainList.length}】张`)
        appendTip(`识别到详情图【${product.descList.length}】张`)
        save(product)
    }

    /**
     * 匹配京东详情内容接口的正则
     * cd.jd.com/description/channel?skuId=100035246704&mainSkuId=100035246704&charset=utf-8&cdn=2
     */
    let jdDescApiUrlRegex = /cd.jd.com\/description\/channel\?skuId=.+&charset=utf-8&cdn=2/igm

    /**
     * 京东页面处理函数
     * @returns
     */
    async function jdPageExecHandler() {
        const resp = await getDescContent(jdDescApiUrlRegex)

        if (resp && resp.status == 200) {
            let descList = []
            let respData = JSON.parse(resp.response)
            let descHtml = `<div>${respData.content}</div>`

            //主要是针对商城主要产品
            $(descHtml).find(".ssd-module-wrap .ssd-module").each((i, item) => {
                let dataId = $(item).attr("data-id");
                let regex = new RegExp(`.${dataId}\{.*background-image:url\\((.*)\\);.*[^\}]\}`, "igm")
                let descImg = regex.exec(descHtml)
                if (!descImg || descImg.length < 2) {
                    return true;
                }
                let url = descImg[1]
                descList.push({ url: url, index: i })
                appendTip(`找到详情图-${i}<br/><a href="${url}" title="${url}" target="_blank"><img src="${url}" width="280"/></a>`)
                console.log("tools", "tmDescImg", url);
            })

            //针对京东医药等兼容处理
            if (descList.length == 0) {
                $(descHtml).find("img").each((i, item) => {
                    let url = $(item).attr("data-lazyload")
                    descList.push({ url: url, index: i })
                    appendTip(`找到详情图-${i}<br/><a href="${url}" title="${url}" target="_blank"><img src="${url}" width="280"/></a>`)
                    console.log("tools", "tmDescImg", url);
                })
            }

            if (descList.length > 1) {
                generateJDProduct(descList)
                return
            }
        }

        appendTip("通过接口获取详情失败")
        console.log("tools", "尝试使用模拟滚动加载匹配")

        let scrollerEvent = setInterval(() => {
            //当前的页面高度
            scrollHeight = $("body").outerHeight();
            let contentObj = $("#J-detail-content .ssd-module-wrap");
            let contentLength = contentObj.find(".ssd-module").length;
            console.log("tools", "scrollHeight:" + scrollHeight + "======scrollTop:" + scrollTop + "======windowHeight:" + windowHeight + "======contentLength:" + contentLength);
            if (
                (scrollHeight > scrollTop + windowHeight + 100 && scrollTop < 2000 && contentLength < 2) ||
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
                    generateJDProduct();
                }, 1000);
                clearInterval(scrollerEvent);
            }
        }, delay);
    }

    /**
    * 生成京东商品
    * @param {Object[]} descList
    */
    function generateJDProduct(descList) {
        let product = { code: productCode, mainList: [], descList: descList ? descList : [] }

        //获取主图
        $(".product-intro #spec-list .lh img").each((i, item) => {
            let url = $(item).attr("data-url");
            if (!url) {
                appendTip(`获取第${i}张主图失败`, "red")
                return
            }
            url = "https://img14.360buyimg.com/n1/s800x800_" + url.replace(/\.(avif)/, '')
            appendTip(`找到主图-${i}<br/><a href="${url}" title="${url}" target="_blank"><img src="${url}" width="280"/></a>`)
            product.mainList.push({ url: url, index: i })
        })

        //如果没有通过接口获取到详情内容，则尝试解析页面
        if (!descList?.length) {
            let styleHtml = $("#J-detail-content style").html();
            $("#J-detail-content .ssd-module-wrap .ssd-module").each((i, item) => {
                let dataId = $(item).attr("data-id");
                let regex = new RegExp(`.${dataId}\{.*background-image:url\\((.*)\\);.*[^\}]\}`, "igm")
                let descImg = regex.exec(styleHtml)
                if (!descImg || descImg.length < 2) {
                    return true;
                }
                let url = descImg[1]
                product.descList.push({ url: url, index: i })
                appendTip(`找到详情图-${i}<br/><a href="${url}" title="${url}" target="_blank"><img src="${url}" width="280"/></a>`)
                console.log("tools", "tmDescImg", url);
            })
        }
        appendTip(`识别到主图【${product.mainList.length}】张`)
        appendTip(`识别到详情图【${product.descList.length}】张`)
        save(product)
    }

    console.log("tools", "默认商品名称", defaultProductCode)
    let productCode = getQueryString('clpProductCode');
    if (productCode) {
        appendTip("批量下载打开的页面，开始自动下载")
        appendTip(`商品编码${productCode}`)
        setTimeout(() => {
            download()
        }, 1000);
    } else {
        productCode = defaultProductCode;
    }
})();
