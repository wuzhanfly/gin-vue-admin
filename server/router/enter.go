package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/IndexPage"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Indexpage IndexPage.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
