## Interface'y

Go have "implicit interfaces". To implement an interface in Go, we
just need to implement all the methods in the interface.

So what is an interface? An interface is two things: it is a set of
methods, but it is also a type.

### The `interface{}` type

The `interface{}` type, the empty interface, is the source of much
confusion. The interface{} type is the interface that has no
methods. Since there is no implements keyword, all types implement at
least zero methods, and satisfying an interface is done automatically,
all types satisfy the empty interface. That means that if you write a
function that takes an `interface{}` value as a parameter, you can
supply that function with any value. So this function:

    func DoSomething(v interface{}) {
    // ...
    }

will accept any parameter whatsoever.

### Type assertions

- https://medium.com/golangspec/type-assertions-in-go-e609759c42e1
