《初识 GOLANG & WEB FRAMEWORK》

面向对象如下：
1.有兴趣的小伙伴
2.有一定的开发语言基础

实现目标：
快手上手

##  1. Golang
   ### 1.1 安装及配置
   https://eddycjy.gitbook.io/golang/di-3-ke-gin/install
   https://studygolang.com/dl

```shell
   go env GOPATH
   go env GO111MODULE
   go env -w GO111MODULE=on
   go env -w GOPROXY=https://goproxy.cn,direct
```
在Mac环境编译成windows版本
```shell
GOOS=windows GOARCH=amd64 go build
```

### 1.2 Golang的与众不同之处

   #### 1.2.1 代码的组织方式

   #### 1.2.2 数组
   + 1. 数组：是同一种数据类型的固定长度的序列。
   + 2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
   + 3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
   + 4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
     for i := 0; i < len(a); i++ {
    }
    for index, v := range a {
    }
   + 5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
   + 6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
   + 7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
   + 8.指针数组 [n]*T，数组指针 *[n]T。


   #### 1.2.3 切片 关键字 make
   + 1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
   + 2. 切片的长度可以改变，因此，切片是一个可变的数组。
   + 3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
   + 4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
   + 5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
   + 6. 如果 slice == nil，那么 len、cap 结果都等于 0。


Golang基础
+ 实践 https://www.topgoer.com/go%E5%9F%BA%E7%A1%80/%E5%88%87%E7%89%87Slice.html
+ 理论 https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-function-call/

   

## 3.WEB框架的基础
   + 1.支持Restful api，支持参数的绑定，校验，自定义校验
   + 2.能够支持jwt token 登录 认证 授权
   + 3.ORM 支持增删改查
   + 4.支持 DB Migration 和 seed 数据
   + 5.支持Log处理
   + 6.支持Test，单元测试，Mock测试， testIf
   + 7.配置文件的热加载
   + 8.路由分组
   + 9.中间件

Gin
+ https://www.kancloud.cn/shuangdeyu/gin_book/949414
+ https://eddycjy.gitbook.io/golang/di-3-ke-gin/api-01

   
## 4. 总结
   + 简洁
   + 关键字少（25个）
   + 面向对象特征
   + 优缺点 https://studygolang.com/articles/12907
   + 资源 https://geekr.dev/awesome-golang#toc-26
   
