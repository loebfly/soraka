package service

import "github.com/loebfly/ezgin/engine"

type diskService int

const DiskService = diskService(0)

func (receiver diskService) DelPath(path string) engine.Result[any] {
	return engine.SuccessRes(nil, "清理完成").ToAnyRes()
}
