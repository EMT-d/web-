// =================================================================================
// 与 XAMPP 常见建表一致：book(id, book_name, author, stock)
// =================================================================================

package entity

// Book 对应表 book。
type Book struct {
	Id       uint64 `json:"id"       orm:"id"        `
	BookName string `json:"bookName" orm:"book_name" `
	Author   string `json:"author"   orm:"author"    `
	Stock    int    `json:"stock"    orm:"stock"     `
}
