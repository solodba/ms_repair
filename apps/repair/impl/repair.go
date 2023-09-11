package impl

import (
	"context"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/solodba/mcube/logger"
	"github.com/solodba/ms_repair/apps/repair"
	"github.com/solodba/ms_repair/conf"
)

// 查看从库是否有错误
func (i *impl) QuerySlaveError(ctx context.Context) (*repair.QuerySlaveErrorResponse, error) {
	sql := `select LAST_ERROR_NUMBER,LAST_ERROR_MESSAGE,LAST_ERROR_TIMESTAMP from performance_schema.replication_applier_status_by_worker ORDER BY LAST_ERROR_TIMESTAMP desc limit 1`
	row := i.db.QueryRowContext(ctx, sql)
	res := repair.NewQuerySlaveErrorResponse()
	err := row.Scan(&res.ErrNum, &res.ErrMsg, &res.ErrTime)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 查询从库是GTID模式或者基于Position模式
func (i *impl) QuerySlaveMode(ctx context.Context) (*repair.QuerySlaveModeResponse, error) {
	sql := `show global variables like 'enforce_gtid_consistency'`
	var name_a, value_a string
	row := i.db.QueryRowContext(ctx, sql)
	err := row.Scan(&name_a, &value_a)
	if err != nil {
		return nil, err
	}
	sql = `show global variables like 'gtid_mode'`
	var name_b, value_b string
	row = i.db.QueryRowContext(ctx, sql)
	err = row.Scan(&name_b, &value_b)
	if err != nil {
		return nil, err
	}
	res := repair.NewQuerySlaveModeResponse()
	if strings.ToUpper(value_a) == "ON" && strings.ToUpper(value_b) == "ON" {
		res.Mode = "GTID"
	}
	if strings.ToUpper(value_a) == "OFF" && strings.ToUpper(value_b) == "OFF" {
		res.Mode = "POSITION"
	}
	return res, nil
}

// Base Position的修复
func (i *impl) BasePositionRepair(ctx context.Context) error {
	sql := `stop slave;set global sql_slave_skip_counter=1;start slave`
	_, err := i.db.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("设置参数[sql_slave_skip_counter=1]跳过错误成功")
	return nil
}

// 查询从库执行过的GTID
func (i *impl) QuerySlaveExcuteGtid(ctx context.Context) (*repair.QuerySlaveExcuteGtidResponse, error) {
	sql := `select APPLYING_TRANSACTION from performance_schema.replication_applier_status_by_worker ORDER BY LAST_ERROR_TIMESTAMP desc limit 1`
	row := i.db.QueryRowContext(ctx, sql)
	res := repair.NewQuerySlaveExcuteGtidResponse()
	err := row.Scan(&res.ApplyingTransaction)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Base GTID的修复
func (i *impl) BaseGtidRepair(ctx context.Context) error {
	excuteGtid, err := i.QuerySlaveExcuteGtid(ctx)
	if err != nil {
		return err
	}
	sql := fmt.Sprintf(`stop slave;set session gtid_next='%s';begin;commit;set session gtid_next='automatic';start slave`, excuteGtid.ApplyingTransaction)
	_, err = i.db.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("从库[%s]跳过GTID[%s]成功", conf.Conf.MySQL.Host, excuteGtid.ApplyingTransaction)
	return nil
}
