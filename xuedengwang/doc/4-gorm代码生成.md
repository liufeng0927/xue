
1、通过工具在数据库新建表格

2、在`cmd/gen/method.go`中新增
```go
//=============================== user_auth ==================================//
type I{表名}Method interface {
}
```
这个接口可以自定义查询方法，可用于处理一些复杂的sql。

语法定义请查阅：[Gorm Gen Annotation Syntax](https://gorm.io/zh_CN/gen/sql_annotation.html)

3、在`cmd/gen/generator.go`中新增
```go
g.ApplyInterface(func(自定义方法的接口){}, g.GenerateModel("表名"))
```
















