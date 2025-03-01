// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package apiserver

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"

	genericoptions "github.com/onexstack/fastgo/pkg/options"
)

// Config 配置结构体，用于存储应用相关的配置.
// 不用 viper.Get，是因为这种方式能更加清晰的知道应用提供了哪些配置项.
type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
}

// Server 定义一个服务器结构体类型.
type Server struct {
	cfg *Config
}

// NewServer 根据配置创建服务器.
func (cfg *Config) NewServer() (*Server, error) {
	return &Server{cfg: cfg}, nil
}

// Run 运行应用.
func (s *Server) Run() error {
	fmt.Printf("Read MySQL host from Viper: %s\n\n", viper.GetString("mysql.host"))

	jsonData, _ := json.MarshalIndent(s.cfg, "", "  ")
	fmt.Println(string(jsonData))

	select {}
}
