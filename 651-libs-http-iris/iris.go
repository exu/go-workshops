package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)
func main() {
	app := iris.New()
	app.Get("/hello/{name:string}", func(ctx context.Context) {
		name := ctx.Params().Get("name")
        	ctx.Writef("Hello %s", name)
	})
	app.Run(iris.Addr(":8080"))
}
