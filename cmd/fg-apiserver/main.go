// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package main

import (
	"os"

	"github.com/onexstack/fastgo/cmd/fg-apiserver/app"
	// 导入 automaxprocs 包，可以在程序启动时自动设置 GOMAXPROCS 配置，
	// 使其与 Linux 容器的 CPU 配额相匹配。
	// 这避免了在容器中运行时，因默认 GOMAXPROCS 值不合适导致的性能问题，
	// 确保 Go 程序能够充分利用可用的 CPU 资源，避免 CPU 浪费。
	_ "go.uber.org/automaxprocs"
)

// Go 程序的默认入口函数。阅读项目代码的入口函数.
func main() {
	// 创建 Go 极速项目
	command := app.NewFastGOCommand()

	// 执行命令并处理错误
	if err := command.Execute(); err != nil {
		// 如果发生错误，则退出程序
		// 返回退出码，可以使其他程序（例如 bash 脚本）根据退出码来判断服务运行状态
		os.Exit(1)
	}
}
