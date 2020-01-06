http://localhost:6060/debug/pprof/

Types of profiles available: 可用的配置文件类型
Count	Profile
14	allocs
0	block
0	cmdline
50	goroutine
14	heap
0	mutex
0	profile
22	threadcreate
0	trace
full goroutine stack dump
Profile Descriptions:

allocs: 对所有过去内存分配的抽样
block: 导致同步原语阻塞的堆栈跟踪
cmdline: 当前程序的命令行调用
goroutine: 所有当前goroutine的堆栈跟踪
heap: 活动对象内存分配的抽样。在获取堆样本之前，可以指定gc GET参数来运行gc。
mutex: 争用互斥锁持有者的堆栈跟踪
profile: CPU配置文件。可以在seconds GET参数中指定持续时间。获得概要文件之后，使用go工具pprof命令研究概要文件。
threadcreate: 导致创建新OS线程的堆栈跟踪
trace: 当前程序的执行轨迹。可以在seconds GET参数中指定持续时间。获取跟踪文件之后，使用go工具跟踪命令来研究跟踪。



 Commands:
    callgrind        输出callgrind格式的图形
    comments         输出所有配置文件注释
    disasm           用示例注释的输出程序集清单
    dot              输出点格式的图形
    eog              通过eog实现图形可视化
    evince           通过evince可视化图形
    gif              输出GIF格式的图形图像
    gv               通过gv可视化图形
    kcachegrind      在KCachegrind可视化报告
    list             为匹配regexp的函数输出带注释的源代码
    pdf              输出PDF格式的图表
    peek             输出与regexp匹配的函数的调用者/callees
    png              输出PNG格式的图形图像
    proto            以压缩的protobuf格式输出概要文件
    ps               输出PS格式的图形
    raw              输出原始配置文件的文本表示形式
    svg              输出SVG格式的图形
    tags             输出概要文件中的所有标记
    text             以文本形式输出顶部条目
    top              以文本形式输出顶部条目
    topproto         以压缩的protobuf格式输出顶部条目
    traces           以文本形式输出所有配置文件样本
    tree             输出调用图的文本呈现
    web              通过web浏览器可视化图形
    weblist          在web浏览器中显示带注释的源代码
    o/options        列出选项及其当前值
    quit/exit/^D     退出pprof

  Options:
    call_tree               创建上下文敏感的调用树
    compact_labels          显示最小的标题
    divide_by               在可视化之前，对所有样本进行比例分割
    drop_negative           忽视消极差异
    edgefraction            隐藏<f>*total的边
    focus                   限制通过匹配regexp的节点的示例
    hide                    跳过与regexp匹配的节点
    ignore                  跳过通过任何与regexp匹配的节点的路径
    mean                    平均样本值除以第一个值(计数)
    nodecount               要显示的最大节点数
    nodefraction            将节点隐藏在<f>*total之下
    noinlines               忽略内联。
    normalize               基于基本轮廓的刻度轮廓。
    output                  基于文件的输出的输出文件名
    prune_from              删除匹配帧下的任何函数。
    relative_percentages    显示相对于聚焦子图的百分比
    sample_index            要报告的示例值(基于0的索引或名称)
    show                    只显示与regexp匹配的节点
    show_from               drop函数位于最高匹配帧之上。
    source_path             搜索源文件的路径
    tagfocus                限制使用范围内的标记或由regexp匹配的示例
    taghide                 跳过与此regexp匹配的标记
    tagignore               使用范围内的标记或regexp匹配的标记丢弃样本
    tagshow                 只考虑与此regexp匹配的标记
    trim                    荣誉nodefraction / edgefraction nodecount违约
    trim_path               在搜索之前从源路径修剪的路径
    unit                    显示测量单位

  Option groups (only set one per group):
    cumulative
      cum              根据累积权重对条目进行排序
      flat             根据权重对条目进行排序
    granularity
      addresses        在地址级别聚合。
      filefunctions    在功能级别聚合。
      files            在文件级别聚合。
      functions        在功能级别聚合。
      lines            在源代码行级别聚合。
  :   Clear focus/ignore/hide/tagfocus/tagignore

  type "help <cmd|option>" for more information


  curl -s http://myapp.cab7b4da4476242bc9ac8743854ce1325.cn-shanghai.alicontainer.com/debug/pprof/heap > base.heap
  curl -s http://myapp.cab7b4da4476242bc9ac8743854ce1325.cn-shanghai.alicontainer.com/debug/pprof/heap > current.heap