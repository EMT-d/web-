// =================================================================================
// 与 XAMPP 常见建表一致：user(id, student_no, username, password)
// =================================================================================

package entity

// User 对应表 user。
type User struct {
	Id        uint64 `json:"id"        orm:"id"         `
	StudentNo string `json:"studentNo" orm:"student_no" `
	Username  string `json:"username"  orm:"username"   `
	Password  string `json:"password"  orm:"password"   `
}
