# channel 

## channel 的使用场景

- [ ] 优雅关闭： 通过接受内核信号量进行判断，




## channel 发送指针遇到的问题记录

为什么chanel生产的指针中的值会有变化？    channel中传递的是指针，通过for向channel中塞数据的时候都指向了一个地址，所以影响了Rule的执行

### 例子1
```go
	fmt.Printf("%p\n",r)  
	fmt.Printf("%T\n",r)
```
分别输出： 
```shell
	0xc000384480
	*models.Rule
```
### 例子2
```go
	func ExampleLoc(){
	fori:=0;i<10;i++{
	fmt.Printf("%d,%p\n",i,&i)
	}
	//output:
	}
```
输出：
```go
	0,0xc000113e60
	1,0xc000113e60
	2,0xc000113e60
	3,0xc000113e60
	4,0xc000113e60
	5,0xc000113e60
	6,0xc000113e60
	7,0xc000113e60
	8,0xc000113e60
	9,0xc000113e60
```
	
### 例子3

```go
	func ExampleLoc(){
	varaint
	a=43
	fmt.Println(a)//43
	fmt.Println(&a)//0xc00000a0b8
	varb=&a
	fmt.Println(b)//0xc00000a0b8
	fmt.Printf("b:%p\n",b)
	fmt.Println(&b)//0xc000006030
	fmt.Printf("%T\n",b)//*int
	fmt.Println(*b)
	//output:
	}
```
	
结论1:  

%p 输出的是地址，

如果b是指针，那么fmt.printf("%p/n",b)则输出的是b指针所指向的地址；   fmt.printf("%p/n",&b)则表示的是b指针本身的地址；

如果b不是指针，是对象，那么fmt.printf("%p/n",b)会报错（如Printf format %p has arg a of wrong type int） fmt.printf("%p/n",&b)则表示的是b对象本身的地址；
	
结论2:
	
如果是for i:=0;i <10 ;i++{ fmt.printf("%p",i)} i的地址是固定的