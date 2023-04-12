package Role

import (
	"encoding/json"
	"fmt"
	"myapp/app/Login/utils"
	onutils "myapp/app/Role/utils"

	"github.com/kataras/iris/v12"
)

func Index(ctx iris.Context) {
	// if ctx.Method() == "POST" {
	// 	ctx.JSON(utils.JsonResult{
	// 		Code: 0,
	// 		Msg:  "操作成功",
	// 	})
	// 	return
	// }
	ctx.View("role/index.html")
}

func List(ctx iris.Context) {
	var role = onutils.Role_mag{
		Id:       1,
		Username: "超级管理员",
		Status:   true,
		Sort:     1,
	}

	data, err := json.Marshal(role)
	if err != nil {
		fmt.Println("角色显示错误")
	}
	if ctx.Method() == "POST" {
		ctx.JSON(utils.JsonResult{
			Code:  0,
			Msg:   "操作成功",
			Data:  string(data),
			Count: 1,
		})
		return
	}
}
