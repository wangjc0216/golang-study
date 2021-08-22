# 常见问题

## json 报错
Json error calling MarshalJSON for type json.RawMessage


## slice问题

一个切片长度为10的切片，
```go

a := []int{0,1,2,3,4,5,6,7,8,9}
fmt.Println(a)

//fmt.Println(a[10])
//panic: runtime error: index out of range [10] with length 10

//fmt.Println(a[11])
//panic: runtime error: index out of range [11] with length 10

//fmt.Println(a[11:])
//panic: runtime error: slice bounds out of range [11:10]

fmt.Println(a[10:])
//[]

a = append(a[:9],a[10:]...)
fmt.Println(a)
```
## slice & string

相互转换问题

string是不可变的，一旦分配好就不再变了，底层实现与[]byte类似。
```go
type StringStruct type{
	str unsafe.Pointer
	len int 
}

type slice struct{
	array unsafe.Pointer
	len int 
	cap int 
}
```
分为string和[]byte标准转换 、与  string与[]byte高性能转换。
高性能转换可能会存在不安全的情况。如下例则会报错： 
```
func stringtoslicebytetmp(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	ret := reflect.SliceHeader{Data: str.Data, Len: str.Len, Cap: str.Len}
	return *(*[]byte)(unsafe.Pointer(&ret))
}

func main()  {
	str := "hello"
	by := stringtoslicebytetmp(str)
	by[0] = 'H'
}
```

