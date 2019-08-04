//04. Variable life cycle and assignmeng
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
