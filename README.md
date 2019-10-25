# beego-orm-create-table
创建beego框架 mysql 表结构对应的struct 结构

# 使用方法

1、git pull https://github.com/zhangyh0924/beego-orm-create-table.git

2、拷贝目录中cmd 文件夹到项目目录src下

3、配置cmd/config/database.conf中数据库相关配置
```
host = 127.0.0.1  //连接地址
port = 3306  //连接端口
database = card //数据库名称
username = root // 数据库用户名
password = root // 数据库用户密码
charset = utf8mb4 // 数据库编码格式
prefix = ds_    //数据库前缀
```
3、配置完成之后可以直接执行
```
go run main.go -h

usage  rum main.go -t=tablename -m=ishump -p=filepath
option:
  -h    帮助
  -m    文件是否驼峰 true false 默认false
  -p string
        文件保存路径
  -t string
        生成表的名称
//例如：生成订单表结构，以驼峰命名，放到当前目录 test/model 下
go run main.go -t=ds_order -m=true -p=./test/model
执行结果
./test/model/Order.go is success
over


```

4、默认生成所有表结构，以下划线命名，生成文件放在当前目录model/下

