package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao 表 user。
type UserDao struct {
	table    string
	group    string
	columns  UserColumns
	handlers []gdb.ModelHandler
}

// UserColumns 列名。
type UserColumns struct {
	Id        string
	StudentNo string
	Username  string
	Password  string
}

var userColumns = UserColumns{
	Id:        "id",
	StudentNo: "student_no",
	Username:  "username",
	Password:  "password",
}

// NewUserDao 创建 DAO。
func NewUserDao(handlers ...gdb.ModelHandler) *UserDao {
	return &UserDao{
		group:    "default",
		table:    "user",
		columns:  userColumns,
		handlers: handlers,
	}
}

func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *UserDao) Table() string {
	return dao.table
}

func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

func (dao *UserDao) Group() string {
	return dao.group
}

func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
