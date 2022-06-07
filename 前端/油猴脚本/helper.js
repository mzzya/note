(function () {
    'use strict';

    //解决zip压缩暂停问题
    let setImmediate = (fn) => { fn(); };
    window.setImmediate = setImmediate;
    console.log("window.setImmediate", setImmediate, window.setImmediate)
})();
