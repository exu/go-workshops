## Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

    ch <- v    // Send v to channel ch.
    v := <-ch  // Receive from ch, and
            // assign value to v.

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

    eventsChannel := make(chan int)

By default, **sends and receives block until the other side is ready**. This allows goroutines to synchronize without explicit locks or condition variables.

### Buffered Channels

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

    ch := make(chan int, 100)

Sends to a buffered channel **block only when the buffer is full**. Receives block when the buffer is empty.
