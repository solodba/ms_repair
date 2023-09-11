package impl

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_repair/apps/repair"
	"github.com/solodba/ms_repair/conf"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	db *sql.DB
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return repair.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	db, err := conf.Conf.MySQL.GetDbConn()
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
}
