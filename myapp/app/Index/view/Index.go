package Index

import (
	"github.com/kataras/iris/v12"
)

func Index(ctx iris.Context) {
	ctx.View("index.html")
}

func Default(ctx iris.Context) {
	ctx.View("welcome.html")
}
