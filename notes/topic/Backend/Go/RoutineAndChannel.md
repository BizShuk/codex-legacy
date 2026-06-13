# Goroutine and Channel note

### usage

buffered channel

```go
ch := make(chan int, 100)
```

unbuffered channel

```go
ch: make(chan int)
```

Check ch whether it is closed

```go
v, ok := <- ch
```

Will for loop util ch gets closed

```go
for v := range ch {

}
```

defer close channels

```go
defer close(ch)
```

### select

Only block when channels are open util value arrives. nil channel will block forever.
closed channel will consumer CPU always.

```go
select {
    case <- ch1:
    case <- done:
        return
}

```

### time.After

it returns a channel.

```go
    select {
        case m:= <- c
        case <- time.After(10 * time.Second)
            return
    }
```

### merge

merge channels into one channel

### nil channel

close on nil channel will cause panic

### tips

- set as buffered channel to prevent publiser hanging.
- Only sender should close channels, only close when receiver must be told

### Pipeline Stages

stages close their outbound channels when all the send operations are done.
stages keep receiving values from inbound channels until those channels are closed.

### Error Handling

1. through channel

2.

### Fan-out, fan-in

Multiple functions can read from the same channel until that channel is closed; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that's closed when all the inputs are closed. This is called fan-in.

### Context

let parent can control goroutine cancelation instead of let goroutine make the cancelation.
It needs to generate context first and
