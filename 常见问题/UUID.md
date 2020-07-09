# UUID

UUID 是 通用唯一识别码（Universally Unique Identifier）的缩写。目前最广泛应用的UUID，是微软公司的全局唯一标识符（GUID）。UUID是指在一台机器上生成的数字，它保证对在同一时空中的所有机器都是唯一的。通常平台会提供生成的API。按照开放软件基金会(OSF)制定的标准计算，用到了以太网卡地址、纳秒级。

UUID是一个128比特的数值，这个数值可以通过一定的算法计算出来。

GUID xxxxxxxx-xxxx-xxxx-xxxxxxxxxxxxxxxx(8-4-4-16)

标准UUID xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (8-4-4-4-12)

每个 x 是 0-9 或 a-f 范围内的一个十六进制的数字。

UUID的编码规则：

1）1~8位采用系统时间，在系统时间上精确到毫秒级保证时间上的惟一性；

2）9~16位采用底层的IP地址，在服务器集群中的惟一性；

3）17~24位采用当前对象的HashCode值，在一个内部对象上的惟一性；

4）25~32位采用调用方法的一个随机数，在一个对象内的毫秒级的惟一性。
