[toc]

# 开发环境的搭建
## 环境的搭建
## hello world
``` go
//01.hello world.main
package main
import "fmt"
func main(){
    fmt.Println("hello world!")
}
```

# 程序结构
## go文件的结构
``` go
//声明此文件属于哪一个包
package xxx
//声明该文件需要用到哪些包
import xxx
//定义包级别的常亮，变量，函数等
var xxx
func xxx(){

}
```

## 变量的声明
``` go
//02.Variable declaration.main
package main
import "fmt"
func main(){
    //单变量声明
    var x int  = 0
    //var x int
    //var x = 0;
    //x := 0

    //多变量同时声明
    var y,z int = 0 ,0
    //var y,z int
    //var y,z = 0,0
    //y,z:=0
    fmt.Println(x,y,z)
}
```
变量声明的基本格式可以看做是
``` go
var name type = expression
```
但是如上代码所示，可以缩减成
``` go
var name = expression//根据表达式推断变量类型
var x int //直接将x的值初始化为0（系统默认初始化）
x := 0 //这是一种短变量声明
```
另外多变量声明在代码中也有展示，就不解释了
## 指针与new

## 生命周期与赋值

## 项目的结构