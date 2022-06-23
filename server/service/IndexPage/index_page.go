package IndexPage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/IndexPage"
	IndexPageReq "github.com/flipped-aurora/gin-vue-admin/server/model/IndexPage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IndexPageService struct {
}

// CreateIndexPage 创建IndexPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (InPageService *IndexPageService) CreateIndexPage(InPage IndexPage.IndexPage) (err error) {
	err = global.GVA_DB.Create(&InPage).Error
	return err
}

// DeleteIndexPage 删除IndexPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (InPageService *IndexPageService) DeleteIndexPage(InPage IndexPage.IndexPage) (err error) {
	err = global.GVA_DB.Delete(&InPage).Error
	return err
}

// DeleteIndexPageByIds 批量删除IndexPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (InPageService *IndexPageService) DeleteIndexPageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]IndexPage.IndexPage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIndexPage 更新IndexPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (InPageService *IndexPageService) UpdateIndexPage(InPage IndexPage.IndexPage) (err error) {
	err = global.GVA_DB.Save(&InPage).Error
	return err
}

// GetIndexPage 根据id获取IndexPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (InPageService *IndexPageService) GetIndexPage(id uint) (InPage IndexPage.IndexPage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&InPage).Error
	return
}

// GetIndexPageInfoList 分页获取IndexPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (InPageService *IndexPageService) GetIndexPageInfoList(info IndexPageReq.IndexPageSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&IndexPage.IndexPage{})
	var InPages []IndexPage.IndexPage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&InPages).Error
	return InPages, total, err
}
