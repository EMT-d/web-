package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BookDao 表 book。
type BookDao struct {
	table    string
	group    string
	columns  BookColumns
	handlers []gdb.ModelHandler
}

// BookColumns 列名。
type BookColumns struct {
	Id       string
	BookName string
	Author   string
	Stock    string
}

var bookColumns = BookColumns{
	Id:       "id",
	BookName: "book_name",
	Author:   "author",
	Stock:    "stock",
}

// NewBookDao 创建 DAO。
func NewBookDao(handlers ...gdb.ModelHandler) *BookDao {
	return &BookDao{
		group:    "default",
		table:    "book",
		columns:  bookColumns,
		handlers: handlers,
	}
}

func (dao *BookDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *BookDao) Table() string {
	return dao.table
}

func (dao *BookDao) Columns() BookColumns {
	return dao.columns
}

func (dao *BookDao) Group() string {
	return dao.group
}

func (dao *BookDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

func (dao *BookDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
