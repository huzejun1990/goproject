package main

import "fmt"

// 演示golang中整数类型的使用
func main() {

	var i int = 1

	fmt.Println("i=", i)

	//测试一个int8的范围 -128~127
	// 其他的 int16 int32 int64,类型推。。。
	var j int8 = 127
	fmt.Println("j=", j)

	//测试一下，uint8的范围，其它的 uint16,uint32,uint64一样
	var k uint16 = 256
	fmt.Println("k=", k)

	//int, uint,rune,byte的使用
	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	var c byte = 255

	fmt.Println("b=", b)
	fmt.Println("c=", c)

}
