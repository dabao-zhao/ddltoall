# ddltoall

把 [go-zero](https://github.com/zeromicro/go-zero) 的 [goctl](https://github.com/zeromicro/go-zero/tree/master/tools/goctl) 摘出来的东西

## 安装

```shell
go install github.com/dabao-zhao/ddltoall@latest
```

## 使用

```
// -s sql 文件
// -d 输出目录
// -u sql dsn
// -t 表名匹配规则

ddltoall ddl -s=".\test.sql" -d="test"

ddltoall datasource -u="root:123456@tcp(127.0.0.1:3306)/test" -t="*" -d="test"
```