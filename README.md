### 长沙学院·计算机工程与应用数学学院 2015级软件工程 毕业设计

#### 主要内容

##### 题目:基于爬虫接口(实际操作是用的三方api)的酒店预订系统(订单部分未调用三方api)

##### 联系方式QQ:269812428 QQ群:778315403

##### 主要技术:微信小程序(前台客户端),vuejs(脚手架后台管理端),iris(golang1.12)+gorm+jwt,配置文件viper热加载,

##### 小程序部分

![image](https://github.com/xhaoxiong/bysj/blob/master/1554730786136.gif)

##### <a href="http://123.207.1.120:9018/admin#/login">后台部分(admin/admin)</a>

##### 前后台后端项目目录

```
bysj
│  go.mod
│  go.sum
│  main.go
└ ─ ─conf[配置信息]
│      config.yaml 
└ ─ ─conifg[热加载与viper的启用]
│      config.go
└ ─ ─log
│     平时的打印日志文件及其转储的压缩文件
└ ─ ─models[表struct]
│  │   auto_migrate.go(同步表)
│  │   init.go(数据库初始化)
│  │   feedback.go
│  │   comment.go
│  │   order.go
│  │   page_result.go
│  │   pay_record.go
│  │   user.go    
│  └ ─ ─mgodb
│  │      mgo.go(mgo的一些封装方法mgo驱动版本的) 
│  └ ─ ─redi
│         redis.go(redis的一些封装方法redigo驱动版本的)
└ ─ ─repositories[持久层]
│       AuthRepositories.go
│       CommentRepositories.go
│       DashBoardRepositories.go
│       FeedBackRepositories.go
│       HotelRepositories.go
│       OrderRepositories.go
│       PayRecordRepositories.go
│       UserRepositories.go
└ ─ ─route[路由文件夹]
│      route.go
└ ─ ─services[服务层正常来说也可以不要]
│   │   AuthRepositories.go
│   │   CommentRepositories.go
│   │   DashBoardRepositories.go
│   │   FeedBackRepositories.go
│   │   HotelRepositories.go
│   │   OrderRepositories.go
│   │   PayRecordRepositories.go
│   │   UserRepositories.go    
│   └ ─ ─hotel_api_services
│   │       city.go 
│   │       common.go
│   │       create_order.go
│   │       detail.go
│   │       real_time_inquiry.go
│   │       room_price.go
│   │       search.go
│   └ ─ ─sms_api_services
│   │       sms.go
│   └ ─ ─wechat_api_services
│           get_userinfo.go   
└ ─ ─web[前台和后台的控制层及其中间件与视图]
│   └ ─ ─admin
│   │      AuthController.go
│   │      CommentController.go
│   │      Common.go(返回给前端封装的json方法)
│   │      DashboardController.go
│   │      FeedBackController.go
│   │      IndexController.go
│   │      OrderController.go
│   │      PayRecordController.go
│   │      UserController.go
│   └ ─ ─controllers
│   │      AuthController.go
│   │      CommentController.go
│   │      Common.go(返回给前端封装的json方法)
│   │      FeedBackController.go
│   │      OrderController.go
│   │      UserController.go
│   └ ─ ─middleware
│   │      jwt.go
│   └ ─ ─views
│       └ ─ ─admin
│           │   index.html
│           └ ─ ─assets
│               └ ─ ─css
│               └ ─ ─fonts
│               └ ─ ─img
│               └ ─ ─js
└ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ 
```
#### 项目安装与运行

```


```
