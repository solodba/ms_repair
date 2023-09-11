package conf

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 创建MySQL连接池
func (m *MySQL) GetConnPool() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&multiStatements=true",
		m.Username, m.Password, m.Host, m.Port, m.DB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("conn mysql<%s> error, reason: %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(int(m.MaxOpenConn))
	db.SetMaxIdleConns(int(m.MaxIdleConn))
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleTime != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, reason: %s", dsn, err.Error())
	}
	return db, nil
}

// 从MySQL连接池获取连接
func (m *MySQL) GetDbConn() (*sql.DB, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.db == nil {
		db, err := m.GetConnPool()
		if err != nil {
			return nil, err
		}
		m.db = db
	}
	return m.db, nil
}
