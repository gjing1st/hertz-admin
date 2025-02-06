// Path: internal/apiserver/store/database/system
// FileName: category.go
// Created by bestTeam
// Author: GJing
// Date: 2022/10/31$ 18:15$

package initdata

import (
	"gorm.io/gorm"
)

type InitConfig struct {
}

// DataInserted
// @description: 数据是否已插入
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:19
// @success:
func (i InitConfig) DataInserted(db *gorm.DB) bool {
	return true
}

// InitializeData
// @description: 初始化数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 17:26
// @success:
func (i InitConfig) InitializeData(db *gorm.DB) (err error) {
	return
}
