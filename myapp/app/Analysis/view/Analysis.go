package Analysis

import (
	"github.com/kataras/iris/v12"
)

func Analysis(ctx iris.Context) {
	ctx.View("analysis/index.html")
}
