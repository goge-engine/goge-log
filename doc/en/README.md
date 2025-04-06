# Introduction
**Goge Log** is a logging package designed for the **Goge Engine**.  
This package provides several functions for logging output.

## Documentation
1. Initialization
```go
package main

import (
    "github.com/goge-engine/goge-log"
    "github.com/goge-engine/goge-basis"
)

func main() {
    var logger GogeLogClass    // Initialize the GogeLogClass struct
    logger.Init("./your_log_file_here.log")              // Use the Init() method to initialize the log output
}
```  

2. Usage  
**Goge Log** provides several methods for the `GogeLogClass` struct to output logs, which are:  
- `Log(logContext string)` is used to output logs
- `Error(errorContext error)` is used to output error messages 
- `Log(warningContext gogeBasis.Warning)` is used to output warning messages
- `DevelopmentLog(logHead string, logContext string)` is used to output logs during development, allowing customization of the log level or type by filling in the `logHead` parameter
- `ErrorAndWarningProcess(errorContext error, warningContext gogeBasis.Warning)` is used to automatically output error and warning messages, without the need to check if `error` or `gogeBasis.Warning` is empty, simply call it directly

3. Note  
**It is absolutely essential to include `defer yourLogger.Close()` in the `main()` function, otherwise, there will be issues with unclosed file handles.**