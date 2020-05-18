# golang常见问题

## slice

- make([]int,0,10) 第二个参数是实际长度，第三个参数是容量，底层实现是数组
- 扩容：1024临界点，小于翻倍，大于1.25倍

## map

- make 创建只有2个参数，1KV类型，2容量
- for range map 为什么是无序的，因为最初设计在初始化时会生成一个随机偏移量。
  - 日常使用对key进行排序，然后在遍历keys数组去使用。

## sync.Mutex和sync.RWMutex

- sync.Mutex 互斥锁
- sync.RWMutex 读写锁
  - 读读不互斥，读写，写写互斥。特例：简单的main函数测试时读读是互斥的，跟GPM模型有关。

## channel

- `make(chan int)`创建无缓存双向通道，加数量参数为有缓存通道
- 双向channel可以传递个单向channel，单向channel一般作为是作为函数参数限制调用
- 读写都是并发安全的，名言`不要以共享内存的方式通信，而是通过通信的方式共享内存`
- 写入方使用完毕需记得close
  - 不能多次close,2+次 panic
  - close之后不会影响读取chnnel，有缓存的会继续读完
- for+select 经常搭配使用，轮询从通道获取数据并处理
  - timer.After 返回一个时间类型channel
  - ctx.Done 返回一个结构体类型channel

## context

- context.WithValue 基于struct链式结构
- context.Background与context.TODO区别
  - 实现没有区别都是基于类型为int的emptyCtx
  - Background()一般作为请求顶级入口使用，TODO的话是作为暂时填充后期要完善来使用的，会被一些工具检测出来，比如vscode插件。
- context.WithTimeout与context.WithDeadline
  - WithTimeout接受时间间隔，WithDeadline接受具体时间，WithTimeout调用的是WithDeadline
- 1.13.1~1.13.4存在一个bug，http server 中为每个链接创建ctx是基于全局定义的ctx做的，导致内存泄漏。


### GPM

- G goroutine
- P processer
- M mathine



## 相关文档

[官方文档](https://golang.google.cn/doc/)