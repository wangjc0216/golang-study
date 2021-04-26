
## Pointer
pointer 通常通过将不同包但是构造相同的结构体（自己的理解是字段数量、类型及顺序均相同）进行转换。
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