# GO TEST 实践

测试分为两种，分为功能测试和非功能测试，用于确保满足它的业务需求和性能标准。
功能测试覆盖了大部分功能，包括黑匣子测试、单元测试、集成测试、系统测试、回归测试、烟雾测试等测试类型。
非功能测试，也就是性能测试，包括基线测试和基准测试。测试更重于速度、稳定性、可扩展性、可靠性、负载容量以及负载情况的性能表现。


- [ ] 单元测试

```shell
go test
```



- [ ] 基准测试

基准测试（benchmarking）是一种测量和评估软件性能指标的活动。你可以在某个时候通过基准测试建立一个已知的性能水平（称为基准线），
当系统的软硬件环境发生变化之后再进行一次基准测试以确定那些变化对性能的影响。 这是基准测试最常见的用途。
其他用途包括测定某种负载水平下的性能极限、管理系统或环境的变化、发现可能导致性能问题的条件，等等。

```shell
go test -bench=.
```