package User

import (
	// "fmt"
	"fmt"
	model "myapp/Model"
	Index_utils "myapp/app/Index/utils"
	utils "myapp/app/Login/utils"

	// role_utils "myapp/app/Role/utils"
	user_utils "myapp/app/User/utils"

	// user_utils "myapp/app/User/utils"
	"strconv"
	"strings"

	// "time"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func Index(ctx iris.Context) {
	ctx.View("user/index.html")
}

// 更新查询显示
func User_append(ctx iris.Context) {

	id := ctx.Params().GetIntDefault("id", 0)
	var req user_utils.User_mag
	// id大于0 显示数据
	if id > 0 {
		//
		var reqs model.Users
		// 查询这个id对应的user的数据信息
		reqs = model.Select_user_id(id)
		// fmt.Println(req)
		// fmt.Println("id大于0")
		// ctx.ViewData("info", req)
		// 用来奖数据一起显示
		req.Id = reqs.Id
		req.Realname = reqs.Realname
		req.Gender = reqs.Gender
		req.Nickname = reqs.Nickname
		req.Password = reqs.Password
		req.Status = reqs.Status
		req.LevelName = reqs.LevelName
		req.PositionName = reqs.PositionName
		req.RoleName = reqs.RoleName
		req.Mobile = reqs.Mobile
		req.Email = reqs.Email
		req.Address = reqs.Address
		req.Sort = reqs.Sort
		req.Note = reqs.Note

	}
	// 将职称/等级/角色的信息从数据库查询到并传给前端
	req.PositionName = model.Select_all_position()
	req.LevelName = model.Select_all_level()
	req.RoleName = model.Select_all_role()

	// 从三个数据库查询的信息返回
	ctx.ViewData("info", req)
	ctx.View("user/edit.html")
}

// 更新用户信息
func Update(ctx iris.Context) {
	var req user_utils.UserAddReq
	// 获取用户信息错误,显示错误信息
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 如果用户昵称名已经存在,且id不同则不能添加
	if model.Select_role_exit(req.Nickname) && model.Select_user_id(req.Id).Nickname != req.Nickname {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "用户已经存在,重新添加",
		})
		return
	}

	// 全部正常后可以更新信息
	if !model.Update_user_mag(&req) {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "用户添加失败",
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

// 添加用户信息
func Add(ctx iris.Context) {

	var req user_utils.UserAddReq
	// 获取用户信息错误,显示错误信息
	if err := ctx.ReadForm(&req); err != nil {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		fmt.Println("获取用户信息错误")
		return
	}

	// 如果添加的用户存在,返回已存在错误
	if model.Select_user_exit(req.Nickname) {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "用户已经存在,重新添加",
		})
		return
	}

	// 更新用户数据库
	if !model.Insert_user_mag(&req) {
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  "用户添加失败",
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

// 删除指定角色信息
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
				// 删除id对应用户的信息
				if !model.Delete_user_mag(int(tm_id)) {
					ctx.JSON(utils.JsonResult{
						Code: -1,
						Msg:  "用户删除失败",
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

// 显示全部用户信息
func List(ctx iris.Context) {
	var req user_utils.UserPageReq
	// 定义一个传输给前端data的列表
	var role_data []interface{}
	if err := ctx.ReadForm(&req); err != nil {
		// 返回错误信息
		ctx.JSON(utils.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	if req.Name == "" && req.Gender == 0 {
		// 获取Role_mag数据表的全部数据(是一个列表的加入的多个结构体)
		var role_mag = model.Select_user()
		// 遍历这个列表里的结构体,将对应数据加入到 返回给前端的数据data里
		for _, v := range role_mag {
			role_data = append(role_data, map[string]interface{}{
				"id":           v.Id,
				"realname":     v.Realname,
				"nickname":     v.Nickname,
				"gender":       v.Gender,
				"status":       v.Status,
				"levelName":    v.LevelName,
				"positionName": v.PositionName,
				"mobile":       v.Mobile,
				"email":        v.Email,
				"sort":         v.Sort,
				"create_time":  v.Create_time,
				"update_time":  v.Update_time,
			})
		}
	} else {
		// fmt.Println("gender=", req.Gender)
		// 获取Role_mag数据表的部分数据(是一个列表的加入的多个结构体)
		var role_mag = model.Select_user_limit(req.Name, req.Gender)
		// 遍历这个列表里的结构体,将对应数据加入到 返回给前端的数据data里
		for _, v := range role_mag {
			role_data = append(role_data, map[string]interface{}{
				"id":           v.Id,
				"realname":     v.Realname,
				"nickname":     v.Nickname,
				"gender":       v.Gender,
				"status":       v.Status,
				"levelName":    v.LevelName,
				"positionName": v.PositionName,
				"mobile":       v.Mobile,
				"email":        v.Email,
				"sort":         v.Sort,
				"create_time":  v.Create_time,
				"update_time":  v.Update_time,
			})
		}
	}

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
