package repair

import "context"

// 模块名称
const (
	AppName = "pos"
)

// 基于Position位点服务接口
type Service interface {
	// 查看从库是否有错误
	QuerySlaveError(context.Context) (*QuerySlaveErrorResponse, error)
	// 查询从库是GTID模式或者基于Position模式
	QuerySlaveMode(context.Context) (*QuerySlaveModeResponse, error)
	// Base Position的修复
	BasePositionRepair(context.Context) error
	// 查询从库执行过的GTID
	QuerySlaveExcuteGtid(context.Context) (*QuerySlaveExcuteGtidResponse, error)
	// Base GTID的修复
	BaseGtidRepair(context.Context) error
}
