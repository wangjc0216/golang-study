# 工具整理

## Goland



**配置breadcrumbs**  ：可以配置出在顶部或者底部看到当前行所在的函数，如果函数本身很长，这个功能很有用处。

**配置bookmarks**： 配置书签。在阅读代码的时候，可以对具体的行进行标记数字/字母，同时可添加描述。




## dlv

dlv作为调试工具。相比于gdb，dlv更加智能易用。当前只有配合Goland远程联调的使用经验。

```shell
调试执行文件
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./${executable file} --  ${PARAM}
也可以将二进制文件所在的目录导入到PATH中，如在终端配置
export  PATH=$PATH:${PWD}
然后执行
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ${executable file} --  ${PARAM}

调试测试文件
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient test -- test.run ^TestDemo1$
```

reference: [使用GoLand进行调试](https://www.jianshu.com/p/ce7e96527a4a)