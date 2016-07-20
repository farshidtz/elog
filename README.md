# elog
[![GoDoc](https://godoc.org/github.com/farshidtz/elog?status.svg)](https://godoc.org/github.com/farshidtz/elog)
elog extends Go's built-in [log](https://golang.org/pkg/log) package to enable simple levelled logging and to modify the formatting. 

## Debugging mode
The debugging mode can be enabled by setting the configured environment variable ("DEBUG" by default) to 1. Alternatively, the debugging can be enabled using a flag.

## Usage
Get the package

    go get github.com/farshidtz/elog\
Import
```go
import "github.com/farshidtz/elog"
```

Debugging not enabled
```go
10  // Initialize with the default configuration
11	logger := elog.New("[main] ", nil)
12	
13	logger.Println("Hello world!")
14	// 2016/07/14 16:47:04 [main] Hello world!
15	
16	logger.Debugln("Hello world!")
17	// Nothing printed
18	
19	logger.Fatalln("Hello world!")
20	// 2016/07/14 16:47:04 [main] Hello world!
21	// exit with status 1
```
Debugging enabled
```go
10  // Initialize with the default configuration
11	logger := elog.New("[main] ", nil)
12	
13	logger.Println("Hello world!")
14	// 2016/07/14 16:47:04 [main] main.go:13: Hello world!
15	
16	logger.Debugln("Hello world!")
17	// 2016/07/14 16:47:04 [debug] main.go:16 Hello world!
18	
19	logger.Fatalln("Hello world!")
20	// 2016/07/14 16:47:04 [main] main.go:19 Hello world!
21	// exit with status 1
```

## Configuration
Debugging enabled
```go
10	logger := elog.New("[I] ", &logger.Config{
11	  TimeFormat: time.RFC3339, 
12	  DebugPrefix: "[D] ", 
13	})
14	
15	logger.Println("Hello world!")
16	// 2016-0714T16:57:15Z [I] main.go:15 Hello world!
17	
18	logger.Debugln("Hello world!")
19	// 2016-0714T16:57:15Z [D] main.go:18 Hello world!
20	
21	logger.Fatalln("Hello world!")
22	// 2016-0714T16:57:15Z [I] main.go:21 Hello world!
23	// exit with status 1
```

## Initialize with init()
Alternatively, you can initialize the logger once and use it anywhere in the package:
```go
import "github.com/farshidtz/elog"

var logger *elog.Logger

func init() {
	logger = elog.New("[main] ", nil)
}
```

## Enable debugging with flag
```go
package main

import (
	"github.com/farshidtz/elog"
	"flag"
)

var debugFlag = flag.String("d", false, "Enable debugging")
func main() {
	logger := elog.New("[main] ", &logger.Config{
	  DebugEnabled: debugFlag,
	})
	
	// 
}

