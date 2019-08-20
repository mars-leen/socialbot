# socialbot
socialbot可以快速建立媒体资源型网站，集成vultr api 自动化创建部署爬虫程序。

# 开始
##项目结构
    项目结构按照 [golang-standards ](https://github.com/golang-standards/project-layout)构建。
##开始

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
    
 
# socialbot
**socialbot可以快速建立媒体资源型网站，集成vultr api 自动化创建部署爬虫程序。**

### 目录说明

1. 执行database.sql创建数据库,并配置/Config/Mysql.go数据库连接地址
2. 部署定时任务/App/GetHot.go爬虫程序，且以守护进程的方式执行Server.go
3. 打开hot.html今日热榜页面


### 安装教程
项目结构按照 [golang-standards ](https://github.com/golang-standards/project-layout)构建。
```
socailbot/
├── cmd
│   ├── web
│         └── main.go  （web入口）
├── config
│   ├── config.json      （配置文件）
│   └── config.json.template  （配置参考文件）
├── internal
│   ├── common                 （公共的内部代码）
│   ├── robot                  （机器人代码，master分支暂不合并进来）
│   └── web                   （web 代码）
│            ├── cache          (业务缓存)
│            ├── common         (常量 、接口返回公共处理)
│            ├── console corn （1.维护vultr服务器信息）
│            ├── controller  （控制器层，简单的接收请求参数，判断参数）
│            ├── logic      （逻辑层，处理每个接口的业务）
│            ├── service    （服务层，不同业务逻辑都会使用到的方法。）
│            ├── model      (model 层，封装了 增删该查灯方法)
│            ├── middleware （中间件，用户登陆验证，cors, logger, recovery, 前后台静态文件渲染）
│            ├── orm        (xorm 简单封装)
│            ├── setting    (设置，初始化变量、读取配置文件等)
│            ├── wblogger （log）
│            ├── router.go （路由）
│            └── web.go  （启动web服务）
├── pkg （公共包）
│   ├── jwtauth 
│   ├── logger 
│   ├── snowflake (推特分布式唯一id) 
│   ├── urils     (公共的工具方法)  
│   └── vultr      (vultr 接口封装) 
├── scripts
│   ├── socialbot.sql  （sql 文件）
│   └── up.sh           （临时上传脚本，自己使用的，如果需要可参考）
│  
└── README.md
```

### API说明

#### 获取所有类型
- Method: **GET**
- URL:  ```https://www.printf520.com:8080/GetType```
- Param：无
- Body:
```
{
    "Code":0,
    "Message":"获取数据成功",
    "Data":[
        {"id":"1","sort":"63908","title":"知乎"},
        {"id":"2","sort":"21912","title":"虎扑"},
        {"id":"6","sort":"11707","title":"天涯"},
        {"id":"7","sort":"12546","title":"知乎日报"},
       ]}
```


### 获取具体类型热榜数据
- Method: **GET**
- URL:  ```  https://www.printf520.com:8080/GetTypeInfo?id=2```
- Param：id
- Body:
```
{
    "Code":0,
    "Message":"获取成功",
    "Data":[
        {
            "title":"
45个经典面试回答提示，分享给即将工作的大家。 zt
",
            "url":"https://bbs.hupu.com//28814429.html"
        },
        {
            "title":"
[名场面]回家的诱惑：洪世贤酒店幽会，抵不住诱惑犯了错！ zt
",
            "url":"https://bbs.hupu.com//28818367.html"
        },
        {
            "title":"
张艺兴回应假唱风波。ZT
",
            "url":"https://bbs.hupu.com//28815609.html"
        }
    ]
}
```
    
        