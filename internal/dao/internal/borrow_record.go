package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BorrowRecordDao 表 borrow_record。
type BorrowRecordDao struct {
	table    string
	group    string
	columns  BorrowRecordColumns
	handlers []gdb.ModelHandler
}

// BorrowRecordColumns 列名。
type BorrowRecordColumns struct {
	Id               string
	UserId           string
	BookId           string
	BorrowTime       string
	ExpectReturnTime string
	ReturnTime       string
	Status           string
}

var borrowRecordColumns = BorrowRecordColumns{
	Id:               "id",
	UserId:           "user_id",
	BookId:           "book_id",
	BorrowTime:       "borrow_time",
	ExpectReturnTime: "expect_return_time",
	ReturnTime:       "return_time",
	Status:           "status",
}

// NewBorrowRecordDao 创建 DAO。
func NewBorrowRecordDao(handlers ...gdb.ModelHandler) *BorrowRecordDao {
	return &BorrowRecordDao{
		group:    "default",
		table:    "borrow_record",
		columns:  borrowRecordColumns,
		handlers: handlers,
	}
}

func (dao *BorrowRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *BorrowRecordDao) Table() string {
	return dao.table
}

func (dao *BorrowRecordDao) Columns() BorrowRecordColumns {
	return dao.columns
}

func (dao *BorrowRecordDao) Group() string {
	return dao.group
}

func (dao *BorrowRecordDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

func (dao *BorrowRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
