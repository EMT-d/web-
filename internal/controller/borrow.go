package controller

import (
	"time"

	"gf-demo-user-master/internal/dao"
	"gf-demo-user-master/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var Borrow = &borrowController{}

type borrowController struct{}

type borrowAddReq struct {
	StudentNo string `json:"studentNo"`
	BookId    uint64 `json:"bookId"`
}

// Add 借阅（增）
func (c *borrowController) Add(r *ghttp.Request) {
	var req borrowAddReq
	if err := r.Parse(&req); err != nil || req.StudentNo == "" || req.BookId == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误"})
	}
	user, err := dao.User.Ctx(r.Context()).Where("student_no", req.StudentNo).One()
	if err != nil || user.IsEmpty() {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "用户不存在"})
	}
	var u entity.User
	if err := user.Struct(&u); err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "用户不存在"})
	}
	now := gtime.Now()
	expect := gtime.New(time.Now().AddDate(0, 0, 7))
	record := entity.BorrowRecord{
		UserId:           u.Id,
		BookId:           req.BookId,
		BorrowTime:       now,
		ExpectReturnTime: expect,
		Status:           1,
	}
	_, err = dao.BorrowRecord.Ctx(r.Context()).Data(record).Insert()
	if err != nil {
		g.Log().Error(r.Context(), err)
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "借阅失败"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "借阅成功"})
}

// List 借阅列表（查）
func (c *borrowController) List(r *ghttp.Request) {
	list, err := dao.BorrowRecord.Ctx(r.Context()).All()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "查询失败"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": list})
}

// Detail 单条借阅（查）
func (c *borrowController) Detail(r *ghttp.Request) {
	id := gconv.Uint64(r.GetQuery("id"))
	if id == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误：需要 id"})
	}
	one, err := dao.BorrowRecord.Ctx(r.Context()).Where("id", id).One()
	if err != nil || one.IsEmpty() {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "记录不存在"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "data": one})
}

// Update 更新借阅记录（改）
func (c *borrowController) Update(r *ghttp.Request) {
	var in struct {
		Id               uint64      `json:"id"`
		UserId           uint64      `json:"userId"`
		BookId           uint64      `json:"bookId"`
		Status           *uint       `json:"status"`
		ExpectReturnTime *gtime.Time `json:"expectReturnTime"`
		ReturnTime       *gtime.Time `json:"returnTime"`
	}
	if err := r.Parse(&in); err != nil || in.Id == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误：需要 id"})
	}
	data := g.Map{}
	if in.UserId > 0 {
		data["user_id"] = in.UserId
	}
	if in.BookId > 0 {
		data["book_id"] = in.BookId
	}
	if in.Status != nil {
		data["status"] = *in.Status
	}
	if in.ExpectReturnTime != nil {
		data["expect_return_time"] = in.ExpectReturnTime
	}
	if in.ReturnTime != nil {
		data["return_time"] = in.ReturnTime
	}
	if len(data) == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "请提供至少一项要修改的字段"})
	}
	_, err := dao.BorrowRecord.Ctx(r.Context()).Where("id", in.Id).Data(data).Update()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "更新失败"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "更新成功"})
}

// Return 归还（改状态为已归还，写入 return_time）
func (c *borrowController) Return(r *ghttp.Request) {
	id := gconv.Uint64(r.GetQuery("id"))
	if id == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误：需要 id"})
	}
	_, err := dao.BorrowRecord.Ctx(r.Context()).Where("id", id).Data(g.Map{
		"status":      2,
		"return_time": gtime.Now(),
	}).Update()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "归还失败"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "归还成功"})
}

// Delete 删除借阅记录（删）
func (c *borrowController) Delete(r *ghttp.Request) {
	id := gconv.Uint64(r.GetQuery("id"))
	if id == 0 {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "参数错误：需要 id"})
	}
	_, err := dao.BorrowRecord.Ctx(r.Context()).Where("id", id).Delete()
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"code": 1, "msg": "删除失败"})
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "删除成功"})
}
