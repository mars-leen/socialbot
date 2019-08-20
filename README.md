# socialbot
**socialbot可以快速建立媒体资源型网站，集成vultr api 自动化创建部署爬虫程序，对爬取的资源链接可配置反向代理，防止被禁**

### 目录说明
项目结构按照 [golang-standards ](https://github.com/golang-standards/project-layout)构建。
```
socailbot/
├── cmd
│   ├── web
│   │    └── main.go  （web入口）
│   └── robot  
│         └── main.go  （robot 入口）
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

### 运行教程
1. 导入sql文件（已包含创建数据库语句）。 mysql构建版本8.0，低版本mysql不支持 索引的 VISIBLE INVISIBLE ，需要自行移除 
2. 依赖管理。golang构建版本为go1.12.5， 使用go mod 管理依赖。低于 1.12 版本不支持 go mod，需升级。 mod 文件中 加入了 replace 解决 GFW（中国长城防火墙） 问题.
如果需要当前目录下vendor目录构建。可在更目录下执行以下代码。
`go mod vendor`
3. 配置文件。congfig/config.json, 可以按照config.json.template 编写。填上数据库地址。
4. 运行。 进入 cmd/web 目录，执行 go build .
5. 访问监听的端口。
        http://localhost:8081 (前台)
        http://localhost:8081/dashboard （后台管理）
6. 配置参数。除了配置文件外，运行时配置 存放于 mysql 中，需要进入管理后台更改。

### 开发
1. 前台 后台 基于vue-cli3构建，考虑到部署成本，代码中并没有写死host，默认无host,会默认请求当前域名。vue.config.js中默认配置了8081的反向代理开发服务器（vue-cli自带公共，讲请求代理到需要的接口地址上，避免跨域等）。

```
// 设置代理
        proxy: {
            "/v1": {
                target: "http://127.0.0.1:8081/", // 域名
                ws: true, // 是否启用websockets
                changOrigin: true, //开启代理：在本地会创建一个虚拟服务端，然后发送请求的数据，并同时接收请求的数据，这样服务端和服务端进行数据的交互就不会有跨域问题
                pathRequiresRewrite: {
                    "^/v1": "/"
                }
            },
            "/storage": {
                target: "http://127.0.0.1:8081/", // 域名
                ws: true, // 是否启用websockets
                changOrigin: true, //开启代理：在本地会创建一个虚拟服务端，然后发送请求的数据，并同时接收请求的数据，这样服务端和服务端进行数据的交互就不会有跨域问题
                pathRequiresRewrite: {
                    "^/storage": "/"
                }
            }
        }
```

2.项目还在开发中，todo list
- [x] 前台 登陆、首页、分类 列表
- [ ]  前台 标签 列表
- [ ]  前台 评论 

- [x] 后台 登陆、配置、资源媒体反向代理（防被禁）、爬虫列表、发布爬虫资源到media库、vultr创建服务器等相关操作..
- [ ]  后台 发布机器人爬虫
- [ ]  完善 资源媒体 操作


### 预览
后台管理

![dashboard.png](https://github.com/lukeyMing/socialbot/blob/master/doc/dashboard.png)

前台pc

![front2.png](https://github.com/lukeyMing/socialbot/blob/master/doc/front2.png)

前台手机 （自适应）

![mobile.png](https://github.com/lukeyMing/socialbot/blob/master/doc/mobile.png)

