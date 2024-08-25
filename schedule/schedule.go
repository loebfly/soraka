package schedule

import (
	"fmt"
	//"github.com/robfig/cron"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"time"
)

func Start() {
	//c := cron.New()
	//checkDiskSpec := ezcfg.GetString("schedule.spec.check_disk")
	//err := c.AddFunc(checkDiskSpec, func() {
	//
	//})
	//if err != nil {
	//	ezlogs.Debug("添加检测硬盘内存定时任务失败:{}", err.Error())
	//	return
	//}
	//ezlogs.Debug("添加检测硬盘内存定时任务成功")
	//c.Start(）

	DiskSchedule.Start()
}

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

// cpu 负载
func getCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}

// mem info
func getMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}

// host info
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// disk info
func getDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	//ioStat, _ := disk.IOCounters()
	//for k, v := range ioStat {
	//	fmt.Printf("%v:%v\n", k, v)
	//}
}

// net IO
func getNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}
