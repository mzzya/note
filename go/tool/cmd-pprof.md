usage:

Produce output in the specified format.

   pprof <format> [options] [binary] <source> ...

Omit the format to get an interactive shell whose commands can be used
to generate various views of a profile

   pprof [options] [binary] <source> ...

Omit the format and provide the "-http" flag to get an interactive web
interface at the specified host:port that can be used to navigate through
various views of a profile.

   pprof -http [host]:[port] [options] [binary] <source> ...

Details:
  Output formats (select at most one):
    -callgrind       输出callgrind格式的图形
    -comments        输出所有配置文件注释
    -disasm          用示例注释的输出程序集清单
    -dot             输出点格式的图形
    -eog             通过eog实现图形可视化
    -evince          通过evince可视化图形
    -gif             输出GIF格式的图形图像
    -gv              通过gv可视化图形
    -kcachegrind     在KCachegrind可视化报告
    -list            为匹配regexp的函数输出带注释的源代码
    -pdf             输出PDF格式的图表
    -peek            输出与regexp匹配的函数的调用者/callees
    -png             输出PNG格式的图形图像
    -proto           以压缩的protobuf格式输出概要文件
    -ps              输出PS格式的图形
    -raw             输出原始配置文件的文本表示形式
    -svg             输出SVG格式的图形
    -tags            输出概要文件中的所有标记
    -text            以文本形式输出顶部条目
    -top             以文本形式输出顶部条目
    -topproto        以压缩的protobuf格式输出顶部条目
    -traces          以文本形式输出所有配置文件样本
    -tree            输出调用图的文本呈现
    -web             通过web浏览器可视化图形
    -weblist         在web浏览器中显示带注释的源代码

  Options:
    -call_tree              创建上下文敏感的调用树
    -compact_labels         显示最小的标题
    -divide_by              在可视化之前，对所有样本进行比例分割
    -drop_negative          忽视消极差异
    -edgefraction           隐藏<f>*total的边
    -focus                  限制通过匹配regexp的节点的示例
    -hide                   跳过与regexp匹配的节点
    -ignore                 跳过通过任何与regexp匹配的节点的路径
    -mean                   平均样本值除以第一个值(计数)
    -nodecount              要显示的最大节点数
    -nodefraction           将节点隐藏在<f>*total之下
    -noinlines              忽略内联。
    -normalize              基于基本轮廓的刻度轮廓。
    -output                 基于文件的输出的输出文件名
    -prune_from             删除匹配帧下的任何函数。
    -relative_percentages   显示相对于聚焦子图的百分比
    -sample_index           要报告的示例值(基于0的索引或名称)
    -show                   只显示与regexp匹配的节点
    -show_from              drop函数位于最高匹配帧之上。
    -source_path            搜索源文件的路径
    -tagfocus               限制使用范围内的标记或由regexp匹配的示例
    -taghide                跳过与此regexp匹配的标记
    -tagignore              使用范围内的标记或regexp匹配的标记丢弃样本
    -tagshow                只考虑与此regexp匹配的标记
    -trim                   荣誉nodefraction / edgefraction nodecount违约
    -trim_path              在搜索之前从源路径修剪的路径
    -unit                   显示测量单位

  Option groups (only set one per group):
    cumulative
      -cum             Sort entries based on cumulative weight
      -flat            Sort entries based on own weight
    granularity
      -addresses       Aggregate at the address level.
      -filefunctions   Aggregate at the function level.
      -files           Aggregate at the file level.
      -functions       Aggregate at the function level.
      -lines           Aggregate at the source code line level.

  Source options:
    -seconds              基于时间的概要文件收集的持续时间
    -timeout              配置文件收集的超时(秒)
    -buildid              覆盖主二进制文件的构建id
    -add_comment          要添加到概要文件的自由格式注释
                          显示在一些报告或与教授的意见
    -diff_base source     用于比较的基本概要的来源
    -base source          配置文件减法的基本配置文件的来源
    profile.pb.gz         压缩的protobuf格式的配置文件
    legacy_profile        在遗留的pprof格式的配置文件
    http://host/profile   配置文件处理程序要检索的URL
    -symbolize=           控制符号信息的来源
      none                不尝试符号化
      local               只检查本地二进制文件
      fastlocal           只从本地二进制文件中获取函数名
      remote              不检查本地二进制文件
      force               力re-symbolization
    Binary                用于符号化的二进制本地路径或构建id
    -tls_cert             用于获取配置文件和符号的TLS客户端证书文件
    -tls_key              用于获取配置文件和符号的TLS私钥文件
    -tls_ca               用于获取配置文件和符号的TLS CA证书文件

  Misc options:
   -http              在主机端口提供web接口。
                      Host是可选的，默认是'localhost'。
                      端口是可选的，默认情况下是随机可用的端口。
   -no_browser        跳过为交互式web UI打开浏览器。
   -tools             搜索对象工具的路径

  Legacy convenience options: 遗留的方便选择
   -inuse_space           Same as -sample_index=inuse_space
   -inuse_objects         Same as -sample_index=inuse_objects
   -alloc_space           Same as -sample_index=alloc_space
   -alloc_objects         Same as -sample_index=alloc_objects
   -total_delay           Same as -sample_index=delay
   -contentions           Same as -sample_index=contentions
   -mean_delay            Same as -mean -sample_index=delay

  Environment Variables:
   PPROF_TMPDIR       Location for saved profiles (default $HOME/pprof)
   PPROF_TOOLS        Search path for object-level tools
   PPROF_BINARY_PATH  Search path for local binary files
                      default: $HOME/pprof/binaries
                      searches $name, $path, $buildid/$name, $path/$buildid
   * On Windows, %USERPROFILE% is used instead of $HOME