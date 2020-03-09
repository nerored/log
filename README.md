[![MIT licensed][10]][11] [![Build Status][3]][4] [![Go Report Card][5]][6] [![Coverage Statusd][7]][8]

# log
>
## 介绍
&emsp;log 是golang语言编写的日志打印库。包含了所有终端支持的打印格式，如前景色，背景色，粗体，闪烁等等。同时支持时间标签，输出文件位置以及函数名称，堆栈信息。    
&emsp;多种终端格式通过设计可以叠加一起使用。方便的自定义接口，可以组合出任意的输出格式，方便进行信息筛选。支持多通道输出，可以输出至终端，文件，网络，数据库等等。    
&emsp;支持多种渠道的预警通知，可以自定义日志预警级别。出发对应级别日志后能通过邮件，钉钉等渠道进行实时预警。    
## 安装
&emsp;仅仅需要在你的工程里面添加如下代码,其他的log会自己完成
```go
import (
	"github.com/nerored/log"
)
```
## 使用
&emsp;如果你没有特别的调整需求，这样就可以开始使用了。如果需要做部分的定制，目前提供如下自定义接口
1. 自定义时间标签的格式
```go
func SetTimeFormat(format string) {
	sharedPrinter.timeFormat = format
}
```
2. 添加新的输出通道
```go
func AddWriter(writers ...writer) {
	for _, w := range writers {
		if w == nil {
			continue
		}

		sharedPrinter.writerList = append(sharedPrinter.writerList, w)
	}
}
```
&emsp;如果需要自己定义输出渠道需要实现writer接口，具体用法以及规则参考源码，很容易：)。
```go
type writer interface {
	init()
	exit()
	needColor() bool
	info() io.Writer
	erro() io.Writer
}
```
3. TODO::启用输出到文件，定义文件名格式，分割规则（时间，大小)    
4. TODO::启用预警通知，可以选择默认的通知方式，也可以增加自己的通知     
5. TODO::定义网络接受方式 
## 示例
&emsp;log 默认提供6个等级（不同格式）的日志输出，直接调用对应的函数即可。  
![][1]  
&emsp;也可以使用能自定义的Ulog接口
```go
Ulog(level int8, flags int16, format string, args ...interface{})
```
&emsp;level 表示日志的等级，影响日志的前缀标签以及基本颜色
```go
const (
	LOG_LEVEL_USER = iota
	LOG_LEVEL_DEBU
	LOG_LEVEL_INFO
	LOG_LEVEL_TRAC
	LOG_LEVEL_WARN
	LOG_LEVEL_ERRO
	LOG_LEVEL_FATA
)
```
&emsp;flags 表示需要组合的日志基本信息,通过位操作"|",完成功能的组合
```go
const (
	PRINT_TIMELAB = 0x0100  //打印时间标签，输出结果受 SetTimeFormat 的影响
	PRINT_FILELOC = 0x0200  //打印日志标记的文件位置
	PRINT_FUNCNAM = 0x0400  //打印调用日志输出的函数名称
	PRINT_STACKIN = 0x0800  //是否需要输出堆栈信息

	PRINT_STACKDP = 0x00FF  //输出的堆栈信息深度，值范围为0-255
    
    //默认日志接口使用的组合方式
	PRINT_DEFINE = PRINT_TIMELAB | PRINT_FILELOC | PRINT_FUNCNAM
    //快速的日志跟踪打印
    PRINT_UTRACE = PRINT_DEFINE | PRINT_STACKIN | 10
)
```    
![][2]    
## 开发状态    
&emsp;目前完成了输出到终端的部分，后续会完成输出到文件，attach 到网络连接。以及日志预警.
## Have Fun !
[1]:https://github.com/nerored/log/blob/master/images/IMG_0539.JPG
[2]:https://github.com/nerored/log/blob/master/images/IMG_0540.JPG
[3]: https://travis-ci.org/nerored/log.svg?branch=master
[4]: https://travis-ci.org/nerored/log
[5]: https://goreportcard.com/badge/github.com/nerored/log
[6]: https://goreportcard.com/report/github.com/nerored/log
[7]: https://codecov.io/gh/nerored/log/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/nerored/log
[10]: https://img.shields.io/badge/license-MIT-blue.svg
[11]: LICENSE
