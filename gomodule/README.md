# Go Modules

## 1. Go Modules介绍
`Go Modules` 是Go1.11 和 Go1.12 才开始逐步引用于管理依赖包版本信息明确和更易于管理的依赖管理系统。相较于Vendor的管理方法(Vendor现在已经废弃了)，依赖包(或者叫依赖模块)的版本信息更加明确，且不需要把依赖包放在项目中做为项目的一部分。

以往Go的一个项目被抽象成模块的概念(就是一个Module)，一个模块会通过根目录下的`go.mod`文件来描述模块的状态，`go.mod`包括两个两个内容：
> module path (模块根目录的导入路径)
> dependency requirements (依赖项要求：也就是其他模块的module path + 语义化版本)

对于Go 1.11来说，如果项目的工作目录不在`$GOPATH/src`里面，并且包含`go.mod`文件，则Go命令行会启动模块机制(相较于GOPATH里执行，此时执行`go get`等操作则会有模块版本的概念)，由于历史原因，如果项目在`$GOPATH/src`中，则Go命令行不会开启模块机制。**但是从Go 1.13开始，`Go Modules` 机制在所有情况下都默认使用。**

## 2. Go Modules相关命令

我们以创建一个模块为示例，从开始初始化模块、增加依赖项、更新依赖项版本、回退依赖项版本、增加没有使用Go Modules管理的依赖项(也就是说增加那些不再维护，还是使用Vendor作为包管理工具的Go项目)、移除依赖、增加私有仓库依赖等步骤，全面展示

### 2.1 初始化模块

​	首先我们创建一个demo模块作用Go Modules 的使用示例。我们创建一个demo目录后在其中分别创建`demo.go`和`demo_test.go`文件。

```
//demo.go
package demo

import "fmt"

func DemoFunc() {
        fmt.Println("this is DemoFunc")
}
//demo_test.go
➜  demo git:(main*)cat demo_test.go 
package demo

import "testing"

func TestDemo(t *testing.T) {
        DemoFunc()
}
```

使用`go mod init`命令初始化该demo模块，并且终端会提示如果需要添加requirements 和 sum 需要执行`go mod tidy`命令。

```
go mod init golang-study/module_demo
go: creating new go.mod: module golang-study/module_demo
go: to add module requirements and sums:
        go mod tidy
```

则会在根目录下生成`go.mod`文件。

```
➜  demo git:(main*)cat go.mod      
module golang-study/module_demo

go 1.16
```

执行`go test`用以执行单元测试，会发现测试结果中也会带有**模块(Go Modules)**的信息。

```
go test    

this is DemoFunc
PASS
ok      golang-study/module_demo        1.915s
```



### 2.2 添加依赖项



### 2.3 更新依赖项版本



### 2.4 回退依赖项版本



### 2.5 增加没有tag的依赖(或者说是使用govendor管理的Go项目)



### 2.6 移除依赖



### 2.7 增加私有仓库依赖





### 2.99 Go Modules 相关命令总结

关于使用Go Modules管理Go项目的命令集中列到下面。

```shell
go mod init

go mod tidy

go list -m all

go mod why -m

go mod graph
```

### 

## 3.Modules一些http请求逻辑

### 查看Modules的版本列表

```
curl -L https://goproxy.cn/github.com/prometheus/prometheus/@v/list  | sort
```
会返回这个module的版本列表。


### 查看Modules版本时间
```
curl -L https://goproxy.cn/github.com/prometheus/prometheus/@v/v2.5.0-rc.2+incompatible.info
```
查看Module具体版本时间

### 查看Module指定版本的go.mod文件
```go
curl -L https://goproxy.cn/github.com/prometheus/prometheus/@v/v2.5.0-rc.2+incompatible.mod
```
当然有一些是有tag，也是在合法的服务器上(如github)进行部署，但是该tag并没有使用Module来管理该Module。所以返回的可能就是如下： 
```go
module github.com/prometheus/prometheus
```

### 下载Module文件
```go
curl -L -O  https://goproxy.cn/github.com/prometheus/prometheus/@v/v2.5.0-rc.2+incompatible.zip
```
这会下载该版本的Module源码zip包，解压后可以看见 `github.com/prometheus/prometheus@v2.5.0-rc.2+incompatible/ …`


pkg主要有两个文件夹，分别为cache和具体版本的源码，cache中会包括 .info .mod .zip list,cache中还包括sumdb

如果想要在本地运行private proxy，可以参考[Athens](https://docs.gomods.io/zh/intro/) ,Authens可以对Module进行缓存，防止公网上数据丢失。

Athens可以通过`docker run -p '3030:3000' --name modserver -d  gomods/athens:latest`启动容器，Athens可以将Modules
缓存下来，防止Module被删除(如github具体仓库被删除)的情况，同时将GOPROXY修改为`localhost:3030`，那么`go get`的时候就会现在从localhost:3030
请求Module数据，如果请求不到，modserver会从官方仓库中去获取Module并缓存本地。






## 4.Go1.16 Modules新特性

```
go env -w GO111MODULE=auto
```


## 相关参考

[Go Modules官方博客](https://blog.golang.org/using-go-modules)
[Go Modules官方博客翻译](https://studygolang.com/articles/19334)
[语义化版本 2.0.0](https://semver.org/lang/zh-CN/)

