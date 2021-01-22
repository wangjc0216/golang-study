 # SOCKET连接

## 疑问
- [ ]  如何使用Golang创建socket，像python3例子一样

## 问题引发的实践

在下班路上和磊总，赐哥讨论TCP server 最大连接数是由什么决定的问题想到的关联问题。

~~端口号？~~  

最大打开文件数？

**一致认为，应该是最大打开文件数决定的。**

好的，作为延伸，那么TCP client 最大连接数是由什么决定的呢？

刚开始在网上查找说是端口号决定的（后来实践证明，这个是错的）

**但是，有个疑问，为什么TCP Server 一个端口可以复用？但是TCP Client 端口就不能复用呢？**（后来实践得出，tcp client 连接的端口也是可以复用的）

该链接给了很大的启发：[TCP Server 为什么一个端口可以建立多个连接?](https://segmentfault.com/q/1010000003101541)

## 实践

首先使用简单的TCP server 和TCP client 来创建TCP请求，如`server.go`和`client.go`这两个文件。执行如下：

```
1. 启动服务端，服务端会暴露两个端口，分别为9999端口和6666端口
----
➜  socket go run server.go  
Server ready to read ...
A client connected :127.0.0.1:64889
127.0.0.1:64889 Say hello to Server... 
127.0.0.1:64889 write data to Server... 
127.0.0.1:64889 write data to Server... 
...
2.启动客户端,客户端<源ip,源端口号,目的ip,目的端口号>配置为<127.0.0.1,10789,127.0.0.1,9999>,执行代码，建立连接
----
➜  socket go run client.go 
127.0.0.1:64889 : Client connected!
ReadString
2021-01-21 23:51:50.738156 +0800 CST m=+10.375832398127.0.0.1:64889 Server Say hello! 

writing...
ReadString
2021-01-21 23:51:55.744068 +0800 CST m=+15.381938494127.0.0.1:64889 Server Say hello! 

writing...
ReadString
2021-01-21 23:52:00.750455 +0800 CST m=+20.388519360127.0.0.1:64889 Server Say hello! 
...
3.修改代码后（目的端口修改为6666），再次启动客户端，发现会报错
----
➜  socket go run client.go 
Client connect error ! dial tcp 127.0.0.1:10789->127.0.0.1:6666: bind: address already in use
```

可以发现，如果仅仅修改目的端口是不成成功创建套接字的，在`bind`的位置就给拦下来了。可以在python3执行程序：

```
>>> import socket
>>> s = socket.socket()
>>> s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEPORT, 1)
>>> s.bind(('127.0.0.1', 10789))
>>> s.connect(('127.0.0.1', 6666))
```

发现连接成功。

从客户端视角，可以在Mac上通过`lsof`命令来查看客户端10789端口的连接状态：

```
➜  ~ lsof -i :10789
COMMAND     PID   USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
com.docke   903 jcwang   72u  IPv6 0x1f811ee94e625af1      0t0  TCP localhost:6666->localhost:10789 (ESTABLISHED)
Python    75984 jcwang    3u  IPv4 0x1f811ee957fdc541      0t0  TCP localhost:10789->localhost:6666 (ESTABLISHED)
server    76082 jcwang    5u  IPv4 0x1f811ee959d4ab61      0t0  TCP localhost:distinct->localhost:10789 (ESTABLISHED)
client    76103 jcwang    3u  IPv4 0x1f811ee94b4206a1      0t0  TCP localhost:10789->localhost:distinct (ESTABLISHED)
```

10789端口是TCP Client的连接端口，发现连接了两个地址，分别是6666和distinct，distinct其实就是9999，因为进行地址反解。和iptables类似。客户端和服务端分别都会在TCP连接时创建一个socket。所以在上面可以看见四个连接。

再看下服务端视角看下9999端口的连接状态：

```
➜  ~ lsof -i :9999
COMMAND   PID   USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
server  76082 jcwang    3u  IPv4 0x1f811ee9498d02e1      0t0  TCP localhost:distinct (LISTEN)
server  76082 jcwang    5u  IPv4 0x1f811ee959d4ab61      0t0  TCP localhost:distinct->localhost:10789 (ESTABLISHED)
client  76103 jcwang    3u  IPv4 0x1f811ee94b4206a1      0t0  TCP localhost:10789->localhost:distinct (ESTABLISHED)
```

有三个连接，可以确定，一个是监听socket（listen状态的那个），还有两个是已连接socket(服务端和客户端各一个)。





