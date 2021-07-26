# Sync.Map解析


## atomic.Value

该结构主要暴露两个接口，Load()与Store(),是原子性操作且并发安全，常用于读多写少的情况。比如定期去更新配置之类的。

## unsafe.Pointer
这里使用一个随机的指针来表示被删除。使用atomic包来保证unsafe.Pointer改动的原子性。 使用atomic默认就认为是atomic是读多写少的，是一种乐观锁。
使用如下函数：
```shell
atomic.LoadPointer
atomic.CompareAndSwapPointer
```
 

## Sync.Map



```shell
type Map struct {
 mu Mutex
 read atomic.Value // readOnly
 dirty map[interface{}]*entry
 misses int
}

type readOnly struct {
 m       map[interface{}]*entry
 amended bool
}

```