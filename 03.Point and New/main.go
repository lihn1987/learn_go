//03.point.main.go
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