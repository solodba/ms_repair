package repair

// QuerySlaveErrorResponse结构体
type QuerySlaveErrorResponse struct {
	ErrNum  int64
	ErrMsg  string
	ErrTime string
}

// QuerySlaveErrorResponse构造函数
func NewQuerySlaveErrorResponse() *QuerySlaveErrorResponse {
	return &QuerySlaveErrorResponse{}
}

// QuerySlaveModeResponse结构体
type QuerySlaveModeResponse struct {
	Mode string
}

// QuerySlaveModeResponse构造函数
func NewQuerySlaveModeResponse() *QuerySlaveModeResponse {
	return &QuerySlaveModeResponse{}
}

// QuerySlaveExcuteGtidResponse结构体
type QuerySlaveExcuteGtidResponse struct {
	ApplyingTransaction string
}

// QuerySlaveExcuteGtidResponse构造函数
func NewQuerySlaveExcuteGtidResponse() *QuerySlaveExcuteGtidResponse {
	return &QuerySlaveExcuteGtidResponse{}
}
