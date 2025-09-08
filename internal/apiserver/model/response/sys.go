// Path: internal/apiserver/model/response
// FileName: sys.go
// Created by bestTeam
// Author: GJing
// Date: 2023/2/9$ 20:01$

package response

// ServerStatus 设备状态
type ServerStatus struct {
	ServiceStatus int `json:"service_status"`
	RunStatus     int `json:"run_status"`
}

// GetNetwork 获取网卡配置
type GetNetwork struct {
	Admin Network `json:"admin"`
	SDF   Network `json:"sdf"`
}

type Network struct {
	Addr    string `json:"addr"`
	Gateway string `json:"gateway"`
	Netmask string `json:"netmask"`
}

// UpdateVersionInfo 升级时的版本信息
type UpdateVersionInfo struct {
	CurrentVersion string `json:"version_current"`
	LatestVersion  string `json:"version_latest"`
	CanUpdate      bool   `json:"can_update"` //可升级
}

// AutoUpdateConfig 自动升级配置
type AutoUpdateConfig struct {
	AutoUpdate  bool   `json:"auto_update"`
	UpdateRange string `json:"update_range"`
	Time        string `json:"time"`
}

type PwdValidDate struct {
	ValidDate int `json:"valid_date"`
}
