package main

import (
	"gf-demo-user-master/internal/router"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	s := g.Server()
	router.Register(s)
	s.Run()
}
