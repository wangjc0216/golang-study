# Go Modules

## Go Modules介绍
`Go Modules` 是Go1.11 和 Go1.12 才开始逐步引用于管理依赖包版本信息明确和更易于管理的依赖管理系统。相较于Vendor的管理方法(Vendor现在已经废弃了)，依赖包(或者叫依赖模块)的版本信息更加明确，且不需要把依赖包放在项目中做为项目的一部分。

以往Go的一个项目被抽象成模块的概念(就是一个Module)，一个模块会通过根目录下的`go.mod`文件来描述模块的状态，`go.mod`包括两个两个内容：
> module path (模块根目录的导入路径)
> dependency requirements (依赖项要求：也就是其他模块的module path + 语义化版本)

对于Go 1.11来说，如果项目的工作目录不在`$GOPATH/src`里面，并且包含`go.mod`文件，则Go命令行会启动模块机制(相较于GOPATH里执行，此时执行`go get`等操作则会有模块版本的概念)，由于历史原因，如果项目在`$GOPATH/src`中，则Go命令行不会开启模块机制。**但是从Go 1.13开始，`Go Modules` 机制在所有情况下都默认使用。**

## 相关命令

关于使用Go Modules管理Go项目的命令集中列到下面。

```shell
//
go list -m all

go mod why -m

go mod graph

```

## 相关参考

[Go官方博客](https://blog.golang.org/using-go-modules)

[语义化版本 2.0.0](https://semver.org/lang/zh-CN/)

