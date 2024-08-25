package schedule

import (
	"github.com/loebfly/ezgin/ezlogs"
	"github.com/shirou/gopsutil/v4/disk"
	"soraka/define"
	"time"
)

/*
disk: # 磁盘路径
  listen:
    path: # 磁盘装载路径
      - "/System/Volumes/Data" # 本机磁盘路径
      - "/Volumes/QingGe"
    interval: 60 # 监听磁盘路径间隔, 默认为60秒
  clean: # 清理磁盘配置
    "/System/Volumes/Data": # 和listen.path一一对应
       - "/Users/luchunqing/Desktop/temp" # 需要清理的文件路径
    "/Volumes/QingGe":
       - "/Volumes/QingGe/temp"
    rule: # 清理规则
      when_usage: 80 # 当磁盘使用率超过多少时, 默认为80%
      include_suffix: "" # 包含规则, 默认包含所有文件
      exclude_suffix: "" # 排除规则, 默认为不排除任何文件
      before_time: 0 # 清理多少天之前的文件, 默认为0天，表示不限制，允许删所有指定路径下的所有文件
*/

type diskSchedule int

const DiskSchedule = diskSchedule(0)

func (receiver diskSchedule) Start() {
	diskListenPaths := define.Yml.DiskListenPaths()
	diskCleanPaths := define.Yml.DiskCleanPaths()
	diskCleanRuleWhenUsage := define.Yml.DiskCleanRuleWhenUsage()
	ticker := time.NewTicker(define.Yml.DiskListenInterval())
	defer ticker.Stop()
	for range ticker.C {
		for _, listenPath := range diskListenPaths {
			receiver.checkDiskUsage(listenPath, diskCleanRuleWhenUsage, diskCleanPaths[listenPath])
		}
	}
}

func (receiver diskSchedule) checkDiskUsage(listenPath string, diskCleanRuleWhenUsage float64, cleanPaths []string) {
	usageStat, err := disk.Usage(listenPath)
	if err != nil {
		ezlogs.Error("获取{}磁盘使用率失败, err:", listenPath, err)
		return
	}

	usagePercent := usageStat.UsedPercent

	if usagePercent > diskCleanRuleWhenUsage {
		ezlogs.Error("{}磁盘使用率超过{}%，开始清理文件!", listenPath, diskCleanRuleWhenUsage)
		for _, cleanPath := range cleanPaths {
			receiver.cleanDirFiles(cleanPath)
		}
	} else {
		ezlogs.Info("{}磁盘使用率为{}%，未超过{}%，无需清理文件!", listenPath, usagePercent, diskCleanRuleWhenUsage)
	}
}

func (receiver diskSchedule) cleanDirFiles(dir string) {
	// Clean up files in the specified directory
}
