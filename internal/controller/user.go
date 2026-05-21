package controller

import (
	"gf-demo-user-master/internal/dao"
	"gf-demo-user-master/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/bcrypt"
)

var User = &userController{}

type userController struct{}

func userPublic(u *entity.User) g.Map {
	if u == nil {
		return nil
	}
	return g.Map{
		"id":        u.Id,
		"studentNo": u.StudentNo,
		"username":  u.Username,
	}
}

// Register 用户注册（增）
func (c *userController) Register(r *ghttp.Request) {
	var in struct {
		StudentNo string `json:"studentNo"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}
	if err := r.Parse(&in); err != nil || in.StudentNo == "" || in.Username == "" || in.Password == "" {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误：学号、姓名、密码必填"})
	}
	n, err := dao.User.Ctx(r.Context()).Where("student_no", in.StudentNo).Count()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "注册失败"})
	}
	if n > 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "学号已存在"})
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "注册失败"})
	}
	_, err = dao.User.Ctx(r.Context()).Data(entity.User{
		StudentNo: in.StudentNo,
		Username:  in.Username,
		Password:  string(hashed),
	}).Insert()
	if err != nil {
		g.Log().Error(r.Context(), err)
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "注册失败：若提示 Data too long，请在库中执行 manifest/sql/upgrade_user_password.sql"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "注册成功"})
}

// Login 用户登录（查 + 校验）
func (c *userController) Login(r *ghttp.Request) {
	var in struct {
		StudentNo string `json:"studentNo"`
		Password  string `json:"password"`
	}
	if err := r.Parse(&in); err != nil || in.StudentNo == "" || in.Password == "" {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误"})
	}
	row, err := dao.User.Ctx(r.Context()).Where("student_no", in.StudentNo).One()
	if err != nil || row.IsEmpty() {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "用户不存在或密码错误"})
	}
	var u entity.User
	if err := row.Struct(&u); err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "登录失败"})
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password)) != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "用户不存在或密码错误"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "登录成功", "data": userPublic(&u)})
}

// List 用户列表（查）
func (c *userController) List(r *ghttp.Request) {
	all, err := dao.User.Ctx(r.Context()).All()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "查询失败"})
	}
	var list []entity.User
	if err := all.Structs(&list); err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "查询失败"})
	}
	out := make([]g.Map, 0, len(list))
	for i := range list {
		out = append(out, userPublic(&list[i]))
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": out})
}

// Update 更新用户（改），仅 username / password（表无 phone、时间戳）
func (c *userController) Update(r *ghttp.Request) {
	var in struct {
		Id       uint64 `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := r.Parse(&in); err != nil || in.Id == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误"})
	}
	data := g.Map{}
	if in.Username != "" {
		data["username"] = in.Username
	}
	if in.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "更新失败"})
		}
		data["password"] = string(hashed)
	}
	if len(data) == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "请提供 username 或 password"})
	}
	_, err := dao.User.Ctx(r.Context()).Where("id", in.Id).Data(data).Update()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "更新失败"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "更新成功"})
}

// Delete 删除用户（删）
func (c *userController) Delete(r *ghttp.Request) {
	id := gconv.Uint64(r.GetQuery("id"))
	if id == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误：需要 id"})
	}
	_, err := dao.User.Ctx(r.Context()).Where("id", id).Delete()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "删除失败"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "删除成功"})
}
