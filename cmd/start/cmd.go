package start

import (
	"github.com/solodba/ms_repair/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "ms_repair start service",
	Long:    "ms_repair service",
	Example: `ms_repair start -u root -p Root@123 -h localhost -P3306`,
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewServer()
		if err := srv.Start(); err != nil {
			return err
		}
		return nil
	},
}

// 服务结构体
type Server struct {
	MsRepairService *protocol.MsRepairService
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		MsRepairService: protocol.NewMsRepairService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	if err := s.MsRepairService.Start(); err != nil {
		return err
	}
	return nil
}

// Server服务停止方法
func (s *Server) Stop() error {
	if err := s.MsRepairService.Stop(); err != nil {
		return err
	}
	return nil
}
