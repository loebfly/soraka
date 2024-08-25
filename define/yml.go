package define

import (
	"fmt"
	"github.com/knadh/koanf"
	"strings"
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

const (
	ymlDiskListenPath             = "disk.listen.path"     // 磁盘装载路径列表(字符串数组)
	ymlDiskListenInterval         = "disk.listen.interval" // 监听磁盘路径间隔（Int）
	ymlDiskCleanPrefix            = "disk.clean"           // 清理磁盘配置
	ymlDiskCleanXPath             = "disk.clean.%s"        // 清理磁盘配置路径
	ymlDiskCleanRuleWhenUsage     = "rule.when_usage"      // 当磁盘使用率超过多少时
	ymlDiskCleanRuleIncludeSuffix = "rule.include_suffix"  // 包含规则
	ymlDiskCleanRuleExcludeSuffix = "rule.exclude_suffix"  // 排除规则
	ymlDiskCleanRuleBeforeTime    = "rule.before_time"     // 清理多少天之前的文件
)

type yml struct {
	data *koanf.Koanf
}

var Yml yml

func (receiver yml) DiskListenPaths() []string {
	return receiver.data.Get(ymlDiskListenPath).([]string)
}

func (receiver yml) DiskListenInterval() time.Duration {
	return time.Duration(receiver.data.Int64(ymlDiskListenInterval)) * time.Minute
}

func (receiver yml) DiskCleanPaths() map[string][]string {
	var cleanPaths = make(map[string][]string)
	listenPaths := receiver.DiskListenPaths()
	if len(listenPaths) == 0 {
		return cleanPaths
	}
	for _, path := range listenPaths {
		cleanPaths[path] = receiver.data.Get(fmt.Sprintf(ymlDiskCleanXPath, path)).([]string)
	}
	return cleanPaths
}

func (receiver yml) DiskCleanRuleWhenUsage() float64 {
	return receiver.data.Float64(ymlDiskCleanRuleWhenUsage)
}

func (receiver yml) DiskCleanRuleIncludeSuffixes() []string {
	suffix := receiver.data.String(ymlDiskCleanRuleIncludeSuffix)
	if suffix == "" {
		return []string{}
	}
	return strings.Split(suffix, ",")
}

func (receiver yml) DiskCleanRuleExcludeSuffixes() []string {
	suffix := receiver.data.String(ymlDiskCleanRuleExcludeSuffix)
	if suffix == "" {
		return []string{}
	}
	return strings.Split(suffix, ",")
}

func (receiver yml) DiskCleanRuleBeforeTime() time.Duration {
	return time.Duration(receiver.data.Int64(ymlDiskCleanRuleBeforeTime))
}
