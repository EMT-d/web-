package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User 表 user 的 DO。
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        any
	StudentNo any
	Username  any
	Password  any
}
