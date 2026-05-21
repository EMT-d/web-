package controller

import (
	"gf-demo-user-master/internal/dao"
	"gf-demo-user-master/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var Book = &bookController{}

type bookController struct{}

// Add 添加图书（增）
func (b *bookController) Add(r *ghttp.Request) {
	var book entity.Book
	if err := r.Parse(&book); err != nil {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 1, "msg": "参数错误"})
	}
	if book.BookName == "" {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 1, "msg": "书名必填"})
	}
	_, err := dao.Book.Ctx(r.Context()).Insert(g.Map{
		"book_name": book.BookName,
		"author":    book.Author,
		"stock":     book.Stock,
	})
	if err != nil {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 1, "msg": "添加失败"})
	}
	r.Response.WriteJson(map[string]interface{}{"code": 0, "msg": "添加成功"})
}

// List 图书列表（查）
func (b *bookController) List(r *ghttp.Request) {
	list, err := dao.Book.Ctx(r.Context()).All()
	if err != nil {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 1, "msg": "查询失败"})
	}
	r.Response.WriteJson(map[string]interface{}{"code": 0, "data": list})
}

// Update 修改图书（改）
func (b *bookController) Update(r *ghttp.Request) {
	var book entity.Book
	if err := r.Parse(&book); err != nil || book.Id == 0 {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 1, "msg": "参数错误：需要 id"})
	}
	_, err := dao.Book.Ctx(r.Context()).Where("id", book.Id).Data(g.Map{
		"book_name": book.BookName,
		"author":    book.Author,
		"stock":     book.Stock,
	}).Update()
	if err != nil {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 1, "msg": "修改失败"})
	}
	r.Response.WriteJson(map[string]interface{}{"code": 0, "msg": "修改成功"})
}

// Delete 删除图书（删）
func (b *bookController) Delete(r *ghttp.Request) {
	id := gconv.Uint64(r.GetQuery("id"))
	_, err := dao.Book.Ctx(r.Context()).Where("id", id).Delete()
	if err != nil {
		r.Response.WriteJsonExit(map[string]interface{}{"code": 1, "msg": "删除失败"})
	}
	r.Response.WriteJson(map[string]interface{}{"code": 0, "msg": "删除成功"})
}
