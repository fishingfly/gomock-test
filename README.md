<!--
 * @Author: zhounanjun
 * @Date: 2022-04-22 12:21:43
 * @LastEditors: zhounanjun
 * @LastEditTime: 2022-04-22 23:06:15
 * @Description: 请填写简介
-->

## 操作文档
### 手动对interface生成mock代码

操作命令,生成mock文件
```
mockgen -destination=mocks/mock_doer.go -package=mocks -source=doer/doer.go  Doer
```
在工程更目录下执行, -destination指定生成mock代码的路径, -package=mocks 指定mock代码所处的包, -source=doer/doer.go 指定需要mock的接口的所在文件, 最后的Doer指source文件中哪个interface是需要mock的

### 结合go:generate使用GoMock
在需要mock的interface上添加一行go:generate注释:
```
//go:generate mockgen -destination=../mocks/mock_doer1.go -package=mocks -source=./doer1.go
type Doer1 interface {
	DoSomething1(int, string) error
}
```
然后在工程根目录中执行:
```
go generate ./...
```
在mock目录下又会生成一个mock文件

### golang 子测试
见/user/add_test.go