package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"time"
)

type JunquanService struct {
}

// CreateJunquan 创建Junquan记录
// Author [piexlmax](https://github.com/piexlmax)
func (JqService *JunquanService) CreateJunquan(Jq autocode.Junquan) (err error) {
	err = global.GVA_DB.Create(&Jq).Error
	return err
}

// DeleteJunquan 删除Junquan记录
// Author [piexlmax](https://github.com/piexlmax)
func (JqService *JunquanService) DeleteJunquan(Jq autocode.Junquan) (err error) {
	err = global.GVA_DB.Delete(&Jq).Error
	return err
}

// DeleteJunquanByIds 批量删除Junquan记录
// Author [piexlmax](https://github.com/piexlmax)
func (JqService *JunquanService) DeleteJunquanByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Junquan{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateJunquan 更新Junquan记录
// Author [piexlmax](https://github.com/piexlmax)
func (JqService *JunquanService) UpdateJunquan(Jq autocode.Junquan) (err error) {
	err = global.GVA_DB.Save(&Jq).Error
	return err
}

// GetJunquan 根据id获取Junquan记录
// Author [piexlmax](https://github.com/piexlmax)
func (JqService *JunquanService) GetJunquan(id uint) (err error, Jq autocode.Junquan) {
	err = global.GVA_DB.Where("id = ?", id).First(&Jq).Error
	return
}

// GetJunquanInfoList 分页获取Junquan记录
// Author [piexlmax](https://github.com/piexlmax)
func (JqService *JunquanService) GetJunquanInfoList(info autoCodeReq.JunquanSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Junquan{})
	var Jqs []autocode.Junquan
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.URL != "" {
		db = db.Where("url = ?", info.URL)
	}
	if info.Username != "" {
		db = db.Where("username = ?", info.Username)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	db.Preload("User")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&Jqs).Error
	return err, Jqs, total
}

func (JqService *JunquanService) Jump(userId uint) (err error, url string) {

	Jq := autocode.Junquan{}
	err = global.GVA_DB.Where("user_id = ?", userId).First(&Jq).Error
	if err != nil {
		return err, ""
	}

	timeUTC := time.Now().UnixNano() / 1e6
	str := Jq.Username + Jq.Password + fmt.Sprintf("%d", timeUTC)
	token := utils.MD5V([]byte(str))



	urlstr := fmt.Sprintf("%s/loginLoading?username=%s&time=%d&token=%s", Jq.URL, Jq.Username, timeUTC, token)
	return nil, urlstr
}
