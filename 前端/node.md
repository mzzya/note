# node 相关

## 私有包配置

私有包最好放在一个`scope`，可以理解成namespace。

```json
//package.json
{
  "name": "@hellojqk/tools"
}
```

这里的`scope`即为`@hellojqk`

```sh
# 登录
npm login --scope=@hellojqk --always-auth --registry=你的私有地址（例如：nexus搭建的独立hosted的npm仓库）

# 发布
npm publish
```

通过`cat ~/.npmrc`我们可以看到类似下方的内容

```conf
registry=https://registry.npm.taobao.org/
//registry.***.com/repository/npm-hosted/:_authToken=NpmToken.******************************
@hellojqk:registry=https://registry.***.com/repository/npm-hosted/
```

在`CI`中可以通过`npm cofig set`将上方参数写入进去，这样即可实现免密登录。


## 常见问题

1. npm包明明安装过了，构建时总报找不到某些包。
   1. 删除node_modules然后重新install
2. npm publish 401错误
   1. 如果是推送到npm官方仓库，需要publish指定`--registry`参数。可能是配置了加速源导致的。
   2. 如果是推送到nexus私有仓库，建议按上方`私有包配置`章节操作。
