## fastgo 项目

- 该实战项目配套课程见： [Go 项目开发极速入门课](https://blog.csdn.net/lnxfei/category_12907774.html)
- 本初级项目的中/高级版本为：[miniblog](https://github.com/onexstack/miniblog)
- 更多「云原生 AI 实战营」项目见：[云原生 AI 实战营项目介绍](https://konglingfei.com/cloudai/intro/%E4%BA%91%E5%8E%9F%E7%94%9F_AI_%E5%AE%9E%E6%88%98%E8%90%A5%E9%A1%B9%E7%9B%AE%E4%BB%8B%E7%BB%8D.html)

## fastgo 项目适宜人群

- 掌握一定 Go 基础语法，想通过一个完整的实战，来快速系统学习 Go 项目开发的初学者；
- 有意从事 Go 项目开发，但尚未入门或入门尚浅的同学。

## 项目快速部署

```bash
$ mkdir -p  $HOME/golang/src/github.com/onexstack/
$ cd $HOME/golang/src/github.com/onexstack/
$ git clone https://github.com/onexstack/fastgo
$ cd fastgo/
$ ./build.sh
$ _output/fg-apiserver -c configs/fg-apiserver.yaml
```

**注意：** 

1. 要登录 MySQL 并且执行 `source configs/fastgo.sql;` 创建 `fastgo` 数据库及表；
2. 更新 `configs/fg-apiserver.yaml` 中 `mysql` 配置项。

## fastgo 包含的技能点

技能点见下图所示：

![](./docs/images/skills.jpg)

## Contacts

<img src="./docs/images/three-code.png" alt="" width="900" />

## License

[MIT](https://choosealicense.com/licenses/mit/)
