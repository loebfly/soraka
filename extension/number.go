package extension

import "fmt"

type number int

const Number = number(0)

func (receiver number) FormatUint64(n uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"} // 单位
	i := 0                                               // 计数器
	for n >= 1024 && i < len(units)-1 {
		n /= 1024
		i++
	}
	// 保留两位小数
	return fmt.Sprintf("%.2f %s", float64(n), units[i])
}
