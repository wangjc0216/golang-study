# 问题清单

- [x] golang 如何比较[]byte 是否相同
```shell
可以参考 https://www.cnblogs.com/apocelipes/p/11116725.html
[]byte比较可以通过 bytes.Equal函数来验证；如果是其他对象则可以通过reflect.DeepEqual来验证，对reflect.DeepEqual的描述如下：

DeepEqual reports whether x and y are “deeply equal,” defined as follows. Two values of identical type are deeply equal if one of the following cases applies. Values of distinct types are never deeply equal.
...
Slice values are deeply equal when all of the following are true: they are both nil or both non-nil, they have the same length, and either they point to the same initial entry of the same underlying array (that is, &x[0] == &y[0]) or their corresponding elements (up to length) are deeply equal. Note that a non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil)) are not deeply equal.
```  

- [ ] base64 为什么 编写方式 处于什么样的用途，为什么不直接提供一个base64 的函数呢？ 而是要初始化一个对象

- [ ] go的条件编译

- [ ] []byte(...)与（...）.([]byte)有什么区别？


- [x] 发现服务总是 exit 1 退出？ 
```
后来发现是因为使用logger.Fatalf()函数来打印日志，Fatalf通常在打印之后就会将程序退出，所以除了单元测试不要使用Fatalf()函数
```