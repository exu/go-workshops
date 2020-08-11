package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/hello/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})
	app.Listen(":8080")
}
