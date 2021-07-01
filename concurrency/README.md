# 并发控制

## key princaple

- goroutine生命周期管理。可以通过channel、context等手段控制goroutine的生命周期何时开始何时结束。
- 优雅退出。1）如果并发量不大，可以通过类似http总shutdown的方式来进行退出；2）如果并发量大，可以通过一些worker来执行。


### 生命周期管理

1） 通过channel。当关闭channel的时候，相应的chan将不再阻塞。
```go
func DoSomething(){
    s := http.Server{
        Handler: handler,
        Addr:    addr,
    }

    go func() {
    		<-stop
	    	log.Printf("server will exiting, addr: %s", addr)
		    s.Shutdown(context.Background())
    }()
    return s.ListenAndServe()
}
```

2） 通过context。


### 优雅退出

在http中，每个Handler可以通过goroutine来做一些后台逻辑，并及时返回用户结果。 下面示例通过goroutine来做一些后台的逻辑操作，从而达到快速返回用户结果的逻辑。
```go
// Tracker knows how to track events for the application.
type Tracker struct{}
 
// Event records an event to a database or stream.
func (t *Tracker) Event(data string) {
    time.Sleep(time.Millisecond) // Simulate network write latency.
    log.Println(data)
}

type App struct {
    track Tracker
}

// Handle represents an example handler for the web service.
func (a *App) Handle(w http.ResponseWriter, r *http.Request) {

    // Do some actual work.
    // Respond to the client.
    w.WriteHeader(http.StatusCreated)

    // Fire and Hope.
    // BUG: We are not managing this goroutine.
    go a.track.Event("this event")
}
```
上述代码存在问题，当server 关闭的时候，可能会有一些Event还没有执行完。我们等待Event执行完成后再退出（优雅退出）。我们对代码进行重构，重构结果如下： 
```go
// Tracker knows how to track events for the application
type Tracker struct{
	wg sync.WaitGroup
}

// Event starts tracking an event. It runs asynchronously to
// not block the caller. Be sure to call the Shutdown function
// before the program exits so all tracked events finish.
func (t *Tracker) Event(data string) {
    // Increment counter so Shutdown knows to wait for this event.
    t.wg.Add(1)
    // Track event in a goroutine so caller is not blocked.
    go func() {
        // Decrement counter to tell Shutdown this goroutine finished.
        defer t.wg.Done()
        time.Sleep(time.Millisecond) // Simulate network write latency.
        log.Println(data)
    }()
}

func (t *Tacker)Shutdown(){
	t.wg.Wait()
}


func main(){
	//start a server
	
	//shutdown gouroutine ,wait fro all event goroutines to finish
	a.tracker.Shutdown()
}
```
但是如果goroutine因为处理时间过长，可能会导致主程序迟迟不能退出，所以我们需要对goroutine做一个时间限制，如果超出对应时间，则强制退出。
```go
// Shutdown waits for all tracked events to finish processing
// or for the provided context to be canceled.
func (t *Tracker) Shutdown(ctx context.Context) error {

    // Create a channel to signal when the waitgroup is finished.
    ch := make(chan struct{})

    // Create a goroutine to wait for all other goroutines to
    // be done then close the channel to unblock the select.
    go func() {
        t.wg.Wait()
        close(ch)
    }()

    // Block this function from returning. Wait for either the
    // waitgroup to finish or the context to expire.
    select {
    case <-ch:
        return nil
    case <-ctx.Done():
        return errors.New("timeout")
    }
}

func main(){
	
	...
	
	// Wait up to 5 seconds for all event goroutines to finish.
	const timeout = 5 * time.Second
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    err := a.track.Shutdown(ctx)
}
```



## Reference

[Concurrency Trap #2: Incomplete Work](https://www.ardanlabs.com/blog/2019/04/concurrency-trap-2-incomplete-work.html)

[Go并发编程(一)goroutine](https://lailin.xyz/post/go-training-week3-goroutine.html)
