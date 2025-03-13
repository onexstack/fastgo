#!/bin/bash

# 获取脚本所在目录作为项目根目录
PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")

# 定义构建产物的输出目录为项目根目录下的_output文件夹
OUTPUT_DIR=${PROJ_ROOT_DIR}/_output

# 指定版本信息包的路径，后续会通过-ldflags参数将版本信息注入到这个包的变量中
VERSION_PACKAGE=github.com/onexstack/fastgo/pkg/version

# 确定VERSION值：如果环境变量中没有设置VERSION，则使用git标签作为版本号
# git describe --tags --always --match='v*'命令会获取最近的v开头的标签，如果没有则使用当前commit的短哈希
if [[ -z "${VERSION}" ]];then
  VERSION=$(git describe --tags --always --match='v*')
fi

# 检查代码仓库状态：判断工作目录是否干净
# 默认状态设为"dirty"（有未提交更改）
GIT_TREE_STATE="dirty"
# 使用git status检查是否有未提交的更改
is_clean=$(shell git status --porcelain 2>/dev/null)
# 如果is_clean为空，说明没有未提交的更改，状态设为"clean"
if [[ -z ${is_clean} ]];then
  GIT_TREE_STATE="clean"
fi

# 获取当前git commit的完整哈希值
GIT_COMMIT=$(git rev-parse HEAD)

# 构造链接器标志（ldflags）
# 通过-X选项向VERSION_PACKAGE包中注入以下变量的值：
# - gitVersion: 版本号
# - gitCommit: 构建时的commit哈希
# - gitTreeState: 代码仓库状态(clean或dirty)
# - buildDate: 构建日期和时间(UTC格式)
GO_LDFLAGS="-X ${VERSION_PACKAGE}.gitVersion=${VERSION} \
  -X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT} \
  -X ${VERSION_PACKAGE}.gitTreeState=${GIT_TREE_STATE} \
  -X ${VERSION_PACKAGE}.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

# 执行Go构建命令
# -v: 显示详细编译信息
# -ldflags: 传入上面定义的链接器标志
# -o: 指定输出文件路径和名称
# 最后参数是入口文件路径
go build -v -ldflags "${GO_LDFLAGS}" -o ${OUTPUT_DIR}/fg-apiserver -v cmd/fg-apiserver/main.go
