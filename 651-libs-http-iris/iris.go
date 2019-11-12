package main

import 	"github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/hello/{name:string}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
        	ctx.Writef("Hello %s", name)
	})
	app.Run(iris.Addr(":8080"))
}
