package City

import (
	// "encoding/json"
	// "myapp/app/Login/utils"

	"github.com/kataras/iris/v12"
)

func City(ctx iris.Context) {
	ctx.View("city/city.html")
}
