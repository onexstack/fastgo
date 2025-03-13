#!/usr/bin/env bash

# Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

# Common utilities, variables and checks for all build scripts.
set -o errexit
set -o nounset
set -o pipefail

# The root of the build/dist directory
PROJ_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

INSECURE_SERVER="127.0.0.1:6666"

Header="-HContent-Type: application/json"
CCURL="curl -s -XPOST" # Create
UCURL="curl -s -XPUT" # Update
RCURL="curl -s -XGET" # Retrieve
DCURL="curl -s -XDELETE" # Delete

# 随机生成用户名
fg::test::username()
{
  echo fastgo$(date +%s)

}

# 注意：使用 root 用户登录系统，否则无法删除指定的用户
fg::test::login()
{
  ${CCURL} "${Header}" http://${INSECURE_SERVER}/login \
    -d'{"username":"'$1'","password":"'$2'"}' | grep -Po 'token[" :]+\K[^"]+'
}

# 用户相关接口测试函数
fg::test::user()
{
  username=$(fg::test::username)
  # 1. 创建 fastgo 用户
  ${CCURL} "${Header}" http://${INSECURE_SERVER}/v1/users \
    -d'{"username":"'${username}'","password":"fastgo1234","nickname":"fastgo","email":"colin404@foxmail.com","phone":"'$(date +%s)'"}'; echo
  echo -e "\033[32m1. 成功创建 ${username} 用户\033[0m"

  token="-HAuthorization: Bearer $(fg::test::login ${username} fastgo1234)"

  # 2. 列出所有用户
  ${RCURL} "${token}" "http://${INSECURE_SERVER}/v1/users?offset=0&limit=10"; echo
  echo -e "\033[32m2. 成功列出所有用户\033[0m"

  # 3. 获取 fastgo 用户的详细信息
  ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/users/${username}; echo
  echo -e "\033[32m3. 成功获取 ${username} 用户详细信息\033[0m"

  # 4. 修改 fastgo 用户
  ${UCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/users/${username} \
    -d'{"nickname":"fastgo(modified)"}'; echo
  echo -e "\033[32m4. 成功修改 ${username} 用户信息\033[0m"

  # 5. 删除 fastgo 用户
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/users/${username}; echo
  echo -e "\033[32m5. 成功删除 ${username} 用户\033[0m"

  echo -e '\033[32m==> 所有用户接口测试成功\033[0m'
}

# 博客相关接口测试函数
fg::test::post()
{

  username=$(fg::test::username)
  # 1. 创建测试用户
  ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/users \
    -d'{"username":"'${username}'","password":"fastgo1234","nickname":"fastgo","email":"colin404@foxmail.com","phone":"'$(date +%s)'"}'; echo
  echo -e "\033[32m1. 成功创建测试用户: ${username}\033[0m"

  token="-HAuthorization: Bearer $(fg::test::login ${username} fastgo1234)"

  # 2. 创建一条博客
  postID=`${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/posts -d'{"title":"installation","content":"installation."}' | grep -Po 'post-[a-z0-9]+'`
  echo -e "\033[32m2. 成功创建博客: ${postID}\033[0m"

  # 3. 列出所有博客
  ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/posts; echo
  echo -e "\033[32m3. 成功列出所有博客\033[0m"

  # 4. 获取所创建博客的信息
  ${RCURL} "${token}" http://${INSECURE_SERVER}/v1/posts/${postID}; echo
  echo -e "\033[32m4. 成功获取博客 ${postID} 详细信息\033[0m"

  # 6. 修改所创建博客的信息
  ${UCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/posts/${postID} -d'{"title":"modified"}'; echo
  echo -e "\033[32m5. 成功更新博客 ${postID} 信息\033[0m"

  # 7. 删除所创建的博客
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/posts -d'{"postIDs":["'${postID}'"]}'; echo
  echo -e "\033[32m6. 成功删除博客 ${postID}\033[0m"

  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/users/${username}; echo
  echo -e "\033[32m7. 成功删除测试用户：${username}\033[0m"

  echo -e '\033[32m==> 所有博客接口测试成功\033[0m'
}

# 测试 user 资源 CURD
fg::test::user

# 测试 post 资源 CURD
fg::test::post

echo -e '\033[32m==> 所有 fastgo 接口测试成功\033[0m'
