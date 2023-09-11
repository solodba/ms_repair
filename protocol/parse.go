package protocol

import (
	"context"
	"fmt"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	"github.com/solodba/ms_repair/apps/repair"
	"github.com/solodba/ms_repair/conf"
)

var (
	ctx = context.Background()
)

// master-slave repair服务结构体
type MsRepairService struct {
	svc repair.Service
}

// master-slave repair服务结构体构造函数
func NewMsRepairService() *MsRepairService {
	return &MsRepairService{
		svc: apps.GetInternalApp(repair.AppName).(repair.Service),
	}
}

// master-slave repair服务启动方法
func (m *MsRepairService) Start() error {
	querySlaveErrorRes, err := m.svc.QuerySlaveError(ctx)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("=========================从库错误信息如下============================")
	logger.L().Info().Msgf("从库[%s]错误代码: %d", conf.Conf.MySQL.Host, querySlaveErrorRes.ErrNum)
	logger.L().Info().Msgf("从库[%s]错误信息: %s", conf.Conf.MySQL.Host, querySlaveErrorRes.ErrMsg)
	logger.L().Info().Msgf("从库[%s]错误发生时间: %s", conf.Conf.MySQL.Host, querySlaveErrorRes.ErrTime)
	logger.L().Info().Msgf("=========================开始修复从库报错=============================")
	querySlaveModeRes, err := m.svc.QuerySlaveMode(ctx)
	if err != nil {
		return err
	}
	switch querySlaveErrorRes.ErrNum {
	case 1062:
		if querySlaveModeRes.Mode == "POSITION" {
			logger.L().Info().Msgf("从库[%s]的复制模式: %s", conf.Conf.MySQL.Host, querySlaveModeRes.Mode)
			err = m.svc.BasePositionRepair(ctx)
			if err != nil {
				logger.L().Info().Msgf("从[%s]库修复失败, 原因: %s", conf.Conf.MySQL.Host, err.Error())
				return err
			}
			logger.L().Info().Msgf("=================从库[%s]修复成功================", conf.Conf.MySQL.Host)
		}
		if querySlaveModeRes.Mode == "GTID" {
			logger.L().Info().Msgf("从库[%s]的复制模式: %s", conf.Conf.MySQL.Host, querySlaveModeRes.Mode)
			err = m.svc.BaseGtidRepair(ctx)
			if err != nil {
				logger.L().Info().Msgf("从[%s]库修复失败, 原因: %s", conf.Conf.MySQL.Host, err.Error())
				return err
			}
			logger.L().Info().Msgf("=================从库[%s]修复成功================", conf.Conf.MySQL.Host)
		}
	default:
		return fmt.Errorf("该类型不支持处理, 请自行手动处理")
	}
	return nil
}

// master-slave repair服务停止方法
func (s *MsRepairService) Stop() error {
	return nil
}
