package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BorrowRecord 表 borrow_record 的 DO。
type BorrowRecord struct {
	g.Meta           `orm:"table:borrow_record, do:true"`
	Id               any
	UserId           any
	BookId           any
	BorrowTime       *gtime.Time
	ExpectReturnTime *gtime.Time
	ReturnTime       *gtime.Time
	Status           any
}
