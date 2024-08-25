package define

import (
	"github.com/loebfly/ezgin/ezcfg"
	"github.com/loebfly/ezgin/ezlogs"
	"os"
	"path/filepath"
	"strings"
)

func Setup() {
	ymlPath := getYmlPath()
	data, err := ezcfg.GetYmlData(ymlPath)
	if err != nil {
		ezlogs.Error("配置文件解析错误:" + err.Error())
		panic(err)
	}
	Yml = yml{data: data}
}

// getLocalYml 获取yml配置文件路径
func getYmlPath() string {
	fileName := ezcfg.GetString("ezgin.app.name") + ".yml"
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if strings.Contains(fileName, "/") || strings.Contains(fileName, "\\") {
		return fileName
	}
	return path + "/" + fileName
}
