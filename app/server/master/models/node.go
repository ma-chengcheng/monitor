package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"master/manager/api"
	"master/utils"
	"time"
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
	IP            string  `json:"ip"`
	HostName      string  `json:"host_name"`
	SSHUsername   string  `json:"ssh_username"`
	SSHPassword   string  `json:"ssh_password"`
	CpuPercent    float64 `json:"cpu_percent"`
	MemoryPercent float64 `json:"memory_percent"`
	DiskPercent   float64 `json:"disk_percent"`
	Status        bool    `json:"status"`
}

func CheckNode(ip string) bool {
	var count int64
	utils.MysqlDB.Model(Node{}).Where(Node{IP: ip}).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func AddNode(node Node) {
	utils.MysqlDB.Create(&node)
	utils.RedisDB.Do(utils.RedisDBCtx, "SADD", "hosts", node.IP)
}

func DeleteNode(ip string) {
	var node Node
	utils.MysqlDB.Where("IP = ?", ip).Delete(&node)
	utils.RedisDB.Do(utils.RedisDBCtx, "SREM", "hosts", node.IP)
}

func GetNodeInfoList() []NodeInfo {
	var nodes []Node
	var nodeInfoList []NodeInfo

	utils.MysqlDB.Find(&nodes)

	for _, node := range nodes {
		res, err := utils.RedisDB.Get(utils.RedisDBCtx, node.IP).Result()
		status := false

		nodeInfo := api.NodeInfo{}
		if err == nil {
			json.Unmarshal([]byte(res), &nodeInfo)
			status = true
		}

		nodeInfoList = append(nodeInfoList, NodeInfo{
			IP:            node.IP,
			HostName:      node.HostName,
			SSHUsername:   node.SSHUsername,
			SSHPassword:   node.SSHPassword,
			CpuPercent:    nodeInfo.CpuPercent,
			DiskPercent:   nodeInfo.DiskPercent,
			MemoryPercent: nodeInfo.MemoryPercent,
			Status:        status,
		})
	}

	return nodeInfoList
}

func GetIPList() []string {
	res, err := utils.RedisDB.SMembers(utils.RedisDBCtx, "hosts").Result()
	var ipList []string
	if err == nil {
		ipList = res
	}
	return ipList
}

func SetNodeInfo(ip, nodeInfo string)  {
	utils.RedisDB.Set(utils.RedisDBCtx, ip, nodeInfo, time.Second * 30)
}
