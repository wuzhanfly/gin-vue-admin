package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/IndexPage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IndexPageSearch struct {
	IndexPage.IndexPage
	request.PageInfo
}
