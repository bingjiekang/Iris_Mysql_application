package Time

import (
	"github.com/kataras/iris/v12"
)

func Time(ctx iris.Context) {
	ctx.View("time/index.html")
}

func Countdown(ctx iris.Context) {
	ctx.View("time/countdown.html")
}
