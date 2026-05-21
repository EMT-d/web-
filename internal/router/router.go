package router

import (
	"gf-demo-user-master/internal/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(s *ghttp.Server) {
	api := s.Group("/api")

	// 用户 CRUD + 注册/登录（抓包实验）
	api.POST("/user/register", controller.User.Register)
	api.POST("/user/login", controller.User.Login)
	api.GET("/user/list", controller.User.List)
	api.PUT("/user/update", controller.User.Update)
	api.DELETE("/user/delete", controller.User.Delete)

	// 图书 CRUD
	api.POST("/book/add", controller.Book.Add)
	api.GET("/book/list", controller.Book.List)
	api.PUT("/book/update", controller.Book.Update)
	api.DELETE("/book/delete", controller.Book.Delete)

	// 借阅表 borrow CRUD + 归还
	api.POST("/borrow/add", controller.Borrow.Add)
	api.GET("/borrow/list", controller.Borrow.List)
	api.GET("/borrow/detail", controller.Borrow.Detail)
	api.PUT("/borrow/update", controller.Borrow.Update)
	api.POST("/borrow/return", controller.Borrow.Return)
	api.DELETE("/borrow/delete", controller.Borrow.Delete)
}
