package Position

import (
	model "myapp/Model"
	Index_utils "myapp/app/Index/utils"
	utils "myapp/app/Login/utils"
	memberlevel_utils "myapp/app/MemberLevel/utils"
	"strconv"
	"strings"

	// "time"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

// 显示会员等级界面
func Index(ctx iris.Context) {
	ctx.View("memberlevel/index.html")
}

// 更新查询显示
func Memberlevel_append(ctx iris.Context) {

	id := ctx.Params().GetIntDefault("id", 0)
	if id > 0 {
		var req model.MemberLevel
		// 查询这个id对应的memberlevel的数据信息
		req = model.Select_memberlevel_id(id)
		// fmt.Println(req.Id, req.Username, req.Status)
		ctx.ViewData("info", req)
	}
	// fmt.Print("这是需要编辑的id:", id)
	ctx.View("memberlevel/edit.html")
}

// 更新memberlevel信息
func Update(ctx iris.Context) {
	var req memberlevel_utils.MemberLevelAddReq
	// 获取会员等级错误,显示错误信息
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 如果会员等级名已经存在,且id不同则不能添加
	if model.Select_memberlevel_exit(req.Name) && model.Select_memberlevel_id(req.Id).Name != req.Name {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "等级名已经存在,请重新添加",
		})
		return
	}

	// 全部正常后可以更新信息
	if !model.Update_memberlevel_mag(&req) {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "会员等级添加失败",
		})
		return
	}
	// 更新成功
	ctx.JSON(utils.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
	return

}

// 添加会员等级信息
func Add(ctx iris.Context) {

	var req memberlevel_utils.MemberLevelAddReq
	// 获取岗位等级信息错误,显示错误信息
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 如果添加的等级存在,返回已存在错误
	if model.Select_memberlevel_exit(req.Name) {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "等级已经存在,重新添加",
		})
		return
	}

	// 更新会员等级数据库
	if !model.Insert_memberlevel_mag(&req) {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "等级添加失败",
		})
		return
	}

	// 更新成功
	ctx.JSON(utils.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
	return

}

// 删除指定会员等级信息
func Delete_mag(ctx iris.Context) {
	var req Index_utils.UserUpadte
	// 获取用户名
	req.Nickname = sessions.Get(ctx).GetString("trickname")
	// 如果不是admin超级管理员不能删除信息
	if req.Nickname != "admin" {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "对不起,您没有权限进行此操作!",
		})
		return
	} else {
		// 获取待删除的id信息,string接收多个参数,(多个待删除id)
		// id := ctx.Params().GetIntDefault("id", 0)
		id := ctx.Params().GetString("id")
		// id为空不删除
		if id == "" {
			ctx.JSON(utils.JsonResult{
				Code: -1,
				Msg:  "记录ID不能为空",
			})
			return
		}
		// 将字符串转为字符串列表
		st := strings.Split(id, ",")
		for _, v := range st {
			// 将每个字符串转为数字id
			tm_id, _ := strconv.ParseInt(v, 10, 64)
			if tm_id > 0 {
				// 删除id对应等级的信息
				if !model.Delete_memberlevel_mag(int(tm_id)) {
					ctx.JSON(utils.JsonResult{
						Code: -1,
						Msg:  "等级删除失败",
					})
					return
				}
			}
		}
		// 全部删除后返回删除成功
		ctx.JSON(utils.JsonResult{
			Code: 0,
			Msg:  "删除成功",
		})
		return

	}

}

// 显示全部岗位信息
func List(ctx iris.Context) {
	// 定义一个传输给前端data的列表
	var role_data []interface{}
	// 获取level数据表的全部数据(是一个列表的加入的多个结构体)
	var role_mag = model.Select_memberlevel()
	// 遍历这个列表里的结构体,将对应数据加入到 返回给前端的数据data里
	for _, v := range role_mag {
		role_data = append(role_data, map[string]interface{}{
			"id":          v.Id,
			"name":        v.Name,
			"sort":        v.Sort,
			"create_time": v.Create_time,
			"update_time": v.Update_time,
		})
	}

	// role_data = append(role_data, map[string]interface{}{
	// 	"id":       1,
	// 	"username": "超级管理员",
	// 	"status":   true,
	// 	"sort":     5,
	// 	"created":  time.Now().Year(),
	// 	"updated":  time.Now().Month(),
	// })

	if ctx.Method() == "POST" {

		ctx.JSON(utils.JsonResult{
			Code:  0,
			Msg:   "调用成功",
			Data:  role_data,
			Count: 10,
		})
		return
	}
}
