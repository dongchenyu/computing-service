# 开发 web 服务程序
### 1.概述
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。
任务目标
+ 熟悉 go 服务器工作原理
+ 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。
+ 使用 curl 工具访问 web 程序
+ 对 web 执行压力测试

### 2.开发cloudgo
在这里我们需要选择一个合适的框架，我选择的是martini
martini框架是使用Go语言进行模块化web应用与服务的开发框架，专门用来处理Web相关内容，而且这个框架比较容易上手
首先要先安装相关框架
```
go get github.com/codegangsta/martini
```
安装好之后我们编写一段helloworld程序来验证一下
```
package main

import "github.com/codegangsta/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
```
我们把这段程序命名为test.go
然后在控制台中输入go run test.go

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111155408807.png)

然后打开http://localhost:3000
看到如下结果

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111155448271.png)

对上面的程序有几点需要说明
+ m := martini.Classic()用来创建一个的martini实例
+ m.Get("/", func() string { ... })用来接收对\的GET方法请求，其中第二个参数是对请求的处理方法
+ m.Run()用来运行服务器
+ 端口号在这里没有特意设置，所以默认是3000

接下来开始进行我们的cloudgo开发
其中main.go可以直接使用老师给出来的代码
```
package main

import (
    "os"
    "github.com/gocloud/server"
    flag "github.com/spf13/pflag"
)
const (
    PORT string = "8080" 
)
func main() {
    port := os.Getenv("PORT") 
    if len(port) == 0 {
        port = PORT
    }
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
    server.NewServer(port)
}
```
server.go是一个helloworld，跟上面的很类似
```
package server
import "github.com/codegangsta/martini"

func NewServer(port string) {   
    m := martini.Classic()

    m.Get("/", func(params martini.Params) string {
        return "hello world"
    })

    m.RunOnAddr(":"+port)   
}
```
其中在结尾多了RunOnAddr， 用来运行程序监听端口，在这里是8080端口
在控制台中输入go run main.go 运行程序

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111160749897.png)

然后打开http://localhost:8080

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111160918737.png)

可以看到运行成功
### 3.使用 curl 测试
我们另打开一个控制台，输入
```
cuil -v http://localhost:8080/
```
得到结果如下

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111161213432.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

### 4.使用 ab 测试

在控制台中输入
```
ab -n 1000 -c 100 http://localhost:8080/
```
结果如下

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111161445174.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)
