# GO TEST 实践

测试分为两种，分为功能测试和非功能测试，用于确保满足它的业务需求和性能标准。
功能测试覆盖了大部分功能，包括黑匣子测试、单元测试、集成测试、系统测试、回归测试、烟雾测试等测试类型。
非功能测试，也就是性能测试，包括基线测试和基准测试。测试更重于速度、稳定性、可扩展性、可靠性、负载容量以及负载情况的性能表现。

对于开发人员而言，需要关注的是单元测试。**单元测试，又称为程序员测试，是程序员们本该做的自我检查工作之一。**

Go语言的缔造者们从一开始就非常重视程序测试，并且为Go程序的开发者提供了丰富的API和工具。利用这些工具，我们可以


testing.M  

testing.T

testing.B



- [ ] 断言



- [ ] 单元测试
`gi `
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





##  Example Test

REF: https://blog.golang.org/examples

Example是可选编译的，作为test suite的一部分。



Example函数写在_test.go文件中，同时以Example作为开头。



```
package stringutil_test

import (
    "fmt"

    "github.com/golang/example/stringutil"
)

func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
    // Output: olleh
}
```

我们通过执行 `go test -v  -test.run ^ExampleReverse$` -不需要额外的参数就可以执行Example参数。

Example测试用例通过output来判断ExampleTest是否通过测试。

### Example 命名规则

```
func ExampleFoo()     // documents the Foo function or type
func ExampleBar_Qux() // documents the Qux method of type Bar
func Example()        // documents the package as a whole
```







