package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Book 表 book 的 DO。
type Book struct {
	g.Meta   `orm:"table:book, do:true"`
	Id       any
	BookName any
	Author   any
	Stock    any
}
