#!/bin/bash

# newbie 项目源码根目录
PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")

# 构建产物输出目录
OUTPUT_DIR=${PROJ_ROOT_DIR}/_output

## 指定应用使用的 version 包，会通过 `-ldflags -X` 向该包中指定的变量注入值
VERSION_PACKAGE=github.com/onexstack/fastgo/pkg/version

## 定义 VERSION 语义化版本号
if [[ -z "${VERSION}" ]];then
  VERSION=$(git describe --tags --always --match='v*')
fi

## 检查代码仓库是否是 dirty（默认dirty）
GIT_TREE_STATE="dirty"
is_clean=$(git status --porcelain 2>/dev/null)
if [[ -z ${is_clean} ]];then
  GIT_TREE_STATE="clean"
fi

# 向代码中传入版本信息
GIT_COMMIT=$(git rev-parse HEAD)
GO_LDFLAGS="-X ${VERSION_PACKAGE}.gitVersion=${VERSION} \
  -X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT} \
  -X ${VERSION_PACKAGE}.gitTreeState=${GIT_TREE_STATE} \
  -X ${VERSION_PACKAGE}.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

# 编译 newbie 源码
go build -v -ldflags "${GO_LDFLAGS}" -o ${OUTPUT_DIR}/fg-apiserver -v cmd/fg-apiserver/main.go
