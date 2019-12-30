# a simple web lib for gin

e.g
``` go
web.GetMapping("/", func(c *ct.Context) interface{} {
    ...
    return xx
}
web.Engine.LoadHTMLGlob("html/*")
web.Engine.Run(":8080")
```
