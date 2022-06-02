// ==UserScript==
// @name         auto-auth
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  desc
// @author       author
// @match        http*://*.demo.com/*
// @match        http*://localhost:*/*
// @match        http*://127.0.0.1:*/*
// @grant        GM_cookie
// @grant        GM_setClipboard
// @grant        GM_setValue
// @grant        GM_getValue
// @grant        GM_xmlhttpRequest
// @grant        GM_registerMenuCommand
// @grant        GM_openInTab
// @run-at       document-start
// @updateURL    https://demo.com/tampermonkey/auto-auth.js
// @downloadURL  https://demo.com/tampermonkey/auto-auth.js
// ==/UserScript==

(function () {
    'use strict';
  
    // 环境和ticket映射
    var envMap = {
      dev: {
        ticketName: 'dev_token',
        home: 'http://dev.demo.com/',
        loginUrl: 'http://devssoapi.demo.com',
      },
      test: {
        ticketName: 'test_token',
        home: 'https://test.demo.com/',
        loginUrl: 'https://testssoapi.demo.com',
      },
      uat: {
        ticketName: 'uat_token',
        home: 'https://uat.demo.com/',
        loginUrl: 'https://uatssoapi.demo.com',
      },
      prd: {
        ticketName: '_token',
        home: 'https://.demo.com/',
        loginUrl: 'https://ssoapi.demo.com',
      },
    };
  
    console.log('ssoTicket 当前页面地址', location.href);
    let ssoTicketRedirectUrl = getQueryString('ssoTicketRedirectUrl');
    if (ssoTicketRedirectUrl) {
      setTimeout(() => {
        location.href = ssoTicketRedirectUrl;
      }, 3000);
    }
  
    // 当前执行环境 dev:开发 test：测试 uat:验收 prd:生产
    var activeEnv = GM_getValue('env', 'test');
    console.log(
      'ssoTicket 当前环境',
      activeEnv,
      GM_getValue(envMap[activeEnv].ticketName),
    );
    // 在匹配到如下正则页面启用该脚本
    var pagePatterns = [/127.0.0.1:*\d*/, /localhost:*\d*/];
  
    // 页面请求符合这些正则的接口（fetch和xhr类型）时不注入Authorization 其他全部注入
    var skipAPIPatterns = [/.*.baidu.com/, /.*.aliyuncs.com/];
  
    function setEnv(env) {
      activeEnv = env;
      GM_setValue('env', env);
      var ticket = GM_getValue(envMap[activeEnv].ticketName);
      if (!ticket) {
        gotoLogin();
        return;
      }
      alert(`设置环境为${env},ticket:${ticket}`);
      location.reload();
    }
  
    GM_registerMenuCommand('开发环境', setEnv.bind(this, 'dev'));
    GM_registerMenuCommand('测试环境', setEnv.bind(this, 'test'));
    GM_registerMenuCommand('UAT环境', setEnv.bind(this, 'uat'));
    GM_registerMenuCommand('生产环境', setEnv.bind(this, 'prd'));
    GM_registerMenuCommand('清除插件缓存的Ticket', () => {
      Object.keys(envMap).forEach((env) => {
        GM_setValue(envMap[env].ticketName, '');
      });
      alert(`已清除所有ticket信息，请先访问在工作台`);
    });
  
    // 缓存ticket
    function storeTicket(ticketName) {
      GM_cookie.list(
        { name: ticketName, domain: '.demo.com' },
        (cookie, error) => {
          if (error) {
            console.log(
              'ssoTicket GM_cookie.list error',
              ticketName,
              cookie,
              error,
            );
            return;
          }
          if (!cookie.length) {
            console.log(
              'ssoTicket GM_cookie.list miss',
              ticketName,
              cookie,
              error,
            );
            return;
          }
          GM_setValue(ticketName, cookie[0].value);
          console.log('ssoTicket 缓存ticket', ticketName, cookie[0].value);
        },
      );
    }
  
    Object.keys(envMap).forEach((env) => {
      storeTicket(envMap[env].ticketName);
    });
  
    pagePatterns.forEach(async (pattern) => {
      var matchRes = pattern.test(document.location.host);
      console.log('ssoTicket pattern', pattern, matchRes);
      if (matchRes) {
        var ticket = GM_getValue(envMap[activeEnv].ticketName);
        if (!ticket) {
          gotoLogin();
          return;
        }
        filterFetch();
        filterXHR();
        await unsafeWindow
          .fetch(envMap[activeEnv].loginUrl + '/api/backend/user/menu')
          .then((res) => {
            switch (res.status) {
              case 401:
                console.log('ssoTicket 校验ticket无效', res);
                gotoLogin();
                break;
              case 200:
                console.log('ssoTicket 校验ticket有效', res);
                break;
            }
          })
          .catch((err) => {
            console.log('ssoTicket 校验ticket异常', err);
          });
      }
    });
  
    function gotoLogin() {
      alert(`请先前往${envMap[activeEnv].home}缓存当前环境ticket`);
      location.href =
        envMap[activeEnv].home +
        '?ssoTicketRedirectUrl=' +
        encodeURIComponent(location.origin + location.pathname);
    }
  
    function matchApis(pagePatterns, apiUrl) {
      if (!apiUrl) {
        return true;
      }
      if (!apiUrl.startsWith('http') && !apiUrl.startsWith('/')) {
        return true;
      }
      var res = false;
      for (let i = 0; i < pagePatterns.length; i++) {
        const pattern = pagePatterns[i];
        res = res || pattern.test(apiUrl);
      }
      return res;
    }
  
    function getQueryString(name) {
      var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
      var r = window.location.search.substr(1).match(reg);
      if (r != null) {
        return unescape(r[2]);
      }
      return null;
    }
  
    // 处理fetch请求的header注入
    function filterFetch() {
      const originFetch = fetch;
      unsafeWindow.fetch = (input, init) => {
        var matchRes = matchApis(skipAPIPatterns, input);
        // console.log('ssoTicket filterFetch', matchRes, input, init);
        if (matchRes) {
          return originFetch(input, init);
        }
  
        let modifyInit = init;
  
        if (!modifyInit) {
          modifyInit = {};
        }
        if (!modifyInit.headers) {
          modifyInit.headers = {};
        }
        let ticket = GM_getValue(envMap[activeEnv].ticketName, '');
        console.log(
          'ssoTicket 附加',
          envMap[activeEnv].ticketName,
          ticket,
          input,
        );
        modifyInit.headers.Authorization = ticket;
        return originFetch(input, modifyInit);
      };
    }
  
    // 处理xhr请求的header注入
    function filterXHR() {
      let req = XMLHttpRequest;
      (function (open, send) {
        XMLHttpRequest.prototype.open = function (method, url) {
          this._url = url;
          open.apply(this, arguments);
        };
        XMLHttpRequest.prototype.send = function () {
          var matchRes = matchApis(skipAPIPatterns, this._url);
          // console.log('ssoTicket filterXHR', this._url, matchRes);
          if (matchRes) {
            send.apply(this, arguments);
            return;
          }
  
          let ticket = GM_getValue(envMap[activeEnv].ticketName, '');
          console.log(
            'ssoTicket filterXHR 附加',
            envMap[activeEnv].ticketName,
            ticket,
            this._url,
          );
          this.setRequestHeader('Authorization', ticket);
          send.apply(this, arguments);
        };
      })(req.prototype.open, XMLHttpRequest.prototype.send);
    }
  })();
  