# docker问题

## Cmd 与 Entrypoint

都是最终执行命令，一般用于应用进程的启动。

### 格式

- 都支持
  - CMD|ENTRYPOINT ["executable","param1","param2"]
  - CMD|ENTRYPOINT command param1 param2 此格式
- 区别支持
  - CMD ["param1","param2"] (as default parameters to ENTRYPOINT)

- 多`CMD`仅最后一个有效
- `CMD`可做为默认执行设置
- 同时定义`CMD`和`Entrypoint`会执行Entrypoint并执行CMD
- `--entrypoint`覆盖`Dockerfile`中的定义


### 坑

- [ "echo", "$HOME" ] 无效
- [ "sh", "-c", "echo $HOME" ] 有效

`docker inspect alpine:3.11.6` 可以看到只有`Cmd`没有`Entrypoint`

## 相关资料

[dockerfile](https://docs.docker.com/engine/reference/builder)