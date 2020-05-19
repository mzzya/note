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

相关问题

- M0和G0主要负责调度器的运作的
- 本地P队列没有G时则从全局G队列获取，没有再从别的P队列偷取一半的G来执行
- 简单main函数开10个go对变量递增，数据与期望一致是因为G无阻塞，调度器就没有再创建P去关联别的M执行
- M数与内核线程数有关，P数量按需创建并关联M，假设M被阻塞了，会将P调度到别的M上执行
- coroutine是协助式的，需要主动让住资源。goroution有调度器调度，最多拥有10ms
- 新的G进来会先写入到P队列，P队列中超过256个G后会写入到全局G队列
- 第一版调度器没有P，导致所有M从全局G队列获取时存在竞态问题，加锁导致性能下降
- 多个goroutione中的通信用channel来同步数据，它是线程安全的



## 相关文档

[官方文档](https://golang.google.cn/doc/)