重構大學後端

### 技術

> Golang: gin, gorm
>
> Vue: element-ui,
>
> MySQL

### 功能 

> 日曆系統 CRUD
>
> 線上程式編譯器

### 架構

```bash
├─authentication
├─dao
│  ├─calendardao
│  └─userdao
├─global
│  └─database
├─init
│  ├─configinit
│  ├─routerinit
│  └─sqlinit
├─model
│  ├─calendar
│  └─user
├─router
│  ├─calendarrouter
│  ├─runprogramrouter
│  └─userrouter
└─service
    ├─calendarservice
    ├─runprogramservice
    └─userservice
```

​	主要分成 router、model、service

​	router: 負責路由分配

​	model: 定義資料結構

​	service: 業務邏輯

authentication: 用來存放 生成、檢驗 jwt

init: 初始化配置

config.ini: 存放可變動參數