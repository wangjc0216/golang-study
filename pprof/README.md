# PProf理解

PProf是分析性能、分析数据的工具，PProf用profile.proto读取分析样本的集合，并生成可视化报告，用以帮助分析数据(支持文本和图形报告) 

profile.proto是一个Protobuf v3的描述文件，它描述了一组callstack和symbolization（符号化）信息，它的作用是统计分析一组采样的调用栈，配置文件格式是stacktrace。

## 采样方式

1. runtine/pprof：采集程序（非server）指定区块的运行数据进行分析

2. net/http/pprof:基于HTTP Server运行，并且可以采集运行时的数据进行分析。
```
在程序中增加：
import (
    _ "net/http/pprof"
)
func main(){
    go http.ListenAndServe("0.0.0.0:6060",nil)
}
```
3. go test： 通过运行测试用例，指定所需标识进行采集。
```
go test -bench=. -cpuprofile=cpu.profile
就会生成cpu.profile
```

## 使用模式

Report Generation：报告生成
Inbteractive Terminal Use： 交互式终端使用
Web Interface：Web界面

## Use Case
比较简单的使用方式是 通过 net/http/pprof 来采集运行时的数据，并进行分析。
```
import (
    _ "net/http/pprof"
)
func main(){
    go http.ListenAndServe("0.0.0.0:6060",nil)
}
```
那么就可以通过访问 http://localhost:6060/debug/pprof/ 地址来查看当前采样。 有多种类型可以对其进行下载，也可以在页面查看（查看的话需要加参数?debug=1）.下载好文件后通过go tool 命令来查看Web界面：
```
go tool pprof -http=:6061 <profile_name>
```

