# proto

## install
1. 安装protoc 

在[下载页面](https://github.com/protocolbuffers/protobuf/releases),下载适用于本机的压缩包，下载并解压后：
```
sudo mv bin/* /usr/local/bin
sudo mv  include/*  /usr/local/include/
sudo chown  $(whoami) /usr/local/bin/protoc
```


2. 安装Go插件生成Go代码： 
```shell
//这个版本略高，proto中的go_package需要加. 或者 /,如果是低版本则不需要
go get github.com/golang/protobuf/protoc-gen-go

//安装低版本插件可以在任意一个Module中执行,这样就可以在$GOAPTH/bin/ 找到对应版本的执行文件了。
go  get github.com/golang/protobuf/protoc-gen-go@v1.1.0
```

## Example

在本地路径有一个test.proto文件，可以执行： 
```shell
protoc --proto_path=. --go_out=. test.proto
```
就可以生成pb.go文件，同时还可以与proto文件一个层级。如果protoc-gen-go 的版本超过v1.3.5的话，可能不能在一个层级了。