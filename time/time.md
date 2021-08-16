# time包

## time.AfterFunc

time.AfterFunc 会另开一个goroutine，等到具体时间会执行func()内容。

```go
func ExampleTestTime() {
	fmt.Println("before..", time.Now())
	//var t *time.Timer

	time.AfterFunc(time.Second*5, func() {
		fmt.Println("func..", time.Now())
	})

	fmt.Println("after1..", time.Now())
	time.Sleep(time.Second)
	//<-t.C
	fmt.Println("after2..", time.Now())
	time.Sleep(7 * time.Second)
	//output:
}
```


## time.Duration(next)*time.Minute
```
var next int
next * time.Minute 会报错
time.Duration(next) * time.Minute 才可以
```

## time.Ticker time.Timer
time.Ticker 是定时的,会按照指定时间发出Ticker，比如下面例子会按照3s打印日志

time.Timer每次都需要Reset

```go
func ExampleTimerTicker() {
ticker := time.NewTicker(2 * time.Second)

for i := 0; i < 5; i++ {
<-ticker.C
fmt.Println("ticker ...", time.Now())
time.Sleep(3 * time.Second)
}
timer := time.NewTimer(2 * time.Second)
for i := 0; i < 5; i++ {
<-timer.C
fmt.Println("timer ...", time.Now())
timer.Reset(time.Second)
}
//output:
}
```