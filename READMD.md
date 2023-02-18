# 项目搭建

0. 搭建基本目录
1. 搭建服务器 gin `go get -u github.com/gin-gonic/gin`
2. 数据库交互 gorm `go get -u gorm.io/gorm`
3. 加载配置文件 viper `go get github.com/spf13/viper`

## 配置文件

```yaml
server:
  port: 8080
databases:
  username:
  password:
  host: 127.0.0.1
  port: 3306
  database:
  charset: utf8
redis:
  port: 6379
```
