# socialbot
socialbot可以快速建立媒体资源型网站，集成vultr api 自动化创建部署爬虫程序。

# 开始
## 项目结构
    项目结构按照 [golang-standards ](https://github.com/golang-standards/project-layout)构建。
## 开始
    1.导入sql文件（已包含创建数据库语句）。 mysql构建版本8.0，低版本mysql不支持 索引的 VISIBLE INVISIBLE ，需要自行移除  
    2.依赖管理。golang构建版本为go1.12.5， 使用go mod 管理依赖。低于 1.12 版本不支持 go mod，需升级。 mod 文件中 加入了 replace 解决 GFW（中国长城防火墙） 问题. 
    如果需要当前目录下vendor目录构建。可在更目录下执行以下代码。
    `go mod vendor`
    3.配置文件。congfig/config.json, 可以按照config.json.template 编写。填上数据库地址。
    4.运行。 进入 cmd/web 目录，执行 go build .
    5.访问监听的端口。
        http://localhost:8081 (前台)
        http://localhost:8081/dashboard （后台管理）
    6.配置参数。除了配置文件外，运行时配置 存放于 mysql 中，需要进入管理后台更改。
        