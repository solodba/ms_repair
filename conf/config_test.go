package conf_test

import (
	"testing"

	"github.com/solodba/ms_repair/conf"
)

func TestGetDbConn(t *testing.T) {
	conf.Conf = conf.NewDefaultConfig()
	conf.Conf.MySQL.Username = "root"
	conf.Conf.MySQL.Password = "Root@123"
	conf.Conf.MySQL.Host = "118.195.139.53"
	conf.Conf.MySQL.Port = 3306
	conf.Conf.MySQL.DB = "mysql"
	conf.Conf.MySQL.MaxOpenConn = 50
	conf.Conf.MySQL.MaxIdleConn = 10
	conf.Conf.MySQL.MaxLifeTime = 600
	conf.Conf.MySQL.MaxIdleTime = 300
	conn, err := conf.Conf.MySQL.GetDbConn()
	if err != nil {
		t.Fatal(err)
	}
	row, err := conn.Query("show slave status")
	defer row.Close()
	if err != nil {
		t.Fatal(err)
	}
	for row.Next() {
		var result string
		err = row.Scan(&result)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}
