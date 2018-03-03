package astidebug_test

import (
	"testing"

	"github.com/asticode/go-astitools/debug"
	"github.com/stretchr/testify/assert"
)

func mockStack() []byte {
	return []byte(`goroutine 1 [running]:
	runtime/astidebug.Stack(0xc42012e280, 0xc4200f8e17, 0xc4200f8d78)
	/usr/local/go/src/runtime/debug/stack.go:24 +0x79
	github.com/asticode/myrepo.glob..func2(0x0, 0x0, 0xc4200f8da8)
	/home/asticode/projects/go/src/github.com/asticode/myrepo/sync.go:55 +0x22
	github.com/asticode/myrepo.(*MyStruct).LogMessage(0xc420105e00, 0x7a2108, 0x11, 0x7d0710, 0xc4200f8e60)
	/home/asticode/projects/go/src/github.com/asticode/myrepo/sync.go:70 +0x3b
	github.com/asticode/myrepo.(*MyStruct).RUnlock.func1(0xc42001d980, 0xc4201dc000, 0x14, 0xc420105e00)
	/home/asticode/projects/go/src/github.com/asticode/myrepo/sync.go:147 +0x70
	github.com/asticode/myrepo.(*MyStruct).RUnlock(0xc420105e00, 0xc42001d980, 0xc4201dc000, 0x14)
	/home/asticode/projects/go/src/github.com/asticode/myrepo/sync.go:151 +0x3d
	main.(*Worker).Retire(0xc4200f1800, 0x90cfa0, 0xc420018d70)
	/home/asticode/projects/go/src/github.com/asticode/myproject/worker.go:174 +0x11d
	main.main()
	/home/asticode/projects/go/src/github.com/asticode/myproject/main.go:76 +0x571`)
}

func TestNewStack(t *testing.T) {
	astidebug.DebugStack = func() []byte {
		return mockStack()
	}
	var s = astidebug.NewStack()
	assert.Equal(t, astidebug.Stack{astidebug.StackItem{Filename: "/usr/local/go/src/runtime/debug/stack.go", Function: "runtime/astidebug.Stack", Line: 24}, astidebug.StackItem{Filename: "/home/asticode/projects/go/src/github.com/asticode/myrepo/sync.go", Function: "github.com/asticode/myrepo.glob..func2", Line: 55}, astidebug.StackItem{Filename: "/home/asticode/projects/go/src/github.com/asticode/myrepo/sync.go", Function: "github.com/asticode/myrepo.(*MyStruct).LogMessage", Line: 70}, astidebug.StackItem{Filename: "/home/asticode/projects/go/src/github.com/asticode/myrepo/sync.go", Function: "github.com/asticode/myrepo.(*MyStruct).RUnlock.func1", Line: 147}, astidebug.StackItem{Filename: "/home/asticode/projects/go/src/github.com/asticode/myrepo/sync.go", Function: "github.com/asticode/myrepo.(*MyStruct).RUnlock", Line: 151}, astidebug.StackItem{Filename: "/home/asticode/projects/go/src/github.com/asticode/myproject/worker.go", Function: "main.(*Worker).Retire", Line: 174}, astidebug.StackItem{Filename: "/home/asticode/projects/go/src/github.com/asticode/myproject/main.go", Function: "main.main", Line: 76}}, s)
}
