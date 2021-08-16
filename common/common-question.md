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
