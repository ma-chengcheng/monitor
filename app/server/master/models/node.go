package models

import (
	"gorm.io/gorm"
	"master/utils"
)

func init() {
	if utils.MysqlDBErr == nil && !utils.MysqlDB.Migrator().HasTable(&Node{}) {
		_ = utils.MysqlDB.Migrator().CreateTable(&Node{})
	}
}

type Node struct {
	gorm.Model
	IP          string `gorm:"not null;unique" json:"ip" `
	HostName    string `gorm:"not null;unique" json:"host_name"`
	SSHUsername string `gorm:"not null" json:"ssh_username"`
	SSHPassword string `json:"ssh_password" gorm:"not null" json:"ssh_username"`
}

type NodeInfo struct {
	IP            string `json:"ip"`
	HostName      string `json:"host_name"`
	Status        bool   `json:"status"`
	CpuPercent    string `json:"cpu_percent"`
	MemoryPercent string `json:"memory_percent"`
	DiskPercent   string `json:"disk_percent"`
}

func AddNode(node Node) {
	utils.MysqlDB.Create(&node)
	utils.RedisDB.Do(utils.RedisDBCtx, "hset", "hosts", node.IP, node.HostName)
}

func DeleteNode(ip string) {
	var node Node
	utils.MysqlDB.Where("IP = ?", ip).Delete(&node)
}

func CheckNode(ip string) bool {
	var count int64
	utils.MysqlDB.Model(Node{}).Where(Node{IP: ip}).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func GetNodeInfoList() []NodeInfo {
	var nodes []Node
	var nodeInfoList []NodeInfo

	utils.MysqlDB.Select("IP", "HostName").Find(&nodes)
	for _, node := range nodes {
		nodeInfoList = append(nodeInfoList, NodeInfo{
			IP:       node.IP,
			HostName: node.HostName,
		})
	}

	return nodeInfoList
}
