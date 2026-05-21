// =================================================================================
// 与 XAMPP 常见建表一致：borrow_record(..., return_time, status)
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BorrowRecord 对应表 borrow_record。
type BorrowRecord struct {
	Id               uint64      `json:"id"               orm:"id"                 `
	UserId           uint64      `json:"userId"           orm:"user_id"            `
	BookId           uint64      `json:"bookId"           orm:"book_id"            `
	BorrowTime       *gtime.Time `json:"borrowTime"       orm:"borrow_time"        `
	ExpectReturnTime *gtime.Time `json:"expectReturnTime" orm:"expect_return_time" `
	ReturnTime       *gtime.Time `json:"returnTime"       orm:"return_time"        `
	Status           uint        `json:"status"           orm:"status"             `
}
