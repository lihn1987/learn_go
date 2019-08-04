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
所谓指针也就是指向数据地址的一个数据类型。</br>
其类型的声明方式，如果是int类型的指针，其类型写为
``` go
var variable * int 
```
若为bool类型的指针，其类型写为
``` go 
var variable * bool  
```
而对指针取值的方式也比较简单，就是
```
*variable
```

至于指针我们常常会考虑的合适释放指针空间的问题，这里暂不考虑，后面章节会说，go语言有他自己的一套来处理指针资源的声明周期。
下面我们来看一套指针相关的代码示例

``` go
//03.Point and New.main.go
package main

import "fmt"

func main(){
	var x int = 15;
	var  x_point *int = &x;
	fmt.Println(x_point)
	fmt.Println(*x_point)
	fmt.Println(x)
	test_func1(*x_point)
	fmt.Println(*x_point)
	fmt.Println(x)
	test_func2(x_point)
	fmt.Println(*x_point)
	fmt.Println(x)

	///////////////
	fmt.Println("==================")

	x_point = new(int);
	*x_point = 5;
	fmt.Println(x_point)
	fmt.Println(*x_point)
	fmt.Println(x)
	test_func1(*x_point)
	fmt.Println(*x_point)
	fmt.Println(x)
	test_func2(x_point)
	fmt.Println(*x_point)
	fmt.Println(x)
}

func test_func1(value int){
	value = 100;
}

func test_func2(value *int){
	*value = 100;
}
```
其输出结果为
```
0xc000066058
15
15
15
15
100
100
==================
0xc0000660f0
5
100
5
100
100
100
```
可以看出在函数*x_point的值永远与x的值保持一致（相同的指针）

在函数内部修改指正指向内容的值，其值依然会发生变化。

另外，指针的赋值除了对已有变量进行取地址操作（```&```）外，也可以直接使用new 分配一个新的变量
## 生命周期与赋值

在Go语言里面，变量的生命周期似乎并不是像C语言那么明确，比如一个大括号外变量就已经被销毁，堆和栈分的那么清楚。go语言似乎在故意让我们忘记生命周期这回事情，只要这个变量还有用，那么他就没有被销毁，而我们并不需要必须知道他的生命周期.

而赋值这块，go语言的特点是可以单次给你个值赋值，也可以单次给多个值赋值，下面看示例
``` go
//04. Variable life cycle and assignment
package main

import "fmt"
func main(){
	//返回函数内变量指针
	var xxx *int = test_func();
	fmt.Println(*xxx);

	//循环中重复定义test
	for i:=0; i < 3; i++{
		var test int = 5
		fmt.Println(test)
		test +=1;
	}
	//赋值
	var a,b,c int;
	a=1
	fmt.Println(a);
	b,c=2,3
	fmt.Println(b,c);
}
func test_func() *int{
	var xxx int = 100;
	return &xxx
}
```

输出内容为
```
100
5
5
5
1
2 3
```
可以看出函数返回指针时，虽然指向的值在函数中，但是在返回后，指针的值依然可以并没有被销毁。

而循环中声明的局部变量和其他语言一样。每次循环都重新创建，上次值得结果并不影响该次。

## 类型与声明                                                                   
为了类型隔离和注记的原因我们可能会使用其他的名称来表示一个变量类型，如下代码所示。
``` go
//05.Type Declaration/main.go
package main

import "fmt"
type Length int
type Height int
func main(){
	var v1 Length = 10;	
	var v2 Height = 20;
	var v3 int = 30;

	v1=40;
	//下面的方法都会导致编译错误
	//v3=v1;
	//v1=v3;

	fmt.Println(v1,v2,v3)
}
```


类型声明还保证了不同类型变量之间无法直接赋值。


## package and file