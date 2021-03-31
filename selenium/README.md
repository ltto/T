# selenium-NoErr 版本

## 背景

> 操作中所有错误统统panic
>
> 节省代码行数,更快的进行代码编写
>
> 灵感来自于[selenium-java](https://mvnrepository.com/artifact/org.seleniumhq.selenium/selenium-java)

## 使用

```go
package main

import (
	"flag"
	"fmt"
	"github.com/ltto/T/selenium"
)

func main() {
	user := flag.String("user", "admin", "user")
	pwd := flag.String("pwd", "admin", "pwd")
	host := flag.String("h", "127.0.0.1", "host")
	port := flag.Int("p", 80, "port")
	flag.Parse()

	driver, err := selenium.NewChromeDriver()
	if err != nil {
		panic(err)
	}
	//defer driver.Ser.Stop()
	driver.Get(fmt.Sprintf("http://%s:%d", *host, *port))
	driver.FindElementByCSSSelector("body > xxx").SendKeys(*user)
	driver.FindElementByCSSSelector("body > div.loginBlock > div.bgContainer > p:nth-child(2) > input").SendKeys(*pwd)
	driver.FindElementByCSSSelector("body > div.loginBlock > p.describeW.describeWW.floatR > input").Click()
	driver.FindElementByCSSSelector("body > div.loginBlock > div.bgContainer > a").Click()
}
```