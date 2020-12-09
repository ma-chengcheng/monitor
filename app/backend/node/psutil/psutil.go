package psutil

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetCPUPercent() float64 {
	percents, _ := cpu.Percent(time.Second, true)
	var sum = 0.0
	
	for _, p := range percents {
		sum += p
	}
	
	if len(percents) != 0 {
		return sum / float64(len(percents))
	}
	return 0
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}
