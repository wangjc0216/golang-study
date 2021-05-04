
## Pointer
pointer 通常通过将不同包但是构造相同的结构体（自己的理解是字段数量、类型及顺序均相同）进行转换。
如果不使用unsafe.Pointer进行转换的话就会报错：
```shell
//panic: interface conversion: interface {} is *ecdsa.PublicKey, not *main.PublicKey
//goroutine 1 [running]:
//main.main()
//        (调用堆栈信息).../.../.../main.go:56 +0x55

func main() {

	pk := new(ecdsa.PublicKey)
	var i interface{}
	i = pk
	pkk := i.(*PublicKey)
	fmt.Printf("%T %v\n", pkk, pkk)
}

type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}
```
下面例子就是将ecdsa.PublicKey转换为本地的PublicKey
```shell
func main() {
	pk := new(ecdsa.PublicKey)
	pkk := (*PublicKey)(unsafe.Pointer(pk))
	fmt.Printf("%T %v\n", pkk, pkk)

}

type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}
```
