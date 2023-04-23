// Path: internal/apiserver/store/database/system
// FileName: category.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 18:15$

package initdata

import (
	"errors"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/dict"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"gorm.io/gorm"
)

type InitCategory struct {
}

// DataInserted
// @description: 数据是否已插入
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:19
// @success:
func (i InitCategory) DataInserted(db *gorm.DB) bool {
	if errors.Is(db.Where("name = ?", "商用密码标准规范").First(&entity.Category{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// InitializeData
// @description: 初始化数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 17:26
// @success:
func (i InitCategory) InitializeData(db *gorm.DB) (err error) {
	if db == nil {
		return DBNullErr
	}
	entities := []entity.Category{
		{Name: "系统通知", Status: dict.StatusEnable, BaseModel: entity.BaseModel{ID: dict.CategoryNotification}}, //ID:1
		{Name: "帮助中心", Status: dict.StatusEnable, BaseModel: entity.BaseModel{ID: dict.CategoryHelp}},
		{Name: "安全资讯", Status: dict.StatusEnable, BaseModel: entity.BaseModel{ID: dict.CategorySafeInformation}},
		{Name: "法律法规", Status: dict.StatusEnable, BaseModel: entity.BaseModel{ID: dict.CategoryLawsRegulations}, ParentId: dict.CategoryIndustryPolicy},
		{Name: "商用密码标准规范", Status: dict.StatusEnable, BaseModel: entity.BaseModel{ID: dict.CategoryCipherStandard}, ParentId: dict.CategoryIndustryPolicy},
		{Name: "等保政策文件", Status: dict.StatusEnable, BaseModel: entity.BaseModel{ID: dict.CategoryGradePolicy}, ParentId: dict.CategoryIndustryPolicy},
		{Name: "密评政策文件", Status: dict.StatusEnable, BaseModel: entity.BaseModel{ID: dict.CategoryEstimatePolicy}, ParentId: dict.CategoryIndustryPolicy},
		{Name: "行业规范", Status: dict.StatusEnable, BaseModel: entity.BaseModel{ID: dict.CategoryIndustryPolicy}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return InitDataErr
	}
	return
}
