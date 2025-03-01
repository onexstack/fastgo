// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package options

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLOptions defines options for mysql database.
type MySQLOptions struct {
	Addr                  string        `json:"addr,omitempty" mapstructure:"addr"`
	Username              string        `json:"username,omitempty" mapstructure:"username"`
	Password              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
}

// NewMySQLOptions create a `zero` value instance.
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:                  "127.0.0.1:3306",
		Username:              "onex",
		Password:              "onex(#)666",
		Database:              "onex",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

// Validate verifies flags passed to MySQLOptions.
func (o *MySQLOptions) Validate() error {
	// 验证MySQL地址格式
	if o.Addr == "" {
		return fmt.Errorf("MySQL server address cannot be empty")
	}
	// 检查地址格式是否为host:port
	host, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("Invalid MySQL address format '%s': %w", o.Addr, err)
	}
	// 验证端口是否为数字
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("Invalid MySQL port: %s", portStr)
	}
	// 验证主机名是否为空
	if host == "" {
		return fmt.Errorf("MySQL hostname cannot be empty")
	}

	// 验证凭据和数据库名
	if o.Username == "" {
		return fmt.Errorf("MySQL username cannot be empty")
	}

	if o.Password == "" {
		return fmt.Errorf("MySQL password cannot be empty")
	}

	if o.Database == "" {
		return fmt.Errorf("MySQL database name cannot be empty")
	}

	// 验证连接池参数
	if o.MaxIdleConnections <= 0 {
		return fmt.Errorf("MySQL max idle connections must be greater than 0")
	}

	if o.MaxOpenConnections <= 0 {
		return fmt.Errorf("MySQL max open connections must be greater than 0")
	}

	if o.MaxIdleConnections > o.MaxOpenConnections {
		return fmt.Errorf("MySQL max idle connections cannot be greater than max open connections")
	}

	if o.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("MySQL max connection lifetime must be greater than 0")
	}

	return nil
}

// DSN return DSN from MySQLOptions.
func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Addr,
		o.Database,
		true,
		"Local")
}

// NewDB create mysql store with the given config.
func (o *MySQLOptions) NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.DSN()), &gorm.Config{
		// PrepareStmt executes the given query in cached statement.
		// This can improve performance.
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(o.MaxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)

	return db, nil
}
