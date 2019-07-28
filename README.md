# ccc

安装

```
go get github.com/funtoy/ccc
```

代码

```golang
package main

import (
	"github.com/funtoy/ccc"
)

func main() {
    var AppName = "app_name"
	ccc.Create(AppName, func() {
	    ...
		your code
		...
	})
}

```

当你编译出`app_name`后
使用

```shell
./app_name start 

```


```shell
./app_name start -d  ;后台执行 

```

```shell
./app_name stop

```

```shell
./app_name status

```