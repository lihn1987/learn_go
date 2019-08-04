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