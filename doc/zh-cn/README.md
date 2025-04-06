# 简介
**Goge Log**是一个用于**Goge Engine**的日志输出包。  
该包提供了一些日志输出的函数

## 文档
1.初始化
```go
package main

import (
    "github.com/goge-engine/goge-log"
    "github.com/goge-engine/goge-basis"
)

func main() {
    var logger GogeLogClass    //初始化GogeLogClass结构体
    logger.Init("./your_log_file_here.log")              //使用Init()方法初始化日志输出
}
```  

2.使用  
**Goge Log**提供了一些日志输出的`GogeLogClass`结构体方法，它们为：  
- `Log(logContext string)`用于输出日志
- `Error(errorContext error)` 用于输出错误信息 
- `Warning(warningContext gogeBasis.Warning)`  用于输出警告信息
- `DevelopmentLog(logHead string, logContext string)`  用于开发时输出日志，可以以填入`logHead`参数的方式自定义输出日志等级或类型
- `ErrorAndWarningProcess(errorContext error, warningContext gogeBasis.Warning)`  用于自动输出错误和警告信息，无需判断`error`或者`gogeBasis.Warning`是否为空，直接调用即可

3.注意
**一定一定要在`main()`函数中加入`defer yourLogger.Close()`，否则会出现文件句柄没有关闭的问题**