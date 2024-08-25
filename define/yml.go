package define

import (
	"fmt"
	"github.com/loebfly/ezgin/ezcfg"
	"strings"
	"time"
)

const (
	ymlDiskListenPath             = "disk.listen.path"               // 磁盘装载路径列表(字符串数组)
	ymlDiskListenInterval         = "disk.listen.interval"           // 监听磁盘路径间隔（Int）
	ymlDiskCleanXPath             = "disk.clean.%s"                  // 清理磁盘配置路径
	ymlDiskCleanRuleWhenUsage     = "disk.clean.rule.when_usage"     // 当磁盘使用率超过多少时
	ymlDiskCleanRuleIncludeSuffix = "disk.clean.rule.include_suffix" // 包含规则
	ymlDiskCleanRuleExcludeSuffix = "disk.clean.rule.exclude_suffix" // 排除规则
	ymlDiskCleanRuleBeforeTime    = "disk.clean.rule.before_time"    // 清理多少天之前的文件
)

type yml int

const Yml = yml(0)

func (receiver yml) DiskListenPaths() []string {
	return ezcfg.GetArray[string](ymlDiskListenPath)
}

func (receiver yml) DiskListenInterval() time.Duration {
	return time.Duration(ezcfg.GetInt64(ymlDiskListenInterval)) * time.Second
}

func (receiver yml) DiskCleanPaths() map[string][]string {
	var cleanPaths = make(map[string][]string)
	listenPaths := receiver.DiskListenPaths()
	if len(listenPaths) == 0 {
		return cleanPaths
	}
	for _, path := range listenPaths {
		cleanPaths[path] = ezcfg.GetArray[string](fmt.Sprintf(ymlDiskCleanXPath, path))
	}
	return cleanPaths
}

func (receiver yml) DiskCleanRuleWhenUsage() float64 {
	return ezcfg.GetFloat64(ymlDiskCleanRuleWhenUsage)
}

func (receiver yml) DiskCleanRuleIncludeSuffixes() []string {
	suffix := ezcfg.GetString(ymlDiskCleanRuleIncludeSuffix)
	if suffix == "" {
		return []string{}
	}
	return strings.Split(suffix, ",")
}

func (receiver yml) DiskCleanRuleExcludeSuffixes() []string {
	suffix := ezcfg.GetString(ymlDiskCleanRuleExcludeSuffix)
	if suffix == "" {
		return []string{}
	}
	return strings.Split(suffix, ",")
}

func (receiver yml) DiskCleanRuleBeforeTime() time.Duration {
	return time.Duration(ezcfg.GetInt64(ymlDiskCleanRuleBeforeTime))
}
