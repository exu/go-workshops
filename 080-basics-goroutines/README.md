## Go routines

A goroutine is a lightweight thread of execution.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives (channels).

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
