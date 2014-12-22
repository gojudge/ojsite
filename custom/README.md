## 用户配置文件

```
# http端口
httpport = 82
# 运行模式：dev开发模式、prod生产模式
runmode = dev

# 数据库配置
mysqluser = "username"
mysqlpass = "password"
mysqlurls = "127.0.0.1"
mysqlport = 3306
mysqldb   = "goj"

# 静态文件map映射表
static_map = "map.json"

# github oauth配置
github_client_id = xxxxxxxxxxxxxxxxxxxx
github_client_secret = xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

若有与`/conf/app.conf`相同的配置项，则会覆盖`/conf/app.conf`中的配置项。
